package main

import (
	"encoding/json"
	"fmt"
	"github.com/aori-io/aori-sdk-go/pkg"
)

type PingRequest struct {
	Id      int      `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

func main() {
	bot, err := pkg.NewAoriProvider()
	if err != nil {
		fmt.Println("error initializing bot")
		return
	}

	//bot.StartReceiver() // Start the receiver goroutine

	req := PingRequest{
		Id:      1,
		Jsonrpc: "2.0",
		Method:  "aori_ping",
		Params:  []string{},
	}

	b, err := json.Marshal(&req)

	fmt.Println("msg: ", string(b))

	err = bot.Send(b)
	if err != nil {
		fmt.Println("error sending msg", err)
		return
	}

	response, err := bot.Receive()
	if err != nil {
		fmt.Println("error receiving response:", err)
		return
	}

	fmt.Println("Received response:", string(response))
}
