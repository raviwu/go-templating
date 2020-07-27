package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqRoot,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func double(i int) int {
	return 2 * i
}

func square(i int) float64 {
	return math.Pow(float64(i), 2)
}

func sqRoot(i float64) float64 {
	return math.Sqrt(i)
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

	ravi := struct {
		Name string
		Age  int
	}{
		Name: "Ravi",
		Age:  35,
	}
	err = tpl.ExecuteTemplate(os.Stdout, "passStruct.gohtml", ravi)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "passPipeline.gohtml", 32)
	if err != nil {
		log.Fatalln(err)
	}
}
