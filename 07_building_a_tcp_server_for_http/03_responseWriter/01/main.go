package main

import (
	"net/http"
	"fmt"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Set our custom Headers like this
	w.Header().Set("Kunkler-key", "this key is from jpkunkler")
	w.Header().Set("content-type", "text/html")

	// print to ResponseWriter
	fmt.Fprint(w, "<h1>Any text you want </h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
