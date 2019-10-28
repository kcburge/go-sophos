// Program gen is used to generate go-sophos types
package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/ChimeraCoder/gojson"
	"github.com/esurdam/go-sophos"
)

var (
	client  *sophos.Client
	rootDir string
	debug   bool

	numberSequence    = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
	numberReplacement = []byte(`$1 $2 $3`)
	header            = `package objects

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
		Link        string    // path to definition
		Swag        *swag     // the definition from the UTM
		Endpoint    *endpoint // a representation of the Endpoint
	}
	// map[path]map[method]methodDescription
	swag struct {
		Paths map[string]methodMap
		// Definitions are Object definitions
		Definitions map[string]subTypeDef
	}
	subTypeDef struct {
		Properties map[string]struct {
			Type        string
			Enum        []string
			Items       map[string]string
			Default     interface{}
			Description string
		}
		Description string
		Type        string
	}
	methodMap map[string]struct {
		Description string
		Parameters  []struct {
			Name        string
			In          string
			Description string
			Type        string
			Required    bool
		}
		Tags      []string
		Responses map[int]struct{ Description string }
	}
	endpoint struct {
		Definition  *definition `json:"-"` //used in template
		Title, Name string
		Bytes       string
		Routes      []string
		Path        string
		Methods     []string
		SubTypes    []subtype
		References  []string
		Paths       map[string]methodMap `json:"-"`
	}
	subtype struct {
		Name              string
		JsonTag           string
		GetPath           string
		GetPaths          []string
		PutPath           string
		PostPath          string
		Type              subTypeDef
		DeletePath        string
		PatchPath         string
		HasRef            bool
		Bytes             string
		Node              *endpoint `json:"-"`
		IsPlural          bool
		IsPluralInterface bool
		IsType            bool
	}
)

func main() {
	var err error

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client, err = sophos.New(os.Getenv("TF_VAR_utm_api_endpoint"), sophos.WithBasicAuth(os.Getenv("TF_VAR_utm_api_user"), os.Getenv("TF_VAR_utm_api_password")))
	if err != nil {
		log.Fatal(err)
	}

	// TODD: version api against UTM
	var v sophos.Version
	r, err := client.Get("/api/status/version")
	if err != nil {
		log.Fatal(err)
	}
	r.MarshalTo(&v)
	fmt.Println(v.Restd)

	rootDir = "api/v" + v.Restd
	os.RemoveAll(rootDir)

	err = os.MkdirAll(rootDir, 0777)
	if err != nil {
		log.Fatal(err)
	}

	var dd []definition
	r, err = client.Get("/api/definitions")
	if err != nil {
		log.Fatal(err)
	}

	err = r.MarshalTo(&dd)
	if err != nil {
		log.Fatal(err)
	}

	var d int
	for _, def := range dd {
		d++
		err := def.process()
		if err != nil {
			log.Fatal(err)
		}

		subDir := rootDir + "/objects/"
		os.MkdirAll(subDir, 0777)
		f, err := os.Create(subDir + strings.ToLower(def.Name) + ".go")
		if err != nil {
			log.Fatal(err)
		}
		if d == 1 {
			f.Write([]byte(`// Package objects contains the generated Sophos object types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
`))
		}

		f.Write([]byte(header))

		if debug {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "    ")
			if err := enc.Encode(def.Endpoint); err != nil {
				panic(err)
			}
		}

		fmt.Printf("writing %s\n", def.Name)
		err = def.Endpoint.ExecuteTemplate(f)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

type nftd struct {
	Name, Val, Path string
}

var nodeTypeFuncsTemplate = `
// {{.Name}} represents the {{.Path}} node and implements sophos.Node
type {{.Name}} struct {	Value {{.Val}} }

// Get gets the {{.Path}} value from the UTM
func({{firstLetter .Name}} *{{.Name}})Get(client sophos.ClientInterface, options ...sophos.Option)(err error) {
	return get(client, "/api/nodes/{{.Path}}", &{{firstLetter .Name}}.Value, options...)
}

// Update is syntactic sugar for Update{{.Name}}
func({{firstLetter .Name}} *{{.Name}})Update(client sophos.ClientInterface, options ...sophos.Option)(err error) {
	return put(client, "/api/nodes/{{.Path}}", {{firstLetter .Name}}.Value, options...)
}
`

var nodeFuncsTemplate = `
// Get{{.Name}} gets the {{.Path}} value from the UTM
func Get{{.Name}}(client sophos.ClientInterface, options ...sophos.Option) (val {{.Val}}, err error) {
	err = get(client, "/api/nodes/{{.Path}}", &val, options...)
	return
}

// Update{{.Name}} PUTs the {{.Path}} value to the UTM
func Update{{.Name}}(client sophos.ClientInterface, val {{.Val}}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/{{.Path}}", val, options...)
}
`

var nodeFuncsTestTemplate = `
func TestGet{{.Name}}(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	_, err := Get{{.Name}}(client)
	if err != nil {
		t.Errorf("TestGet{{.Name}} should not have error: %s", err.Error())
	}
}

func TestUpdate{{.Name}}(t *testing.T) {
	td := setupTestCase(t)
	defer td(t)

	var v {{.Val}}
	err := Update{{.Name}}(client, v)
	if err != nil {
		t.Error(err.Error())
	}
}
`

func executeTmpl(f io.Writer, v string, data interface{}) {
	tmpl, err := template.New("").Parse(v)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

var nodesHeader = `package nodes

import "github.com/esurdam/go-sophos"
import "encoding/json"
`

func handleNodesNode() error {
	var nodes map[string]interface{}
	r, err := client.Get("/api/nodes")
	if err != nil {
		log.Fatal(err)
	}
	r.MarshalTo(&nodes)

	keys := make([]string, 0, len(nodes))
	for key := range nodes {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	subDir := rootDir + "/nodes"
	err = os.MkdirAll(subDir, 0777)
	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.Create(subDir + "/nodes.go")
	if err != nil {
		log.Fatal(err)
	}
	f2.Write([]byte(`// Package nodes contains generated types and Get/Update functions for sophos.Node(s)
//
// This file was generated by bin/gen.go! DO NOT EDIT!
` + nodesHeader))
	//
	// for _, key := range keys {
	//	strKey := strings.Replace(key, ".", "_", -1)
	//	// f2.Write([]byte(fmt.Sprintf(" %sNode = sophos.Endpoint(\"%s\")\n", toCamelInitCase(strKey, true), key)))
	//	f2.Write([]byte(fmt.Sprintf("type %s struct{ Path string, Value %s}", toCamelInitCase(strKey, true), key)))
	// }

	for _, key := range keys {
		value := nodes[key]
		valueType := typeForValue(value)
		strKey := strings.Replace(key, ".", "_", -1)
		tmpl, err := template.New("").Funcs(funcMap).Parse(nodeTypeFuncsTemplate)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = tmpl.Execute(f2, &nftd{
			Name: toCamelInitCase(strKey, true),
			Val:  valueType,
			Path: key,
		})
		if err != nil {
			log.Fatal(err.Error())
		}
		// f.Write([]byte(fmt.Sprintf(" %sValue %s\n", toCamelInitCase(strKey, true), valueType)))
	}
	f2.Close()

	f, err := os.Create(subDir + "/handlers.go")
	if err != nil {
		log.Fatal(err)
	}
	// ft, err := os.Create(subDir + "/handlers_test.go")
	// if err != nil {
	//	log.Fatal(err)
	// }
	// ft.Write([]byte(nodesTestTemp))

	f.Write([]byte(nodesHeader))
	f.Write([]byte(`
func get(c sophos.ClientInterface, path string, val interface{}, options ...sophos.Option) (err error) {
	res, err := c.Get(path, options...)
	if err != nil {
		return err
	}
	err = res.MarshalTo(val)
	return
}

func put(c sophos.ClientInterface, path string, val interface{}, options ...sophos.Option) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put(path, bytes.NewReader(byt), options...)
	return
}
`))

	for _, key := range keys {
		value := nodes[key]
		valueType := typeForValue(value)
		strKey := strings.Replace(key, ".", "_", -1)

		executeTmpl(f, nodeFuncsTemplate, &nftd{
			Name: toCamelInitCase(strKey, true),
			Val:  valueType,
			Path: key,
		})

		// executeTmpl(ft, nodeFuncsTestTemplate, &nftd{
		//	Name: toCamelInitCase(strKey, true),
		//	Val:  valueType,
		//	Path: key,
		// })

		// f.Write([]byte(fmt.Sprintf(" %sValue %s\n", toCamelInitCase(strKey, true), valueType)))
	}
	f.Close()
	// ft.Close()

	f3, err := os.Create(subDir + "/directory.go")
	if err != nil {
		log.Fatal(err)
	}
	f3.Write([]byte(nodesHeader))
	f3.Write([]byte(`// Lookup will retrieve a sophos.Node by its name
	func Lookup(name string)sophos.Node{ return nodeDirectory[name] }
	`))
	f3.Write([]byte("var nodeDirectory = map[string]sophos.Node{\n"))
	for _, key := range keys {
		strKey := strings.Replace(key, ".", "_", -1)
		f3.Write([]byte(fmt.Sprintf("\"%s\": &%s{},\n", key, toCamelInitCase(strKey, true))))
	}
	f3.Write([]byte("}"))
	f3.Close()
	return nil
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

	// format the path, see if its a endpoint endpoint
	path := fmt.Sprintf("/api/%s", strings.ToLower(def.Name))
	if def.Name != "Nodes" {
		path = fmt.Sprintf("/api/nodes/%s", strings.ToLower(def.Name))
	}

	// the endpoint will represent this definition
	def.Endpoint = &endpoint{
		Definition: def,
		Title:      toCamelInitCase(def.Name, true),
		Name:       def.Name,
		Path:       path,
	}

	ep := def.Endpoint

	def.Endpoint.fetch()
	ep.Paths = def.Swag.Paths
	if def.Name == "Nodes" {
		return handleNodesNode()
	}

	// Swag.Paths contains a mapping of path -> map[method]methodDescription
	for path, methodMap := range def.Swag.Paths {
		// add the path to the known Endpoint Paths
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
		for method, d := range methodMap {
			// add the method to the known Endpoint methods
			ep.AddMethod(method)
			s := subtype{
				Name:    toCamelInitCase(def.Name+"_"+name, true),
				JsonTag: def.Name + "_" + name,
				Node:    def.Endpoint,
			}

			if s.Name == "StatusStatus" {
				s.Name = "StatusVersion"
			}

			if strings.Contains(path, "{ref}") {
				// HasRef typically means we can de-pluralize the struct returned from client in "get"
				s.HasRef = true
				// change the name to uppercase Ref
				name = strings.Replace(name, "{ref}", "{Ref}", -1)
				// add the name to the known Endpoint references
				ep.AddReference(toCamelInitCase(def.Name+"_"+name, true))
			} else {
				s.GetPath = path
			}

			if method == "get" {
				s.Type = def.Swag.Definitions[strings.Replace(d.Tags[0], "/", ".", -1)]
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

			// Add the subtype to the endpoint
			ep.AddSubType(s)
		}

	}

	sort.Strings(ep.Routes)
	sort.Strings(ep.Methods)
	sort.Strings(ep.References)
	sort.SliceStable(ep.SubTypes, func(i, j int) bool { return ep.SubTypes[i].Name < ep.SubTypes[j].Name })
	return nil
}

// fetch fetches the endpoint itself
func (n *endpoint) fetch() error {
	// get the struct
	r, err := client.Get(n.Path)
	if err != nil {
		// error here is okay since endpoint/endpoint is not a /endpoint/{{Path}} subtype
		// objects wil be retrieved from endpoints
		// log.Println("could not get path")
		return err
	}
	// write the endpoint data
	byt, err := gojson.Generate(r.Body, gojson.ParseJson, n.Title, "main", []string{"json"}, false, true)
	if err != nil {
		log.Printf("could not gojson response: %s, %s\n", n.Path, err.Error())
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

func (n *endpoint) AddReference(ref string) {
	for _, se := range n.References {
		if se == ref {
			return
		}
	}
	n.References = append(n.References, ref)
}

func (n *endpoint) AddMethod(m string) {
	for _, se := range n.Methods {
		if se == m {
			return
		}
	}
	n.Methods = append(n.Methods, m)
}

func (n *endpoint) AddSubType(s subtype) {
	for idx, se := range n.SubTypes {
		if se.Name == s.Name {
			if s.HasRef {
				n.SubTypes[idx].HasRef = true
			}
			if len(s.GetPaths) > 0 {
				n.SubTypes[idx].GetPaths = append(n.SubTypes[idx].GetPaths, s.GetPaths...)
			}

			if s.GetPath != "" {
				n.SubTypes[idx].Type = s.Type
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

			if !se.IsPluralInterface && s.IsPluralInterface {
				n.SubTypes[idx].IsPluralInterface = s.IsPluralInterface
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
	byt, err := gojson.Generate(resp.Body, gojson.ParseJson, name, "main", []string{"json"}, false, true)
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
			if strings.HasSuffix(line, " []interface{}\n") {
				//	// make pluralized
				s.IsPlural = true
				s.IsPluralInterface = true
				newLine := strings.Replace(line, " []interface{}", "s []"+name, 1)
				outBuf.Write([]byte(newLine))
				outBuf.Write([]byte("// " + name + " represents a UTM " + s.Type.Description + "\n"))
				newType := strings.Replace(line, " []interface{}", " struct {", 1)

				// Write the struct from the subtType
				outBuf.Write([]byte(newType))
				outBuf.Write([]byte("Locked string `json:\"_locked\"`\n"))
				outBuf.Write([]byte("ObjectType string `json:\"_type\"`\n"))
				outBuf.Write([]byte("Reference string `json:\"_ref\"`\n"))
				for k, p := range s.Type.Properties {
					if p.Description != "" {
						outBuf.Write([]byte(fmt.Sprintf("// %s description: %s\n", toCamelInitCase(k, true), p.Description)))
					}
					if len(p.Enum) > 0 {
						outBuf.Write([]byte(fmt.Sprintf("// %s can be one of: %#v\n", toCamelInitCase(k, true), p.Enum)))
					}
					switch p.Type {
					case "string":
						if v, ok := p.Default.(string); ok {
							outBuf.Write([]byte(fmt.Sprintf("// %s default value is %#v\n", toCamelInitCase(k, true), v)))
						}
						outBuf.Write([]byte(fmt.Sprintf("%s string `json:\"%s\"`\n", toCamelInitCase(k, true), k)))
					case "integer":
						if v, ok := p.Default.(int); ok {
							outBuf.Write([]byte(fmt.Sprintf("// %s default value is %d\n", toCamelInitCase(k, true), v)))
						}
						outBuf.Write([]byte(fmt.Sprintf("%s int `json:\"%s\"`\n", toCamelInitCase(k, true), k)))
					case "boolean":
						if v, ok := p.Default.(bool); ok {
							outBuf.Write([]byte(fmt.Sprintf("// %s default value is %#v\n", toCamelInitCase(k, true), v)))
						}
						outBuf.Write([]byte(fmt.Sprintf("%s bool `json:\"%s\"`\n", toCamelInitCase(k, true), k)))
					case "array":
						outBuf.Write([]byte(fmt.Sprintf("%s []interface{} `json:\"%s\"`\n", toCamelInitCase(k, true), k)))
					default:
						fmt.Printf("Do not know type \"%s\" for %s: %s\n", p.Type, k, p.Description)
						outBuf.Write([]byte(fmt.Sprintf("%s interface{} `json:\"%s\"`\n", toCamelInitCase(k, true), k)))
					}
				}
				outBuf.Write([]byte("}\n"))
				continue
			}
		}

		if strings.Contains(line, "`json:\"_type\"`") {
			s.IsType = true
			line = strings.Replace(line, "Type", "ObjectType", -1)
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

var endpointTemplate = `
// {{.Title}} is a generated struct representing the Sophos {{.Title}} Endpoint
// GET {{.Path}}
{{if eq .Bytes ""}}type {{.Title}} struct {
	{{range .SubTypes}}{{.Name}} {{.Name}} ` + "`json:\"{{.JsonTag}}\"`" + `
	{{end}}
}{{else}}{{.Bytes}}{{end}}

var _ sophos.Endpoint = &{{.Title}}{}

var defs{{.Title}} =  map[string]sophos.RestObject{
		{{range .SubTypes}}"{{.Name}}": &{{.Name}}{},
		{{end}}
	}

// RestObjects implements the sophos.Node interface and returns a map of {{.Title}}'s Objects
func({{.Title}}) RestObjects() map[string]sophos.RestObject { return defs{{.Title}} }

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

var _ sophos.RestGetter = &{{.Name}}{}
// GetPath implements sophos.RestObject and returns the {{.Name}}s GET path{{getDesc . .GetPath "get"}}
func(*{{.Name}}s) GetPath() string { return "/api{{.GetPath}}" }
// RefRequired implements sophos.RestObject
func(*{{.Name}}s) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the {{.Name}}s GET path{{getDesc . .PatchPath "get"}}
func({{firstLetter .Name}} *{{.Name}}) GetPath() string { return fmt.Sprintf("/api{{asRefUrl .PatchPath}}", {{firstLetter .Name}}.Reference) }
// RefRequired implements sophos.RestObject
func({{firstLetter .Name}} *{{.Name}}) RefRequired() (string, bool) { return {{firstLetter .Name}}.Reference, true }

{{else}}

var _ sophos.RestObject = &{{.Name}}{}
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

// UsedByPath implements sophos.RestObject{{getUsedBy .}}
func(*{{.Name}}) UsedByPath(ref string) string {
	return fmt.Sprintf("/api{{asRefUrl .PutPath}}/usedby", ref)
}

{{if .IsType}}
// GetType implements sophos.Object
func({{firstLetter .Name}} *{{.Name}}) GetType() string { return {{firstLetter .Name}}.ObjectType }
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
	"getUsedBy": func(s *subtype) string {
		v := s.Node.Paths[s.PatchPath+"/usedby"]["get"].Description
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

func (n *endpoint) ExecuteTemplate(w io.Writer) error {
	tmpl, err := template.New("").Funcs(funcMap).Parse(endpointTemplate)
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

// generate an appropriate struct type entry
func typeForValue(value interface{}) string {
	//Check if this is an array
	if objects, ok := value.([]interface{}); ok {
		types := make(map[reflect.Type]bool, 0)
		for _, o := range objects {
			types[reflect.TypeOf(o)] = true
		}
		if len(types) == 1 {
			return "[]" + typeForValue(mergeElements(objects).([]interface{})[0])
		}
		return "[]interface{}"
	} else if _, ok := value.(map[interface{}]interface{}); ok {
		return ""
	} else if _, ok := value.(map[string]interface{}); ok {
		return "map[string]interface{}"
	} else if reflect.TypeOf(value) == nil {
		return "interface{}"
	}
	v := reflect.TypeOf(value).Name()
	if v == "float64" {
		v = disambiguateFloatInt(value)
	}
	return v
}

// All numbers will initially be read as float64
// If the number appears to be an integer value, use int instead
func disambiguateFloatInt(value interface{}) string {
	const epsilon = .0001
	vfloat := value.(float64)
	if math.Abs(vfloat-math.Floor(vfloat+epsilon)) < epsilon {
		var tmp int64
		return reflect.TypeOf(tmp).Name()
	}
	return reflect.TypeOf(value).Name()
}

func mergeElements(i interface{}) interface{} {
	switch i := i.(type) {
	default:
		return i
	case []interface{}:
		l := len(i)
		if l == 0 {
			return i
		}
		for j := 1; j < l; j++ {
			i[0] = mergeObjects(i[0], i[j])
		}
		return i[0:1]
	}
}

func mergeObjects(o1, o2 interface{}) interface{} {
	if o1 == nil {
		return o2
	}

	if o2 == nil {
		return o1
	}

	if reflect.TypeOf(o1) != reflect.TypeOf(o2) {
		return nil
	}

	switch i := o1.(type) {
	default:
		return o1
	case []interface{}:
		if i2, ok := o2.([]interface{}); ok {
			i3 := append(i, i2...)
			return mergeElements(i3)
		}
		return mergeElements(i)
	case map[string]interface{}:
		if i2, ok := o2.(map[string]interface{}); ok {
			for k, v := range i2 {
				if v2, ok := i[k]; ok {
					i[k] = mergeObjects(v2, v)
				} else {
					i[k] = v
				}
			}
		}
		return i
	case map[interface{}]interface{}:
		if i2, ok := o2.(map[interface{}]interface{}); ok {
			for k, v := range i2 {
				if v2, ok := i[k]; ok {
					i[k] = mergeObjects(v2, v)
				} else {
					i[k] = v
				}
			}
		}
		return i
	}
}

var nodesTestTemp = `
package nodes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/esurdam/go-sophos"
)

var client *sophos.Client

var errOption = func(r *http.Request) error {
	return fmt.Errorf("this is a fake error")
}

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/error" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		var v interface{}
		json.NewEncoder(w).Encode(&v)
	}))
	sophos.DefaultHTTPClient = ts.Client()
	clientF, err := sophos.New(ts.URL, sophos.WithAPIToken("abc"))
	if err == nil {
		client = clientF
	}
	if client == nil {
		t.Error("errror setting up client, client is nil")
	}
	return func(t *testing.T) {
		ts.Close()
		t.Log("teardown test case")
	}
}
`
