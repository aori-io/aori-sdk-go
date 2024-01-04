package ethers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
)

func TestPersonalSign(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("error generating private key: %v", err)
	}

	tests := []struct {
		name    string
		message string
	}{
		{"ShortMessage", "hello"},
		{"LongMessage", "this is a long message to be signed"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signature, err := PersonalSign(tt.message, privateKey)
			if err != nil {
				t.Errorf("error signing message: %v", err)
			}
			if signature == "" {
				t.Error("empty signature returned")
			}
		})
	}
}
