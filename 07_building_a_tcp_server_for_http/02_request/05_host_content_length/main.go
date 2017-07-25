package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// try parsing whatever was entered into the form
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method        string
		URL           *url.URL
		Submissions   url.Values
		Header        http.Header
		ContentLength int64
		Host          string
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.ContentLength,
		req.Host,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
