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

func CreateAuthWalletPayload(address, signature string) ([]byte, error) {
	req := types.AuthWalletRequest{
		Id:      1,
		JsonRPC: "2.0",
		Method:  "aori_authWallet",
		Params:  []types.AuthWalletParams{{Address: address, Signature: signature}},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateCheckAuthPayload(jwt string) ([]byte, error) {
	req := types.CheckAuthRequest{
		Id:      1,
		JsonRPC: "2.0",
		Method:  "aori_checkAuth",
		Params:  []types.CheckAuthParams{{Auth: jwt}},
	}

	b, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	return b, nil
}
