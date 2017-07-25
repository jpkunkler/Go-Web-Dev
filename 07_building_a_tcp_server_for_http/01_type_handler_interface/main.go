package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// since hotdog now implements the ServeHTTP method it automatically becomes a Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Any code you want can go here")
}

func main() {
	var d hotdog
	//now we can pass in a variable of type d as a handler into ListenAndServe!
	http.ListenAndServe(":8080", d)
}

/*
type Handler interface {
	ServerHTTP(ResponseWriter, *Request)
}

Any type that implements the ServeHTTP method is also a handler!
*/
