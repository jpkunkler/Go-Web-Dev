package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	http.HandleFunc("/dog/", d) // the trailing backslash catches everything, i.e. /dog/something/test
	http.HandleFunc("/cat", c) // without the backslash, only an exact match will be routed, else 404 Error

	http.ListenAndServe(":8080", nil)
}