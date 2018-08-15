package main

import (
	"fmt"
	"net"
)

func main() {
	/*
		Read from a connection
	*/
	// conn, err := net.Dial("tcp", "localhost:8080")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer conn.Close()

	// bs, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(bs))

	/*
		Write from a connection
	*/
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "I dialed you")
}
