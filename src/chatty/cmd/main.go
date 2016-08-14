package main

import (
	"chatty"
	"flag"
	"log"
)

func main() {
	log.Printf("Starting bot\n")

	var token = flag.String("token", "", "integration token for your bot, required")
	flag.Parse()
	if *token == "" {
		log.Printf("\nError, -token required\n\n")
		return
	}

	c := chatty.NewConnection(*token, "#dar")
	err := c.Connect()
	if err != nil {
		log.Printf("Error from c.Connect %v\n", err)
		return
	}
	defer c.Close()

	c.Run()

	log.Printf("Done\n")
}
