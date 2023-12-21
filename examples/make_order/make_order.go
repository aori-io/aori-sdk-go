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

	orderParams := pkg.MakeOrderInput{
		SellToken:  "0xD3664B5e72B46eaba722aB6f43c22dBF40181954",
		SellAmount: "1000000000000000000",
		BuyToken:   "0x2715Ccea428F8c7694f7e78B2C89cb454c5F7294",
		BuyAmount:  "1000000",
	}
	response, err := bot.MakeOrder(orderParams)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received response:", response)
}
