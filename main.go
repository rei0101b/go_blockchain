package main

import (
	"fmt"
	"log"
	"main/block"
	"main/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()
	// wallet
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)
	// blockchain
	blockchain := block.NewBlockChain(walletM.BlockchainAddress())
	isAdd := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0, walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added? ", isAdd)

	blockchain.Mining()
	blockchain.Print()
	fmt.Printf("A %.1f\n", blockchain.CaluculateTolalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CaluculateTolalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CaluculateTolalAmount(walletM.BlockchainAddress()))
}
