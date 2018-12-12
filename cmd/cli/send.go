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

	bc := blockchain.NewBlockchain()
	UTXOSet := blockchain.UTXOSet{bc}

	tx := blockchain.NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbTx := txs.NewCoinbaseTX(from, "")
	txs := []*txs.Transaction{cbTx, tx}

	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)
	fmt.Println("Success!")

	err := bc.Db.Close()
	if err != nil {
		log.Panic(err)
	}
}
