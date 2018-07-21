package insight

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	insight *insight
	Address string
	*balancesQuery
}

func (a *insight) NewAddress(address string) *Address {
	return &Address{insight: a, Address: address}
}

func (a *Address) Balance() float64 {
	if a.balancesQuery == nil {
		a.balancesQuery = a.balanceQuery()
	}
	btc := float64(a.BalanceSat) * 0.000000008
	return btc
}

func (a *Address) PendingBalance() float64 {
	if a.balancesQuery == nil {
		a.balancesQuery = a.balanceQuery()
	}
	btc := float64(a.UnconfirmedBalanceSat) * 0.000000008
	return btc
}

func (a *Address) balanceQuery() *balancesQuery {
	url := fmt.Sprintf("%v/addr/%v", a.insight.Endpoint, a.Address)
	body, err := httpMethod(url, nil)
	var balance *balancesQuery
	err = json.Unmarshal(body, &balance)
	if err != nil {
		panic(err)
	}
	a.balancesQuery = balance
	return balance
}

func (a *Address) UTXO() []*UTXO {
	url := fmt.Sprintf("%v/addr/%v/utxo", a.insight.Endpoint, a.Address)
	body, err := httpMethod(url, nil)
	var utxos []*UTXO
	err = json.Unmarshal(body, &utxos)
	if err != nil {
		panic(err)
	}
	return utxos
}

func (a *Address) Transactions() []*AddressTxs {
	pages := a.txPages()
	var transactions []*AddressTxs
	for p := 0; p <= pages; p++ {
		txs := a.transactions(p)
		for _, tx := range txs {
			transactions = append(transactions, tx)
		}
	}
	return transactions
}

func (a *Address) txPages() int {
	url := fmt.Sprintf("%v/txs?address=%v", a.insight.Endpoint, a.Address)
	body, err := httpMethod(url, nil)
	var txs *AddressTransactions
	err = json.Unmarshal(body, &txs)
	if err != nil {
		panic(err)
	}
	return txs.PagesTotal
}

func (a *Address) transactions(page int) []*AddressTxs {
	url := fmt.Sprintf("%v/txs?address=%v&pageNum=%v", a.insight.Endpoint, a.Address, page)
	body, err := httpMethod(url, nil)
	var txs *AddressTransactions
	err = json.Unmarshal(body, &txs)
	if err != nil {
		panic(err)
	}
	return txs.Txs
}

type balancesQuery struct {
	AddrStr                 string   `json:"addrStr"`
	Balance                 float64  `json:"balance"`
	BalanceSat              int      `json:"balanceSat"`
	TotalReceived           float64  `json:"totalReceived"`
	TotalReceivedSat        int64    `json:"totalReceivedSat"`
	TotalSent               float64  `json:"totalSent"`
	TotalSentSat            int64    `json:"totalSentSat"`
	UnconfirmedBalance      int      `json:"unconfirmedBalance"`
	UnconfirmedBalanceSat   int      `json:"unconfirmedBalanceSat"`
	UnconfirmedTxApperances int      `json:"unconfirmedTxApperances"`
	TxApperances            int      `json:"txApperances"`
	Transactions            []string `json:"transactions"`
}

type UTXOS struct {
	UTXOs []UTXO
}

type UTXO struct {
	Address       string  `json:"address"`
	Txid          string  `json:"txid"`
	Vout          int     `json:"vout"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	Amount        float64 `json:"amount"`
	Satoshis      int     `json:"satoshis"`
	Height        int     `json:"height"`
	Confirmations int     `json:"confirmations"`
}

type AddressTransactions struct {
	PagesTotal int           `json:"pagesTotal"`
	Txs        []*AddressTxs `json:"txs"`
}

type AddressTxs struct {
	Txid     string `json:"txid"`
	Version  int    `json:"version"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Coinbase string `json:"coinbase"`
		Sequence int64  `json:"sequence"`
		N        int    `json:"n"`
	} `json:"vin"`
	Vout []struct {
		Value        string `json:"value"`
		N            int    `json:"n"`
		ScriptPubKey struct {
			Hex       string   `json:"hex"`
			Asm       string   `json:"asm"`
			Addresses []string `json:"addresses"`
			Type      string   `json:"type"`
		} `json:"scriptPubKey"`
		SpentTxID   interface{} `json:"spentTxId"`
		SpentIndex  interface{} `json:"spentIndex"`
		SpentHeight interface{} `json:"spentHeight"`
	} `json:"vout"`
	Blockhash     string  `json:"blockhash"`
	Blockheight   int     `json:"blockheight"`
	Confirmations int     `json:"confirmations"`
	Time          int     `json:"time"`
	Blocktime     int     `json:"blocktime"`
	IsCoinBase    bool    `json:"isCoinBase,omitempty"`
	ValueOut      float64 `json:"valueOut"`
	Size          int     `json:"size"`
	ValueIn       float64 `json:"valueIn,omitempty"`
	Fees          float64 `json:"fees,omitempty"`
}
