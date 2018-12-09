package main

import "fmt"

func main() {
	bc := NewBlockchain()


	bc.AddBlock("Remittance tx from Moscow to Dushanbe")
	bc.AddBlock("Remittance tx from Dushanbe to Tashkent")

	//fmt.Println(bc.blocks[0].Data)

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Timestamp: %x\n", block.Timestamp)
		fmt.Println()
	}
}
