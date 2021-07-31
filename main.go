package main

import "flag"

func main() {
	var (
		apiKey  string
		apiRate int
	)

	flag.StringVar(&apiKey, "key", "", "etherscan.io api key")
	flag.IntVar(&apiRate, "max api rate", 5, "limits api rates")
}
