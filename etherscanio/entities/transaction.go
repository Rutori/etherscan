package entities

import (
	"strconv"

	"github.com/pkg/errors"
)

type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	Type             string `json:"type"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

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
