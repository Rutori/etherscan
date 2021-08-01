package entities

import (
	"math/big"
	"reflect"
	"testing"
)

func TestTransaction_CalculateCost(t1 *testing.T) {
	type fields struct {
		From     string
		Gas      string
		GasPrice string
		To       string
		Value    string
	}
	tests := []struct {
		name      string
		fields    fields
		wantGas   *big.Int
		wantValue *big.Int
	}{
		{
			name: "get gas and value amount",
			fields: fields{
				Gas:      "0x02",
				GasPrice: "0x02",
				Value:    "0x01",
			},
			wantGas:   new(big.Int).SetInt64(4),
			wantValue: new(big.Int).SetInt64(1),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				From:     tt.fields.From,
				Gas:      tt.fields.Gas,
				GasPrice: tt.fields.GasPrice,
				To:       tt.fields.To,
				Value:    tt.fields.Value,
			}
			gotGas, gotValue := t.CalculateCost()
			if !reflect.DeepEqual(gotGas, tt.wantGas) {
				t1.Errorf("CalculateCost() gotGas = %v, want %v", gotGas, tt.wantGas)
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t1.Errorf("CalculateCost() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
