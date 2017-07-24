package main

import (
	"net"
	"bufio"
	"log"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8080") // start tcp listener
	if err != nil {
		log.Panic(err)
	}
	defer li.Close() // close our listener after we are done

	// accept every connection that is coming in
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn) // for every connection, start a handler in a goroutine
	}
}

// handle will print out the connection info of a request
func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
}