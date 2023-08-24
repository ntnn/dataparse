package main

import (
	_ "embed"
	"flag"
	"log"
	"os"
	"strings"
	"text/template"
)

var (
	fOutput = flag.String("output", "value_numbers_gen.go", "File to write to")
)

func main() {
	flag.Parse()
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

func makeTemplate(text string) (*template.Template, error) {
	tmpl := template.New("")
	tmpl.Funcs(template.FuncMap{
		"hasPrefix": strings.HasPrefix,
	})

	tmpl, err := tmpl.Parse(text)
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
	ret := []string{}
	for _, member := range numberTypes {
		if member != nd.Datatype() {
			ret = append(ret, member)
		}
	}
	return ret
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

//go:embed value_numbers.gotmpl
var numbersTmpl string

func numbers() error {
	tmpl, err := makeTemplate(numbersTmpl)
	if err != nil {
		return err
	}

	gen, err := os.Create(*fOutput)
	if err != nil {
		return err
	}
	defer gen.Close()

	if _, err := gen.WriteString(`package dataparse

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
)

`); err != nil {
		return err
	}

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
		if err := tmpl.Execute(gen, data); err != nil {
			return err
		}
		if _, err := gen.WriteString("\n"); err != nil {
			return err
		}
	}

	return nil
}
