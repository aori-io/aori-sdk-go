package main

import (
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg"
)

func main() {
	customUrl := "wss://dev.api.aori.io/"

	bot, err := pkg.NewAoriProviderWithURL(customUrl)
	if err != nil {
		fmt.Println("error initializing bot: ", err)
		return
	}

	response, err := bot.Ping()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received response:", response)
}
