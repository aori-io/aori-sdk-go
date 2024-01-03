package main

import (
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg"
)

func main() {
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

	response, err := bot.AccountOrders(auth.Result.Auth)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received response:", response.ID)
	for _, order := range response.Result.Orders {
		fmt.Println("Order: ", order.OrderHash)
	}
}
