package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handler(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner()
	for scanner.Scan() {
		ln, err := scanner.Text()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(ln)
	}

}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go handler(conn)
	}
}
