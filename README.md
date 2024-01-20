# Aori Go SDK

[![Test Status](https://github.com/aori-io/aori-sdk-go/actions/workflows/build.yml/badge.svg)](https://github.com/aori-io/aori-sdk-go/actions)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/aori-io/aori-sdk-go/)
[![Aori Dev Telegram Chat][tg-badge]][tg-link]

[tg-badge]: https://img.shields.io/endpoint?url=https%3A%2F%2Ftg.sumanjay.workers.dev%2F%2Bdvbw0fIyS-llMmI0&logo=telegram&color=neon
[tg-link]: https://t.me/+dvbw0fIyS-llMmI0

![H](assets/aori.svg)

Aori is a high-performance orderbook protocol for high-frequency trading on-chain and facilitating OTC settlement. This repository provides a Golang SDK for interacting with the Aori Websocket-based API to help developers integrate and build on top of the protocol as easily as possible.

This SDK is released under the [MIT License](LICENSE).

---

If you have any further questions, refer to [the technical documentation](https://www.aori.io/developers). Alternatively, please reach out to us [on Discord](https://discord.gg/K37wkh2ZfR) or [on Twitter](https://twitter.com/aori_io).

## Table of Contents

- [Installation](#installation)
  - [Initialization](#initialization)
- [Examples](#examples)
  - [Subscribing to the orderbook](#subscribing-to-the-orderbook)

# Installation

To install the SDK:

```bash
go get github.com/aori-io/aori-sdk-go
```

## Initialization

Ensure you have these environment variables defined:

```bash
WALLET_ADDRESS=0x_your_wallet_address
PRIVATE_KEY=your_private_key (without the 0x prefix)
NODE_URL=(wss node url) // can use 'wss://ethereum-goerli.publicnode.com' for example //
```

Then you can go ahead and start initializing the bot

```go
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

  // ...
}
```

## Examples

Please refer to the ./examples folder to see working examples

### Subscribing to the orderbook

Alternatively, one can utilise the subscription to global orderbook events. By default, this is provided. Relevant events for the updating of the state of public orders will be emitted to allow clients to manage a local view of the orderbook for their own purposes.

All relevant orderbook events are under the enum SubscriptionEvents.

> IMPORTANT TO NOTE: The callback in `OnSubscribe` takes a byte slice as an input, meaning it is your responsibility to Unmarshal it into the correct struct. To help you out, below are examples for all Subscription Events.

```go
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

	response, err := bot.SubscribeOrderbook()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Received response:", response)

	err = bot.OnSubscribe(types.OrderCancelled, func(payload []byte) error {
		var order types.SubscribeOrderViewResponse
		err := json.Unmarshal(payload, &order)
		if err != nil {
			log.Fatalln("failed unmarshalling", err)
			return err
		}

		// Do more stuff with order

		return nil
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// ...
	// ...
	// ...

	err = bot.OnSubscribe(types.OrderCreated, func(payload []byte) error {
		var order types.SubscribeOrderViewResponse
		err := json.Unmarshal(payload, &order)
		if err != nil {
			log.Fatalln("failed unmarshalling", err)
			return err
		}

		// Do more stuff with order

		return nil
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// ...
	// ...
	// ...

	err = bot.OnSubscribe(types.OrderTaken, func(payload []byte) error {
		var order types.SubscribeTakeOrderResponse
		err := json.Unmarshal(payload, &order)
		if err != nil {
			log.Fatalln("failed unmarshalling", err)
			return err
		}

		// Do more stuff with order

		return nil
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// ...
	// ...
	// ...

	err = bot.OnSubscribe(types.OrderFulfilled, func(payload []byte) error {
		var order types.SubscribeFulfilledOrderResponse
		err := json.Unmarshal(payload, &order)
		if err != nil {
			log.Fatalln("failed unmarshalling", err)
			return err
		}

		// Do more stuff with order

		return nil
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// ...
	// ...
	// ...

	err = bot.OnSubscribe(types.QuoteRequested, func(payload []byte) error {
		var order types.SubscribeQuoteRequestedResponse
		err := json.Unmarshal(payload, &order)
		if err != nil {
			log.Fatalln("failed unmarshalling", err)
			return err
		}

		// Do more stuff with order

		return nil
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// ...
	// ...
	// ...

	err = bot.OnSubscribe(types.OrderToExecute, func(payload []byte) error {
		var order types.SubscribeOrderToExecuteResponse
		err := json.Unmarshal(payload, &order)
		if err != nil {
			log.Fatalln("failed unmarshalling", err)
			return err
		}

		// Do more stuff with order

		return nil
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

```
