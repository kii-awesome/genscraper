package main

import (
	"os"
	"fmt"
	"encoding/json"
	"flag"
	"github.com/kii-awesome/genscraper/models"
)

func main()  {
	fileName := flag.String("file", "", "path to genesis.json")
	// filter
	address := flag.String("address", "", "your address for filtering")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("usage --file path-to-genesis")
		os.Exit(1)
	}

	// read
	file, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("failed to read file")
		os.Exit(1)
	}

	var genesis models.Genesis
	if err := json.Unmarshal(file, &genesis); err != nil {
		fmt.Println("failed to parsing")
		os.Exit(1)
	}

	if *address == "" {
		for _, balance := range genesis.AppState.Bank.Balances {
			for _, coin := range balance.Coins {
				fmt.Printf("Address: %s\nBalances: %s %s\n" , balance.Address , coin.Amount, coin.Denom)
			}
		}
	} else {
		for _, balance := range genesis.AppState.Bank.Balances {
			if balance.Address == *address {
				for _, coin := range balance.Coins {
					fmt.Printf("Address: %s\nBalances: %s %s\n" , balance.Address , coin.Amount, coin.Denom)
				}
				return
			}
		}
		fmt.Println("Address not found")
	}
	
}