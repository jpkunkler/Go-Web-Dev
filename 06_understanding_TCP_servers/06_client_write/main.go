package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // "dial in" to our TCP server
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you!") // print to connection (write to our server)
}

/* 
Use in combination with 02_read tcp server 
*/