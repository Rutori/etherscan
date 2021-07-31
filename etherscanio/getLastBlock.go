package etherscanio

import (
	"context"
	"encoding/json"

	"etherscan_parse/etherscanio/client"
	"etherscan_parse/etherscanio/entities"
	"github.com/pkg/errors"
)

const lastBlockEndpoint = "https://api.etherscan.io/api?module=proxy&action=eth_blockNumber"

// GetLastBlock returns last block tag
func GetLastBlock(ctx context.Context, api *client.API) (string, error) {
	data, err := api.Query(ctx, lastBlockEndpoint)
	if err != nil {
		return "", err
	}

	block := entities.LastBlock{}

	err = json.Unmarshal(data, &block)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return block.Result, nil
}
