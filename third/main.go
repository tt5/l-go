package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	funcMap := template.FuncMap{
			"upper": strings.ToUpper,
	}



	const nameTmpl = `{{define "name"}} {{upper .Name}} {{end}}`
	const tmpl = `Hello, {{range . -}}{{template "name" .}}{{end -}}!`

	t, err := template.New("greet").Funcs(funcMap).Parse(nameTmpl)
	if err != nil {
		panic(err)
	}
	t, err = t.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	type Person struct {
			Name  string
	}

	data := []Person{
		{"one"},
		{"two"},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
