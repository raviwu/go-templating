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

type person struct {
	Name string
	Age  int
}

func main() {
	ppl := map[string]int{
		"Ravi": 35,
		"Rene": 33,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "passMap.gohtml", ppl)
	if err != nil {
		log.Fatalln(err)
	}

	ravi := person{
		Name: "Ravi",
		Age:  35,
	}
	err = tpl.ExecuteTemplate(os.Stdout, "passStruct.gohtml", ravi)
	if err != nil {
		log.Fatalln(err)
	}
}
