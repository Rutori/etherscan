package scenarios

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"etherscan_parse/etherscanio"
	etherclient "etherscan_parse/etherscanio/client"
	"etherscan_parse/etherscanio/entities"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// GetMaxBalanceChangeAddr returns the address that had its balance changed more that every other one
// for the requested amount of last blocks
func GetMaxBalanceChangeAddr(ctx context.Context, apiKey string, rps, blockRange int) (string, error) {
	api := etherclient.NewAPI(apiKey, rps)
	lastBlockTag, err := etherscanio.GetLastBlock(ctx, api)
	if err != nil {
		return "", err
	}

	lastBlockNum, err := strconv.ParseInt(lastBlockTag[2:], 16, 64)
	if err != nil {
		return "", errors.WithStack(err)
	}

	blocksRecv := make(chan *entities.Block, blockRange)
	errG, errCTX := errgroup.WithContext(ctx)
	for i := 0; i < blockRange; i++ {
		lastBlockNum--
		prevBlock := lastBlockNum
		errG.Go(func() error {
			return queryBlock(errCTX, api, blocksRecv, fmt.Sprintf("0x%s", strconv.FormatInt(prevBlock, 16)))
		})
	}

	amounts := make(chan map[string]*big.Int, 1)
	errc := make(chan error)
	go func() {
		errc <- calculateAmounts(blocksRecv, amounts)
	}()

	err = errG.Wait()
	if err != nil {
		return "", err
	}

	//closing the receiving channel so calculation could end
	close(blocksRecv)

	// checking for errors in calculateAmounts
	err = <-errc
	if err != nil {
		return "", err
	}

	var topAddr string

	topAmount := new(big.Int).SetInt64(0)

	for addr, amount := range <-amounts {
		if amount.CmpAbs(topAmount) < 1 {
			continue
		}

		topAddr = addr
		topAmount = amount
	}

	return topAddr, nil
}

// queryBlock fetches a block and puts it into the outgoing channel
func queryBlock(ctx context.Context, api *etherclient.API, out chan<- *entities.Block, blockTag string) error {
	block, err := etherscanio.GetBlock(ctx, api, blockTag)
	if err != nil {
		return err
	}

	out <- block
	return nil
}

// calculateAmounts constantly recieves blocks, calculates the change to the overall transactions data
// and after the inbound channel is closed writes the result into the output channel
func calculateAmounts(in <-chan *entities.Block, out chan<- map[string]*big.Int) error {
	amounts := make(map[string]*big.Int)
	for block := range in {
		blockChanges, err := block.GetSpending()
		if err != nil {
			return err
		}

		addAmounts(amounts, blockChanges)
	}

	out <- amounts

	return nil
}

// addAmounts adds the new transaction data into the main map
func addAmounts(everyone map[string]*big.Int, delta map[string]*big.Int) {
	for addr, change := range delta {
		_, exists := everyone[addr]
		if !exists {
			everyone[addr] = change
			continue
		}

		everyone[addr].Add(everyone[addr], delta[addr])
	}
}
