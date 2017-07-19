package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {

	nf, err := os.Create("default.html")
	if err != nil {
		log.Println(err)
	}

	err = tpl.ExecuteTemplate(nf, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
