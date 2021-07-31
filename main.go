package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"etherscan_parse/scenarios"
)

func main() {
	var (
		apiKey      string
		apiRate     int
		blockAmount int
	)

	flag.StringVar(&apiKey, "k", "", "etherscan.io api key")
	flag.IntVar(&apiRate, "r", 5, "limits api rates")
	flag.IntVar(&blockAmount, "a", 100, "selects the amount of blocks that will be queried")
	flag.Parse()

	switch {
	case apiKey == "":
		log.Fatalln("API Key missing")
		return

	case apiRate < 1:
		log.Fatalln("API Rate cannot be lower than 1")
		return

	case blockAmount < 1:
		log.Fatalln("block amount cannot be lower than 1")
		return
	}

	addr, err := scenarios.GetMaxBalanceChangeAddr(context.Background(), apiKey, apiRate, blockAmount)
	if err != nil {
		log.Fatalf("%+v", err)
		return
	}

	fmt.Println(addr)
}
