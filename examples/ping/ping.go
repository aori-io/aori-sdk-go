package main

import (
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg"
)

func main() {
	bot, err := pkg.NewAoriProvider()
	if err != nil {
		fmt.Println("error initializing bot")
		return
	}

	response, err := bot.Ping()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received response:", string(response))
}
