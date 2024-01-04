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

# Installation

To install the SDK:

```bash
go get github.com/aori/aori-sdk-go
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
