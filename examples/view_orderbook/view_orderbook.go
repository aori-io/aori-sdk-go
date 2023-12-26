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

	query := types.ViewOrderbookParams{
		ChainId: 5,
		Query: types.ViewOrderbookQuery{
			Base:  "0xD3664B5e72B46eaba722aB6f43c22dBF40181954",
			Quote: "0x2715Ccea428F8c7694f7e78B2C89cb454c5F7294",
		},
	}

	response, err := bot.ViewOrderbook(query)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if len(response.Result.Orders) == 0 {
		fmt.Println("Received response:", response)
	}
	for _, s := range response.Result.Orders {
		fmt.Println("orderHash:", s.OrderHash)
	}
}
