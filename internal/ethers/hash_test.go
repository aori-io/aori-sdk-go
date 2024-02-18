package ethers

import (
	"github.com/aori-io/aori-sdk-go/pkg/types"
	"testing"
)

func TestGetOrderHash(t *testing.T) {
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
			name: "",
			args: args{order: types.OrderParameters{
				Offerer:       "",
				InputToken:    "",
				InputAmount:   "",
				InputChainID:  0,
				InputZone:     "",
				OutputToken:   "",
				OutputAmount:  "",
				OutputChainID: 0,
				OutputZone:    "",
				StartTime:     "",
				EndTime:       "",
				Salt:          "",
				Counter:       0,
				ToWithdraw:    false,
			}},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignOrderHash(tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignOrderHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SignOrderHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}
