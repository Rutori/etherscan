package entities

import "math/big"

// Block describes etherium block data
type Block struct {
	Transactions []*Transaction `json:"transactions"`
}

// GetSpending returns a map that contains addresses and their balance change in this block
func (b *Block) GetSpending() (map[string]*big.Int, error) {
	wallets := make(map[string]*big.Int)
	for i := range b.Transactions {
		if wallets[b.Transactions[i].From] == nil {
			wallets[b.Transactions[i].From] = new(big.Int)
		}

		if wallets[b.Transactions[i].To] == nil {
			wallets[b.Transactions[i].To] = new(big.Int)
		}

		gas, value := b.Transactions[i].CalculateCost()

		wallets[b.Transactions[i].To].Add(wallets[b.Transactions[i].To], value)
		wallets[b.Transactions[i].From].Sub(wallets[b.Transactions[i].From],
			new(big.Int).Add(value, gas),
		)
	}

	return wallets, nil
}
