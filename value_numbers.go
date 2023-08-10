//go:build gen

package main

import (
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/ntnn/dataparse"
)

func main() {
	if err := doMain(); err != nil {
		log.Fatal(err)
	}
}

func doMain() error {
	if err := numbers(); err != nil {
		return err
	}

	return nil
}

func makeTemplate(path string) (*template.Template, error) {
	tmpl := template.New(path)
	tmpl.Funcs(template.FuncMap{
		"hasPrefix": strings.HasPrefix,
	})

	tmpl, err := tmpl.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

var numberTypes = []string{
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"float32",
	"float64",
}

type numberData struct {
	Name     string
	datatype string
	def      string
	Bitsize  int
}

func (nd numberData) NativeConverts() []string {
	return dataparse.FilterSlice(numberTypes, nd.Datatype())
}

func (nd numberData) Datatype() string {
	if nd.datatype != "" {
		return nd.datatype
	}
	return strings.ToLower(nd.Name)
}

func (nd numberData) Default() string {
	if nd.def != "" {
		return nd.def
	}
	return "0"
}

func numbers() error {
	tmpl, err := makeTemplate("value_numbers.tmpl")
	if err != nil {
		return err
	}

	f, err := os.Create("value_numbers.gen.go")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("package dataparse\n\n")
	f.WriteString("import (\n")
	f.WriteString("	\"fmt\"\n")
	f.WriteString("	\"strconv\"\n")
	f.WriteString(")\n\n")

	for _, data := range []numberData{
		{
			Name:    "Int",
			Bitsize: 64,
		},
		{
			Name:    "Int8",
			Bitsize: 8,
		},
		{
			Name:    "Int16",
			Bitsize: 16,
		},
		{
			Name:    "Int32",
			Bitsize: 32,
		},
		{
			Name:    "Int64",
			Bitsize: 64,
		},
		{
			Name:    "Uint",
			Bitsize: 64,
		},
		{
			Name:    "Uint8",
			Bitsize: 8,
		},
		{
			Name:    "Uint16",
			Bitsize: 16,
		},
		{
			Name:    "Uint32",
			Bitsize: 32,
		},
		{
			Name:    "Uint64",
			Bitsize: 64,
		},
		{
			Name:    "Float32",
			Bitsize: 32,
		},
		{
			Name:    "Float64",
			Bitsize: 64,
		},
	} {
		if err := tmpl.Execute(f, data); err != nil {
			return err
		}
		f.WriteString("\n")
	}

	return nil
}
