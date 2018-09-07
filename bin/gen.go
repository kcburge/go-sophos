package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/ChimeraCoder/gojson"
	"github.com/esurdam/go-sophos"
)

var (
	client *sophos.Client

	numberSequence    = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
	numberReplacement = []byte(`$1 $2 $3`)
	header            = `// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)
`
)

//noinspection GoDuplicate
type (
	definition struct {
		Description string
		Name        string
		Link        string // path to definition
		Swag        *swag  // the definition from the UTM
		Node        *node  // a representation of the Node
	}
	// map[path]map[method]methodDescription
	swag struct {
		Paths map[string]methodMap
	}
	methodMap map[string]methodDescriptions

	methodDescriptions struct {
		Description string
		Parameters  []parameter
		Tags        []string
		Responses   map[int]struct{ Description string }
	}

	parameter struct {
		Name        string
		In          string
		Description string
		Type        string
		Required    bool
	}
	node struct {
		Definition  *definition
		Title, Name string
		Bytes       string
		Routes      []string
		Path        string
		Methods     []string
		SubTypes    []subtype
		References  []string
		Paths       map[string]methodMap
	}
	subtype struct {
		Name        string
		JsonTag     string
		GetPath     string
		GetPaths    []string
		PutPath     string
		PostPath    string
		DeletePath  string
		PatchPath   string
		HasRef      bool
		Bytes       string
		MethodDescs methodMap
		Node        *node
		IsPlural    bool
		IsType      bool
	}
)

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

	var dd []definition
	r, err := client.Get("/api/definitions")
	if err != nil {
		log.Fatal(err)
	}

	err = r.MarshalTo(&dd)
	if err != nil {
		log.Fatal(err)
	}

	for _, def := range dd {
		err := def.process()
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create("types/" + strings.ToLower(def.Name) + ".go")
		if err != nil {
			log.Fatal(err)
		}

		f.Write([]byte(header))

		// enc := json.NewEncoder(os.Stdout)
		// enc.SetIndent("", "    ")
		// if err := enc.Encode(def.Node); err != nil {
		// 	panic(err)
		// }
		fmt.Printf("writing %s\n", def.Name)
		err = def.Node.ExecuteTemplate(f)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func (def *definition) process() error {
	// get definition itself from UTM
	resp, err := client.Get(def.Link)
	if err != nil {
		return err
	}

	if err = resp.MarshalTo(&def.Swag); err != nil {
		return err
	}

	// format the path, see if its a node endpoint
	path := fmt.Sprintf("/api/%s", strings.ToLower(def.Name))
	if def.Name != "Nodes" {
		path = fmt.Sprintf("/api/nodes/%s", strings.ToLower(def.Name))
	}

	// the node will represent this definition
	def.Node = &node{
		Definition: def,
		Title:      toCamelInitCase(def.Name, true),
		Name:       def.Name,
		Path:       path,
	}

	ep := def.Node

	def.Node.fetch()
	ep.Paths = def.Swag.Paths
	if def.Name == "Nodes" {
		return nil
	}
	// Swag.Paths contains a mapping of path -> map[method]methodDescription
	for path, methodMap := range def.Swag.Paths {
		// add the path to the known Node Paths
		ep.Routes = append(ep.Routes, path)

		// make a human readable name
		parts := strings.Split(path, "/")
		name := parts[len(parts)-2]
		usedby := parts[len(parts)-1] == "usedby"
		if usedby {
			// /objects/user_preferences/webadmin/{ref}/usedby should map to webadmin and not ref
			name = parts[len(parts)-3]
		}

		// parse each method and generate subtypes
		for method := range methodMap {
			// add the method to the known Node methods
			ep.AddMethod(method)
			s := subtype{
				Name:        toCamelInitCase(def.Name+"_"+name, true),
				JsonTag:     def.Name + "_" + name,
				MethodDescs: methodMap,
				Node:        def.Node,
			}

			if s.Name == "StatusStatus" {
				s.Name = "StatusVersion"
			}

			if strings.Contains(path, "{ref}") {
				// HasRef typically means we can de-pluralize the struct returned from client in "get"
				s.HasRef = true
				// change the name to uppercase Ref
				name = strings.Replace(name, "{ref}", "{Ref}", -1)
				// add the name to the known Node references
				ep.AddReference(toCamelInitCase(def.Name+"_"+name, true))
			} else {
				s.GetPath = path
			}

			if method == "get" {
				s.GetPaths = []string{path}
				if !s.HasRef && !usedby {
					// s.GetPath = path
					// // if the path does not have ref, then we can fetch it and make a struct for it
					byt, err := makeStructBytes(&s, path, def.Name+"_"+name)
					if err == nil {
						s.Bytes = string(bytes.TrimLeft(byt.Bytes(), "\n"))
					}
				}
			}

			// TODO: make struct based on properties defined in s.MethodDescs
			if method == "put" {
				s.PutPath = path
			}
			if method == "post" {
				s.PostPath = path
			}
			if method == "delete" {
				s.DeletePath = path
			}
			if method == "patch" {
				s.PatchPath = path
			}

			// Add the subtype to the node
			ep.AddSubType(s)
		}

	}

	sort.Strings(ep.Routes)
	sort.Strings(ep.Methods)
	sort.Strings(ep.References)
	sort.SliceStable(ep.SubTypes, func(i, j int) bool { return ep.SubTypes[i].Name < ep.SubTypes[j].Name })
	return nil
}

// fetch fetches the node itself
func (n *node) fetch() error {
	// get the struct
	r, err := client.Get(n.Path)
	if err != nil {
		log.Println("could not get path")
		return err
	}
	// write the node data
	byt, err := gojson.Generate(r.Body, gojson.ParseJson, n.Title, "main", []string{"json"}, false)
	if err != nil {
		// log.Printf("could not gojson response: %s, %s\n", n.Path, err.Error())
		// error here means we will manually create a parent struct
		return err
	}

	if n.Name == "Nodes" {
		// read all the Object lines to remove apitoken
		buf := bytes.NewBuffer(byt)
		outBuf := bytes.NewBuffer([]byte{})
		for {
			line, err := buf.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}

				return fmt.Errorf("read file line error: %v", err)
			}
			_ = line

			if strings.Contains(line, "Auth_apiTokens") {
				outBuf.Write([]byte("AuthApiTokens map[string]string `json:\"auth.api_tokens\"`\n"))
				buf.ReadString('\n')
				buf.ReadString('\n')
				continue
			}

			// fix the Nodes formatted Keys
			parts := strings.Split(strings.TrimSpace(line), " ")
			for _, p := range parts {
				if p == "type" || p == "}" {
					break
				}
				if p != "" {
					line = strings.Replace(line, p, toCamelInitCase(p, true), -1)
					break
				}
			}

			outBuf.Write([]byte(line))
		}
		byt = outBuf.Bytes()
	}
	// write the generated struct to the file
	buf := bytes.NewBuffer(byt)
	_, _ = buf.ReadString('\n')
	n.Bytes = string(bytes.TrimLeft(buf.Bytes(), "\n"))

	return nil
}

func (n *node) AddReference(ref string) {
	for _, se := range n.References {
		if se == ref {
			return
		}
	}
	n.References = append(n.References, ref)
}

func (n *node) AddMethod(m string) {
	for _, se := range n.Methods {
		if se == m {
			return
		}
	}
	n.Methods = append(n.Methods, m)
}

func (n *node) AddSubType(s subtype) {
	for idx, se := range n.SubTypes {
		if se.Name == s.Name {
			if s.HasRef {
				n.SubTypes[idx].HasRef = true
			}
			if len(s.GetPaths) > 0 {
				n.SubTypes[idx].GetPaths = append(n.SubTypes[idx].GetPaths, s.GetPaths...)
			}

			if s.GetPath != "" {
				n.SubTypes[idx].GetPath = s.GetPath
			}

			if s.PutPath != "" {
				n.SubTypes[idx].PutPath = s.PutPath
			}

			if s.PostPath != "" {
				n.SubTypes[idx].PostPath = s.PostPath
			}

			if s.DeletePath != "" {
				n.SubTypes[idx].DeletePath = s.DeletePath
			}

			if s.PatchPath != "" {
				n.SubTypes[idx].PatchPath = s.PatchPath
			}

			if s.Bytes != "" {
				n.SubTypes[idx].Bytes = s.Bytes
			}

			if !se.IsPlural && s.IsPlural {
				n.SubTypes[idx].IsPlural = s.IsPlural
			}

			if !se.IsType && s.IsType {
				n.SubTypes[idx].IsType = s.IsType
			}
			return
		}
	}

	n.SubTypes = append(n.SubTypes, s)
}

func addWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}

func makeStructBytes(s *subtype, path, name string) (*bytes.Buffer, error) {
	if !strings.HasPrefix(path, "/api") {
		path = "/api" + path
	}
	resp, err := client.Get(path)
	if err != nil {
		log.Printf("could not get path: %s\n", path)
		return nil, err
	}

	name = toCamelInitCase(name, true)
	if name == "StatusStatus" {
		name = "StatusVersion"
	}
	byt, err := gojson.Generate(resp.Body, gojson.ParseJson, name, "main", []string{"json"}, false)
	if err != nil {
		log.Printf("could not gojson response: %s\n", path)
		return nil, err
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

			return nil, fmt.Errorf("read file line error: %v", err)
		}
		_ = line
		if strings.HasPrefix(line, "type") {
			if strings.HasSuffix(line, " []struct {\n") {
				s.IsPlural = true
				// make pluralized
				newLine := strings.Replace(line, " []struct {", "s []"+name, 1)
				outBuf.Write([]byte(newLine))
				outBuf.Write([]byte("// " + name + " is a generated Sophos object\n"))
				newType := strings.Replace(line, " []struct {", " struct {", 1)
				outBuf.Write([]byte(newType))
				continue
			}
			// if strings.HasSuffix(line, " []interface{}\n") {
			// 	// make pluralized
			// 	newLine := strings.Replace(line, " []interface{}", "s []"+toCamelInitCase(name, true), 1)
			// 	outBuf.Write([]byte(newLine))
			// 	newType := strings.Replace(line, " []interface{}", " interface{}", 1)
			// 	outBuf.Write([]byte(newType))
			// 	continue
			// }
		}

		if strings.Contains(line, "`json:\"_type\"`") {
			s.IsType = true
			line = strings.Replace(line, "Type", "_type", -1)
		}
		if strings.Contains(line, "`json:\"_ref\"`") {
			line = strings.Replace(line, "Ref", "Reference", -1)
		}
		outBuf.Write([]byte(line))
	}

	// write the generated Object struct to the file, removing package name
	_, _ = outBuf.ReadString('\n')
	return outBuf, nil
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

// Methods returns the {{.Title}}'s available HTTP methods
// func({{.Title}}) Methods() []string {
// return []string{
// {{range $i, $element := .Methods}}"{{$element}}",
// {{end}}
// }
// }

var nodeTemplate = `
// {{.Title}} is a generated struct representing the Sophos {{.Title}} Endpoint
// GET {{.Path}}
{{if eq .Bytes ""}}type {{.Title}} struct {
	{{range .SubTypes}}{{.Name}} {{.Name}} ` + "`json:\"{{.JsonTag}}\"`" + `
	{{end}}
}{{else}}{{.Bytes}}{{end}}

var defs{{.Title}} =  map[string]sophos.RestObject{
		{{range .SubTypes}}"{{.Name}}": &{{.Name}}{},
		{{end}}
	}

// RestObjects implements the sophos.Node interface and returns a map of {{.Title}}'s Objects
func({{.Title}}) RestObjects() map[string]sophos.RestObject {
	return defs{{.Title}}
}

// GetPath implements sophos.RestGetter
func(*{{.Title}}) GetPath() string { return "{{.Path}}" }
// RefRequired implements sophos.RestGetter
func(*{{.Title}}) RefRequired() (string, bool) { return "", false }

var def{{.Title}} = &sophos.Definition{Description: "{{.Definition.Description}}",Name: "{{.Definition.Name}}",Link: "{{.Definition.Link}}"}

// Definition returns the /api/definitions struct of {{.Title}}
func({{.Title}}) Definition() sophos.Definition { return *def{{.Title}} }

// ApiRoutes returns all known {{.Title}} Paths
func({{.Title}}) ApiRoutes() []string{
	return []string{
		{{range $i, $element := .Routes}}"/api{{$element}}",
		{{end}}
	}
}

// References returns the {{.Title}}'s references. These strings serve no purpose other than to demonstrate which 
// Reference keys are used for this Endpoint
func({{.Title}}) References() []string {
	return []string{
		{{range $i, $element := .References}}"REF_{{$element}}",
		{{end}}
	}
}

{{range .SubTypes}}
// {{.Name}}{{if .IsPlural}}s{{end}} is an Sophos Endpoint subType and implements sophos.RestObject
{{.Bytes}}

{{if .IsPlural}}

// GetPath implements sophos.RestObject and returns the {{.Name}}s GET path{{getDesc . .GetPath "get"}}
func(*{{.Name}}s) GetPath() string { return "/api{{.GetPath}}" }
// RefRequired implements sophos.RestObject
func(*{{.Name}}s) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the {{.Name}}s GET path{{getDesc . .PatchPath "get"}}
func({{firstLetter .Name}} *{{.Name}}) GetPath() string { return fmt.Sprintf("/api{{asRefUrl .PatchPath}}", {{firstLetter .Name}}.Reference) }
// RefRequired implements sophos.RestObject
func({{firstLetter .Name}} *{{.Name}}) RefRequired() (string, bool) { return {{firstLetter .Name}}.Reference, true }

{{else}}

// GetPath implements sophos.RestObject and returns the {{.Name}} GET path{{getDesc . .GetPath "get"}}
func(*{{.Name}}) GetPath() string {	return "/api{{.GetPath}}" }
// RefRequired implements sophos.RestObject
func(*{{.Name}}) RefRequired() (string, bool) { return "", false }

{{end}}

// DeletePath implements sophos.RestObject and returns the {{.Name}} DELETE path{{getDesc . .DeletePath "delete"}}
func(*{{.Name}}) DeletePath(ref string) string {
	return fmt.Sprintf("/api{{asRefUrl .DeletePath}}", ref)
}

// PatchPath implements sophos.RestObject and returns the {{.Name}} PATCH path{{getDesc . .PatchPath "patch"}}
func(*{{.Name}}) PatchPath(ref string) string {
	return fmt.Sprintf("/api{{asRefUrl .PatchPath}}", ref)
}

// PostPath implements sophos.RestObject and returns the {{.Name}} POST path{{getDesc . .PostPath "post"}}
func(*{{.Name}}) PostPath() string {
	return "/api{{.PostPath}}"
}

// PutPath implements sophos.RestObject and returns the {{.Name}} PUT path{{getDesc . .PutPath "put"}}
func(*{{.Name}}) PutPath(ref string) string {
	return fmt.Sprintf("/api{{asRefUrl .PutPath}}", ref)
}

{{if .IsType}}
// GetType implements sophos.Object
func({{firstLetter .Name}} *{{.Name}}) GetType() string { return {{firstLetter .Name}}._type }
{{end}}

{{end}}
`

var funcMap = template.FuncMap{
	"asRefUrl": func(path string) string { return strings.Replace(path, "{ref}", "%s", -1) },
	"getDesc": func(s *subtype, p, m string) string {
		v := s.Node.Paths[p][m].Description
		if strings.TrimSpace(v) == "" {
			return ""
		}
		return fmt.Sprintf("\n// %s", v)
	},
	"firstLetter": func(name string) string { return strings.ToLower(name)[0:1] },
	"asSwag": func(swag *swag) string {
		a := strings.Replace(fmt.Sprintf("%#v", swag.Paths), "main.methodMap", "sophos.MethodMap", -1)
		a = strings.Replace(a, "main.methodDescriptions", "sophos.MethodDescriptions", -1)
		a = strings.Replace(a, "main.parameter", "sophos.Parameter", -1)
		return a
	},
}

func (n *node) ExecuteTemplate(w io.Writer) error {
	tmpl, err := template.New("").Funcs(funcMap).Parse(nodeTemplate)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, n)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
