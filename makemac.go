package main

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"
)

//go:embed template.mac
var tmplBody string

func MakeMac(p *Project) (string, error) {
	fm := template.FuncMap{
		"str": func(s string) string {
			return strings.ReplaceAll(s, "'", "''")
		},
		"inc": func(i int) int {
			return i + 1
		},
		"dbset": func(s string) bool {
			return strings.HasPrefix(s, "/")
		},
	}
	tmpl, err := template.New("main").Funcs(fm).Parse(tmplBody)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, p); err != nil {
		return "", err
	}
	return buf.String(), nil
}
