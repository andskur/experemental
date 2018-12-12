package cli

import (
	"fmt"
	"log"

	"github.com/andskur/experemental/services/blockchain"
	"github.com/andskur/experemental/services/blockchain/txs"
	"github.com/andskur/experemental/services/wallets"
)

// Make transaction from one address to another
func (cli *CLI) send(from, to string, amount int) {
	if !wallets.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !wallets.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := blockchain.NewBlockchain(from)
	defer bc.Db.Close()

	tx := blockchain.NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*txs.Transaction{tx})
	fmt.Println("Success!")
}
