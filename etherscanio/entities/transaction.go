package entities

import (
	"strconv"

	"github.com/pkg/errors"
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
func (t *Transaction) CalculateCost() (gas int64, value int64, err error) {
	gasAmount, err := strconv.ParseInt(t.Gas[2:], 16, 64)
	if err != nil {
		return 0, 0, errors.WithStack(err)
	}

	gasPrice, err := strconv.ParseInt(t.GasPrice[2:], 16, 64)
	if err != nil {
		return 0, 0, errors.WithStack(err)
	}

	value, err = strconv.ParseInt(t.Value[2:], 16, 64)
	if err != nil {
		return 0, 0, errors.WithStack(err)
	}

	return gasAmount * gasPrice, value, nil
}
