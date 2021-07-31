package entities

// AnswerBlockInfo describes the answer from the block info endpoint
type AnswerBlockInfo struct {
	Result *Block `json:"result"`
}
