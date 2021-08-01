package entities

import (
	"math/big"
)

// Transaction describes transaction data from etherscan API
type Transaction struct {
	From     string `json:"from"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Input    string `json:"input"`
	To       string `json:"to"`
	Value    string `json:"value"`
}

// CalculateCost returns transaction value and the amount of gas spent
func (t *Transaction) CalculateCost() (gas *big.Int, value *big.Int) {
	gasAmount := new(big.Int)
	gasAmount.SetString(t.Gas[2:], 16)

	gasPrice := new(big.Int)
	gasPrice.SetString(t.GasPrice[2:], 16)

	value = new(big.Int)
	value.SetString(t.Value[2:], 16)

	gas = new(big.Int).Mul(gasPrice, gasAmount)

	return gas, value
}
