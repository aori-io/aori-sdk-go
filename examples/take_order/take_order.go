package main

import (
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg"
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

	orderHash := "0x140c5af4f95b5e3f9f1160cc99da9e6ab5ee73741da3f656ebbf39995b13be19"

	response, err := bot.OrderStatus(orderHash)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	takeResponse, err := bot.TakeOrder(response.Result.Order.Order.Parameters, orderHash, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received response:", takeResponse)
}
