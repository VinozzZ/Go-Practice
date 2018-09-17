package main

import (
	"log"
	"flag"
	"fmt"
)

// 1. Get User in put from the command-line
// 2. Store Instagram API info in en
var insName = flag.String("insName", "", "Please provide a instagram user name: ")
var ClientID string

func main() 
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	flag.Parse()
	inputUser := *insName
	fmt.Println("insName", inputUser)
}
