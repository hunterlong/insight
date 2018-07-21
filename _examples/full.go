package main

import (
	"fmt"
	"github.com/hunterlong/insight"
)

func main() {

	bitcoin := insight.New("https://btc.coinapp.io/api")

	err := bitcoin.Ping()
	if err != nil {
		panic(err)
	}

	bitcoin.SetThreads(4)

	block := bitcoin.NewBlock(nil)

	address := bitcoin.NewAddress("162yoHmpvqCq5SztkKaiEd5LiPyxcPLYzu")

	fmt.Println("Current block: ", block.Height)
	fmt.Println("Current hash: ", block.Hash)

	transactions, err := block.Transactions()
	if err != nil {
		panic(err)
	}

	fmt.Println("Current block Transactions: ", len(transactions))

	peers := bitcoin.Peers()
	sync := bitcoin.Sync()

	fmt.Println("Fully Syned: ", sync.Status)
	fmt.Println("Sync Height: ", sync.Height)
	fmt.Println("Host: ", peers.Host)

	fmt.Println("Server up-to-date: ", block.IsLatest())

	fmt.Println("Block JSON output: ", block.ToJSON())

	fmt.Println("Address: ", address.String())
	fmt.Println("Address Balance: ", address.Balance())
	fmt.Println("Address UTXOs: ", len(address.UTXO()))
	fmt.Println("Address Total Sent: ", address.TotalSent)

	bitcoin.SetThreads(16)

	block2 := bitcoin.NewBlock(532924)

	transactions, _ = block2.Transactions()

	fmt.Println("Block: ", block2.Height)
	fmt.Println("Block Transactions: ", len(transactions))

}
