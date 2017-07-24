package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second)) // time out after 10 seconds! Connection will close
	if err != nil {
		log.Println("Connection TimeOut")
	}
	
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln) // print to terminal
		fmt.Fprintf(conn, "I heard you say: %s\n", ln) // print back to our connection
	}
	defer conn.Close()
}