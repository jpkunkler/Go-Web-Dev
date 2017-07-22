package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080") // Listeon on TCP Protocoll and port 8080
	if err != nil {
		log.Panic(err)
	}
	defer li.Close() // don't forget to close our Listener after we are done!

	for {
		conn, err := li.Accept() // accept incoming connection
		if err != nil {
			log.Println(err)
			continue
		}

		// write to our connected client
		io.WriteString(conn, "\nHey, it's your TCP Server!\n")
		fmt.Fprintln(conn, "How are you today?")
		fmt.Fprintf(conn, "%v", "Well, I hope!\n")

		conn.Close() // close our client connection
	}
}
