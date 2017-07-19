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
		"the belief of no beliefs.",
	}
	jesus := sage{
		"Jesus",
		"Love all.",
	}
	gandhi := sage{
		"Gandhi",
		"Be the change.",
	}
	muhammad := sage{
		"Muhammad",
		"To overcome evil with good is good, to resist evil by evil is evil.",
	}

	sages := []sage{buddha, jesus, gandhi, muhammad}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
