package cli

import (
	"fmt"
	"github.com/andskur/experemental/services/blockchain"
)

func (cli *CLI) reindexUTXO() {
	bc := blockchain.NewBlockchain()
	UTXOSet := blockchain.UTXOSet{bc}
	UTXOSet.Reindex()
	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}