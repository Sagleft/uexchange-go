package main

import (
	"log"
	"os"

	uexchange "github.com/Sagleft/uexchange-go"
)

func main() {
	// create client
	client := uexchange.NewClient()

	// auth
	_, err := client.Auth(uexchange.Credentials{
		AccountPublicKey: os.Getenv("PUBLIC_KEY"),
		Password:         os.Getenv("PASSWORD"),
	})
	if err != nil {
		log.Fatalln(err)
	}

	// get balance
	balanceData, err := client.GetBalance()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(balanceData)
}
