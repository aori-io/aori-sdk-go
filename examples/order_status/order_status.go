package main

import (
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg"
)

func main() {
	// TODO not working
	bot, err := pkg.NewAoriProvider()
	if err != nil {
		fmt.Println("error initializing bot:", err)
		return
	}

	auth, err := bot.AuthWallet()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received response:", auth)

	hash := "0x59eba772fac0c9d9f767ef1d3147f44b578801c1bef51169d8986cac2c32031f"

	response, err := bot.OrderStatus(hash)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received response:", response)
}
