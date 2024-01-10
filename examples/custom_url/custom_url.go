package main

import (
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg"
)

func main() {
	customReq := "wss://dev.api.aori.io/"
	customFeed := "wss://dev.beta.feed.aori.io/"

	bot, err := pkg.NewAoriProviderWithURL(customReq, customFeed)
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
