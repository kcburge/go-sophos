package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/ChimeraCoder/gojson"
	"github.com/esurdam/go-sophos"
)

type Definition struct {
	Description, Name, Link string
}

type Swag struct {
	Paths map[string]map[string]interface{}
}

type endpoint struct {
	Title, Name string
	Routes      []string
	Path        string
	SubTypes    []subtype
}

type subtype struct {
	Name    string
	JsonTag string
	GetPath string
}

type data struct {
	Endpoints []endpoint
}

func (ep data) ExecuteTemplate(w io.Writer, temp string) error {
	tmpl, err := template.New("").Parse(temp)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, ep)
	if err != nil {
		return err
	}
	return nil
}

func (ep endpoint) ExecuteSubTemplate(w io.Writer) error {
	tmpl, err := template.New("").Parse(subTypeTemplate)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, ep)
	if err != nil {
		return err
	}
	return nil
}

var funcTemp = `
// Definitions implements the Resource interface and returns a map of {{.Title}}'s RestObjects
func({{.Title}}) Definitions() map[string]RestObject {
	return map[string]RestObject{
		{{range .SubTypes}}"{{.Name}}": {{.Name}}{},
		{{end}}
	}
}

// GetPath implements RestObject interface and returns the {{.Title}} GET path
func({{.Title}}) GetPath() string {
	return "{{.Path}}"
}
{{range .SubTypes}}
// GetPath implements RestObject interface and returns the {{.Name}} GET path
func({{.Name}}) GetPath() string {
	return "{{.GetPath}}"
}
{{end}}

// ApiRoutes returns all known {{.Title}} paths
func({{.Title}}) ApiRoutes() []string{
	return []string{
		{{range $i, $element := .Routes}}"/api{{$element}}",
		{{end}}
	}
}`

func (ep endpoint) ExecuteTemplate(w io.Writer) error {
	tmpl, err := template.New("").Parse(funcTemp)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, ep)
	if err != nil {
		return err
	}
	return nil
}

var subTypeTemplate = `
// {{.Title}} is a generated struct representing the Sophos {{.Name}} Node Leaf
// GET {{.Path}}
type {{.Title}} struct {
	{{range .SubTypes}}{{.Name}} {{.Name}} ` + "`json:\"{{.JsonTag}}\"`" + `
	{{end}}
}

// Definitions implements the Resource interface and returns a map of {{.Title}}'s RestObjects
func({{.Title}}) Definitions() map[string]RestObject {
	return map[string]RestObject{
		{{range .SubTypes}}"{{.JsonTag}}": {{.Name}}{},
		{{end}}
	}
}

{{range .SubTypes}}
// GetPath implements RestObject and returns the {{.Name}} GET path
func({{.Name}}) GetPath() string {
	return "/api{{.GetPath}}"
}
{{end}}`

var objTemplate = `
var Objects = map[string]Resource {
	{{range .Endpoints}}"{{.Name}}": {{.Title}}{},
	{{end}}
}
`

var nodeTemplate = `
var NodeLeaves = map[string]RestObject {
	{{range .Endpoints}}"{{.Name}}": {{.Title}}{},
	{{end}}
}
`

var client *sophos.Client

func main() {

	var ep, token string
	if len(os.Args) == 2 {
		ep = os.Args[1]
		token = os.Args[2]
	}

	if ep == "" {
		ep = os.Getenv("ENDPOINT")
	}
	if token == "" {
		token = os.Getenv("TOKEN")
	}

	if ep == "" || token == "" {
		panic("need endpoint and token as args or from env ($ENDPOINT, $TOKEN)")
	}

	var err error
	client, err = sophos.New(ep, sophos.WithAPIToken(token))
	if err != nil {
		log.Fatal(err)
	}

	var res []Definition
	r, err := client.Get("/api/definitions")
	if err != nil {
		log.Fatal(err)
	}
	err = r.MarshalTo(&res)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("types/generated.go")
	if err != nil {
		log.Fatal(err)
	}

	var (
		objEndpoints  = data{}
		nodeEndpoints = data{}
	)

	f.Write([]byte(`// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

// RestObject is an interface used to retrieve endpoint paths from a generated Object
type RestObject interface {
	// GetPath returns the Object's GET path
	GetPath() string
}

// Resource is an interface representing an endpoint (Sophos Definition) 
type Resource interface {
	// Definitions returns the Resource's RestObjects
	Definitions() map[string]RestObject
}
`))

	for _, def := range res {
		path := fmt.Sprintf("/api/%s", strings.ToLower(def.Name))
		if def.Name != "Nodes" {
			path = fmt.Sprintf("/api/nodes/%s", strings.ToLower(def.Name))
		}

		// get the struct
		r, err := client.Get(path)
		if err != nil {
			log.Println("could not get path")
			continue
		}

		ep := endpoint{
			Title: toCamelInitCase(def.Name, true),
			Name:  def.Name,
			Path:  path,
		}

		// get definition
		resp, err := client.Get(def.Link)
		if err != nil {
			log.Fatal("could not get definition path")
		}

		var swag Swag
		if err = resp.MarshalTo(&swag); err != nil {
			log.Fatal("could not unmarshal swag")
		}

		// add the paths to the endpoint
		for path := range swag.Paths {
			ep.Routes = append(ep.Routes, path)
		}

		// fetch the subTypes of Objects
		if def.Name != "Nodes" {
			// error here means its not a node endpoint, but an Object
			// go through all the paths and fetch a all structs provided by the GET method
			for path, defs := range swag.Paths {
				ep.Routes = append(ep.Routes, path)
				for method := range defs {
					if method == "get" && !strings.Contains(path, "{ref}") {
						parts := strings.Split(path, "/")
						name := parts[len(parts)-2]

						ep.SubTypes = append(ep.SubTypes, subtype{
							Name:    toCamelInitCase(def.Name+"_"+name, true),
							JsonTag: def.Name + "_" + name,
							GetPath: path,
						})
						fmt.Printf("make %s\n", toCamelInitCase(def.Name+"_"+name, true))
						byt, err := makeStructBytes(path, def.Name+"_"+name)
						if err != nil {
							log.Printf("could not generate gojson from %s: %s", path, err.Error())
							continue
						}

						// read all the Object lines to remove duplicate Type declarations
						buf := bytes.NewBuffer(byt)
						outBuf := bytes.NewBuffer([]byte{})
						for {
							line, err := buf.ReadString('\n')
							if err != nil {
								if err == io.EOF {
									break
								}

								log.Fatalf("read file line error: %v", err)
								return
							}
							_ = line
							if strings.Contains(line, "`json:\"_type\"`") {
								continue
							}
							if strings.Contains(line, "`json:\"_ref\"`") {
								continue
							}
							outBuf.Write([]byte(line))
						}

						// write the generated Object struct to the file, removing package name
						_, _ = outBuf.ReadString('\n')
						f.Write(outBuf.Bytes())
					}
				}
			}

		}

		// write the node data
		var isNodePath = true
		byt, err := gojson.Generate(r.Body, gojson.ParseJson, ep.Title, "main", []string{"json"}, false)
		if err != nil {
			isNodePath = false
			objEndpoints.Endpoints = append(objEndpoints.Endpoints, ep)
			err := ep.ExecuteSubTemplate(f)
			if err != nil {
				fmt.Println(err.Error())
			}

			continue
		}

		if def.Name == "Nodes" {
			// read all the Object lines to remove apitoken
			buf := bytes.NewBuffer(byt)
			outBuf := bytes.NewBuffer([]byte{})
			for {
				line, err := buf.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						break
					}

					log.Fatalf("read file line error: %v", err)
					return
				}
				_ = line
				if strings.Contains(line, "Auth_apiTokens") {
					outBuf.Write([]byte("Auth_apiTokens map[string]string `json:\"auth.api_tokens\"`\n"))
					buf.ReadString('\n')
					buf.ReadString('\n')
					continue
				}
				outBuf.Write([]byte(line))
			}
			byt = outBuf.Bytes()
		}

		// write the generated struct to the file
		buf := bytes.NewBuffer(byt)
		_, _ = buf.ReadString('\n')

		f.Write(buf.Bytes())

		if isNodePath {
			nodeEndpoints.Endpoints = append(nodeEndpoints.Endpoints, ep)
		}
		ep.ExecuteTemplate(f)
	}

	// Write the objEndpoints struct
	objEndpoints.ExecuteTemplate(f, objTemplate)
	// Write the nodeEndpoints struct
	nodeEndpoints.ExecuteTemplate(f, nodeTemplate)
}

func makeStructBytes(path, name string) ([]byte, error) {
	resp, err := client.Get("/api" + path)
	if err != nil {
		log.Printf("could not get path: %s", path)
		return nil, err
	}

	return gojson.Generate(resp.Body, gojson.ParseJson, toCamelInitCase(name, true), "main", []string{"json"}, false)
}

var numberSequence = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
var numberReplacement = []byte(`$1 $2 $3`)

func addWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}
