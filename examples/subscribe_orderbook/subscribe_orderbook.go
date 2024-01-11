package main

import (
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg"
	"github.com/aori-io/aori-sdk-go/pkg/types"
)

func main() {
	// TODO not working
	bot, err := pkg.NewAoriProvider()
	if err != nil {
		fmt.Println("error initializing bot: ", err)
		return
	}

	auth, err := bot.AuthWallet()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received response:", auth)

	response, err := bot.SubscribeOrderbook()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received response:", response)

	err = bot.OnSubscribe(types.OrderCreated, func(payload []byte) {
		fmt.Print("new order: ", string(payload))
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
