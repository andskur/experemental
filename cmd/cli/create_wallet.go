package cli

import (
	"fmt"
	"github.com/andskur/experemental/services/wallets"
)

// Create new wallet
func (cli *CLI) createWallet() {
	keystore, _ := wallets.NewWallets()
	address := keystore.CreateWallet()
	keystore.SaveToFile()

	fmt.Printf("Your new address: %s\n", address)
}
