package main

import (
	"andskur/blockchain/experemental/blockchain"
	"andskur/blockchain/experemental/cmd/cli"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Db.Close()

	run := cli.CLI{Bc: bc}
	run.Run()
}
