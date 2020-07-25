package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	ppl := map[string]int{
		"Ravi": 35,
		"Rene": 33,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "passData.gohtml", ppl)
	if err != nil {
		log.Fatalln(err)
	}
}
