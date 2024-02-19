package ethers

import (
	"github.com/aori-io/aori-sdk-go/pkg/types"
	"testing"
)

func TestCalculateOrderHash(t *testing.T) {
	type args struct {
		order types.OrderParameters
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "default order",
			args: args{order: types.OrderParameters{
				Offerer:       "0x000061444E91C5e75de8dCa7BD6C4dC406c1FA56",
				InputToken:    "0x2222222222222222222222222222222222222222",
				InputAmount:   "2000000000000000000", // 1 token in wei
				InputChainID:  5,
				InputZone:     "0xCB93Ed64a1b4C61b809F556532Ff36EA25DaD473",
				OutputToken:   "0x4444444444444444444444444444444444444444",
				OutputAmount:  "2000000000000000000", // 2 tokens in wei
				OutputChainID: 5,
				OutputZone:    "0xCB93Ed64a1b4C61b809F556532Ff36EA25DaD473",
				StartTime:     "1622505600",
				EndTime:       "1725107600",
				Salt:          "12345678",
				Counter:       1,
				ToWithdraw:    false,
			}},
			want:    "0x74b2f45fdff335c40d476d0979bb91c6bf573fa58d122326e604dda457e5627e",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateOrderHash(tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateOrderHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateOrderHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}
