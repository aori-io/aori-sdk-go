package main

import (
	"encoding/json"
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg"
	"github.com/aori-io/aori-sdk-go/pkg/types"
	"log"
)

func main() {
	devReq := "wss://dev.api.aori.io/"
	prodFeed := "wss://feed.aori.io/"

	bot, err := pkg.NewAoriProviderWithURL(devReq, prodFeed)
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

	err = bot.OnSubscribe(types.OrderCreated, func(payload []byte) error {
		var order types.SubscribeOrderViewResponse
		err := json.Unmarshal(payload, &order)
		if err != nil {
			log.Fatalln("failed unmarshalling", err)
			return err
		}

		fmt.Println()
		fmt.Println("new update: ", order.Result.Data.OrderHash)
		fmt.Println()

		// Do more stuff with order

		return nil
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
