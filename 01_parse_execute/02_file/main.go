package main

import "text/template"
import "log"
import "os"

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// create new file index.html
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close()

	// since Execute takes in a Writer, we can use our newly created file nf as such
	// so that the contents of our template will be written to index.html
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
