package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// try parsing whatever was entered into the form
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// pass the form (of type map[string][]string) as data into our template
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
