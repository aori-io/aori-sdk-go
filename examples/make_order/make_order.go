package main

import (
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg"
	"github.com/aori-io/aori-sdk-go/pkg/types"
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

	orderParams := types.MakeOrderInput{
		SellToken:  "0xD3664B5e72B46eaba722aB6f43c22dBF40181954",
		SellAmount: "100000000", // 100 usdc (6 decimals)
		BuyToken:   "0x2715Ccea428F8c7694f7e78B2C89cb454c5F7294",
		BuyAmount:  "750000000000000000", // 0.75 eth (18 decimals)
	}
	response, err := bot.MakeOrder(orderParams)
	// or private:
	//response, err := bot.MakePrivateOrder(orderParams)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received response:", response)
}
