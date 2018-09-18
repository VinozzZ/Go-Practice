package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gedex/go-instagram/instagram"
	"github.com/joho/godotenv"
)

// 1. Get User in put from the command-line
// 2. Store Instagram API info in en
var insName = flag.String("insName", "", "Please provide a instagram user name: ")
var ClientID string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ClientID = os.Getenv("CLIENT_ID")
	flag.Parse()
	// inputUser := *insName
	client := instagram.NewClient(nil)
	client.ClientID = ClientID
	media, _, _ := client.Users.RecentMedia("357679952", nil)
	fmt.Println("media", media)
}
