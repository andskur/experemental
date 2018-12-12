package cli

import (
	"fmt"
	"log"

	"github.com/andskur/experemental/services/wallets"
)

// List all addresses from .dat file
func (cli *CLI) listAddresses() {
	keystore, err := wallets.NewWallets()
	if err != nil {
		log.Panic(err)
	}
	addresses := keystore.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
