package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // "dial in" to our TCP server
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn) // read everything the TCP server prints
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs)) // convert byteslice to string and print
}

/* 
Use in combination with 01_write tcp server 
*/