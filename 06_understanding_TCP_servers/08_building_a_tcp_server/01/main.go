package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0] // get first token of our request, which is GET or POST
			uri := strings.Fields(ln)[1] // get the uri
			fmt.Println("***METHOD", m)
			fmt.Println("***ROUTE", uri)
		}
		if ln == "" { // an empty string indicates the end of our header
			break
		}
		i++
	}
}

func respond(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	// these are determined by the IEF and their guideline
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") // Status Code
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) // Length of content
	fmt.Fprint(conn, "Content-Type: text/html\r\n") // Content Type
	fmt.Fprint(conn, "\r\n") // an empty line to indicate end of the header
	fmt.Fprint(conn, body) //finally, our html body (CONTENT)
}