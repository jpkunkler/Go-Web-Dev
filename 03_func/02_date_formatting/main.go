package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

// monthDayYear takes in a time.Time object and returns a string in MONTH-DAY-YEAR format
func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

// create our FuncMap
var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
	// we can use func expressions in here
	"fdateKitchen": func(t time.Time) string { return t.Format(time.Kitchen) },
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
