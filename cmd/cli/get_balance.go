package cli

import (
	"fmt"
	"log"

	"github.com/andskur/experemental/services/blockchain"
	"github.com/andskur/experemental/services/wallets"
	"github.com/andskur/experemental/utils"
)

// Get balance for specific address
func (cli *CLI) getBalance(address string) {
	if !wallets.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.NewBlockchain(address)

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := bc.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)

	err := bc.Db.Close()
	if err != nil {
		log.Panic(err)
	}
}
