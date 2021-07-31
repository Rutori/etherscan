package etherscanio

import (
	"context"
	"encoding/json"
	"fmt"

	"etherscan_parse/etherscanio/client"
	"etherscan_parse/etherscanio/entities"
	"github.com/pkg/errors"
)

const blockInfoEndpoint = "https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=%s&boolean=true"

// GetBlock fetches block info
func GetBlock(ctx context.Context, api *client.API, blockTag string) (*entities.Block, error) {
	data, err := api.Query(ctx, fmt.Sprintf(blockInfoEndpoint, blockTag))
	if err != nil {
		return nil, err
	}

	block := entities.AnswerBlockInfo{}
	err = json.Unmarshal(data, &block)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return block.Result, nil
}
