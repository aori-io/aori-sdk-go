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

	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoiQXBpS2V5IiwiYWRkciI6IjB4dGVzdGFkZHJlc3MiLCJuYW1lIjoiQnJ1Y2UgV2F5bmUiLCJyYXRlTGltaXQiOjEwMDAsImlhdCI6MTY5NDE4NTc3NX0.RZPdn5OQ2d967hztveEVj5l2tCftYMenyXAYKWb2JRE"
	orderId := "0x59eba772fac0c9d9f767ef1d3147f44b578801c1bef51169d8986cac2c32031f"

	response, err := bot.CancelOrder(orderId, jwt)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received response:", response)
}
