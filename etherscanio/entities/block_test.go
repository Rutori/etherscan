package entities

import (
	"math/big"
	"reflect"
	"testing"
)

func TestBlock_GetSpending(t *testing.T) {
	type fields struct {
		Transactions []*Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]*big.Int
		wantErr bool
	}{
		{
			name: "calculate the correct change in balance for wallets in transactions",
			fields: fields{
				Transactions: []*Transaction{
					{
						From:     "first",
						Gas:      "0x01",
						GasPrice: "0xfa",
						To:       "second",
						Value:    "0xaa",
					},
					{
						From:     "second",
						Gas:      "0x01",
						GasPrice: "0xfa",
						To:       "first",
						Value:    "0x1b",
					},
				},
			},
			want: map[string]*big.Int{
				"first":  new(big.Int).SetInt64(-393),
				"second": new(big.Int).SetInt64(-107),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Block{
				Transactions: tt.fields.Transactions,
			}
			got, err := b.GetSpending()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpending() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSpending() got = %v, want %v", got, tt.want)
			}
		})
	}
}
