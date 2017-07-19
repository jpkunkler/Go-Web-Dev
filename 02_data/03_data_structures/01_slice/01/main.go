package main

import "text/template"
import "os"
import "log"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {

	sages := []string{"Ghandi", "Jesus", "MLK", "Muhammad"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
