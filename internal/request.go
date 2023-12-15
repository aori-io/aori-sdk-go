package internal

import (
	"encoding/json"
	"github.com/aori-io/aori-sdk-go/internal/types"
)

func CreatePingPayload() ([]byte, error) {
	req := types.PingRequest{
		Id:      1,
		JsonRPC: "2.0",
		Method:  "aori_ping",
		Params:  []string{},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	return b, nil
}
