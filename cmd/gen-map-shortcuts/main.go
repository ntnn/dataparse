package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

var (
	fOutput = flag.String("output", "mapShortcuts.gen.go", "")
)

func main() {
	if err := doMain(); err != nil {
		log.Fatal(err)
	}
}

type helperData struct {
	Name         string
	Datatype     string
	DefaultValue string
}

var helperTemplate = `
// {{.Name}} is a shortcut to retrieve a value and call a function on
// the resulting Value.
//
// Calling this method is equivalent to:
//
//	val, err := m.Get("a")
//	if err != nil {
//		// error handling
//	}
//	parsed, err := val.{{.Name}}()
//	if err != nil {
//		// error handling
//	}
func (m Map) {{.Name}}(keys ...any) ({{.Datatype}}, error) {
	v, err := m.Get(keys...)
	if err != nil {
		return {{.DefaultValue}}, err
	}
	return v.{{.Name}}()
}

// Must{{.Name}} is the error-ignoring version of {{.Name}}.
func (m Map) Must{{.Name}}(keys ...any) {{.Datatype}} {
	v, _ := m.{{.Name}}(keys...)
	return v
}
`

var data = []helperData{
	{
		Name:         "Int",
		Datatype:     "int",
		DefaultValue: "0",
	},
	{
		Name:         "Int64",
		Datatype:     "int64",
		DefaultValue: "0",
	},
	{
		Name:         "Uint",
		Datatype:     "uint",
		DefaultValue: "0",
	},
	{
		Name:         "Uint64",
		Datatype:     "uint64",
		DefaultValue: "0",
	},
	{
		Name:         "String",
		Datatype:     "string",
		DefaultValue: "\"\"",
	},
	{
		Name:         "Time",
		Datatype:     "time.Time",
		DefaultValue: "time.Time{}",
	},
}

func doMain() error {
	t, err := template.New("").Parse(helperTemplate)
	if err != nil {
		return err
	}

	out, err := os.Create(*fOutput)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := out.WriteString(`package dataparse

import (
	"time"
)
`); err != nil {
		return err
	}

	for _, d := range data {
		if err := t.Execute(out, d); err != nil {
			return err
		}
	}

	return nil
}
