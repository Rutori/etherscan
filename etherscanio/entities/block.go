package entities

type Block struct {
	Number       string         `json:"number"`
	Transactions []*Transaction `json:"transactions"`
}

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
