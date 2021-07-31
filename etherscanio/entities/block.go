package entities

// Block describes etherium block data
type Block struct {
	Transactions []*Transaction `json:"transactions"`
}

// GetSpending returns a map that contains addresses and their balance change in this block
func (b *Block) GetSpending() (map[string]int64, error) {
	wallets := make(map[string]int64)
	for i := range b.Transactions {
		gas, value, err := b.Transactions[i].CalculateCost()
		if err != nil {
			return nil, err
		}

		wallets[b.Transactions[i].To] += value
		wallets[b.Transactions[i].From] -= value + gas
	}

	return wallets, nil
}
