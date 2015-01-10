package main

// How to use HTML templates in Go.
// Mostly from reading the source here:
// https://github.com/golang/gddo/tree/master/gddo-server

import (
	"fmt"
	"os"
	"html/template"
	"strings"
)

type tmplVars struct {
	Count int
	Names []string
}

func main() {
	fmt.Println("Parsing templates...")
	tvars := tmplVars{Count: 10, Names: []string{"a", "b", "c"}}

	// User defined functions
	funcMap := template.FuncMap {
		// Apparently, the comma here is required.
		"upper": strings.ToUpper,
	}

	// The name argument to template.New can be anything,
	// including an already defined name (such as "ROOT").
	t := template.New("zero").Funcs(funcMap)
	// These can be in any order. They need to contain all the
	// defined templates.
	t, err := t.ParseFiles(
		// Paths can be relative
		"t/index.html",
		"t/root.html",
		"t/common.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t = t.Lookup("ROOT")
	if t == nil {
		fmt.Println("t is nil")
		return
	}
	err = t.Execute(os.Stdout, tvars)
	if err != nil {
		fmt.Println(err)
		return
	}
}
