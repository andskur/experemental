package cli

import (
	"fmt"
	"log"

	"github.com/andskur/experemental/services/blockchain"
	"github.com/andskur/experemental/services/wallets"
)

// Create new blockchain Bd and mine genesis block
func (cli *CLI) createBlockchain(address string) {
	if !wallets.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.CreateBlockchain(address)

	UTXOSet := blockchain.UTXOSet{bc}
	UTXOSet.Reindex()

	err := bc.Db.Close()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Done!")
}