package scenarios

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_addAmounts(t *testing.T) {
	type args struct {
		everyone map[string]*big.Int
		delta    map[string]*big.Int
	}
	tests := []struct {
		name    string
		args    args
		wantMap map[string]*big.Int
	}{
		{
			name: "add amounts to the already existing map",
			args: args{
				everyone: map[string]*big.Int{
					"first": new(big.Int).SetInt64(1),
				},
				delta: map[string]*big.Int{
					"first": new(big.Int).SetInt64(1),
				},
			},
			wantMap: map[string]*big.Int{
				"first": new(big.Int).SetInt64(2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addAmounts(tt.args.everyone, tt.args.delta)
			if !reflect.DeepEqual(tt.args.everyone, tt.wantMap) {
				t.Errorf("addAmounts() got map = %v, want map %v", tt.args.everyone, tt.wantMap)
			}
		})
	}
}
