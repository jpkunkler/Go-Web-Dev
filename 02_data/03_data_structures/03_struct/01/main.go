package main

import "text/template"
import "os"
import "log"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

type sage struct {
	Name  string
	Motto string
}

func main() {

	buddha := sage{
		"Buddha",
		"the belief of no beliefs",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
