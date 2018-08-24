package main

import (
	"io"
	"log"
	"net"
)

func handler(conn net.Conn) {
	io.WriteString(conn, "Server irunning")
	defer conn.Close()
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handler(conn)
	}
}
