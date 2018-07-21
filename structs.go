package insight

type blockHash struct {
	BlockHash string `json:"blockHash"`
}

type blockJson struct {
	Pages        int            `json:"pagesTotal"`
	Transactions []*transaction `json:"txs"`
}

type transaction struct {
	Txid     string `json:"txid"`
	Version  int    `json:"version"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		Sequence  int64  `json:"sequence"`
		N         int    `json:"n"`
		ScriptSig struct {
			Hex string `json:"hex"`
			Asm string `json:"asm"`
		} `json:"scriptSig"`
		Addr            string      `json:"addr"`
		ValueSat        int         `json:"valueSat"`
		Value           float64     `json:"value"`
		DoubleSpentTxID interface{} `json:"doubleSpentTxID"`
	} `json:"vin"`
	Vout          []*InsightVout `json:"vout"`
	Blockhash     string         `json:"blockhash"`
	Blockheight   int            `json:"blockheight"`
	Confirmations int            `json:"confirmations"`
	Time          int            `json:"time"`
	Blocktime     int            `json:"blocktime"`
	ValueOut      float64        `json:"valueOut"`
	Size          int            `json:"size"`
	ValueIn       float64        `json:"valueIn"`
	Fees          float64        `json:"fees"`
}

type InsightVout struct {
	Value        string               `json:"value"`
	N            int                  `json:"n"`
	ScriptPubKey *InsightScriptPubKey `json:"scriptPubKey"`
	SpentTxID    interface{}          `json:"spentTxId"`
	SpentIndex   interface{}          `json:"spentIndex"`
	SpentHeight  interface{}          `json:"spentHeight"`
}

type InsightScriptPubKey struct {
	Hex       string   `json:"hex"`
	Asm       string   `json:"asm"`
	Addresses []string `json:"addresses"`
	Type      string   `json:"type"`
}

type sync struct {
	Status           string      `json:"status"`
	BlockChainHeight int         `json:"blockChainHeight"`
	SyncPercentage   int         `json:"syncPercentage"`
	Height           int         `json:"height"`
	Error            interface{} `json:"error"`
	Type             string      `json:"type"`
}

type blockTransactions struct {
	PagesTotal int        `json:"pagesTotal"`
	BlockTxs   []*blockTx `json:"txs"`
}

type blockTx struct {
	Txid     string `json:"txid"`
	Version  int    `json:"version"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Coinbase string `json:"coinbase"`
		Sequence int    `json:"sequence"`
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
	Blockhash     string      `json:"blockhash"`
	Blockheight   int         `json:"blockheight"`
	Confirmations int         `json:"confirmations"`
	Time          int         `json:"time"`
	Blocktime     int         `json:"blocktime"`
	IsCoinBase    bool        `json:"isCoinBase,omitempty"`
	ValueOut      float64     `json:"valueOut"`
	Size          int         `json:"size"`
	ValueIn       float64     `json:"valueIn,omitempty"`
	Fees          interface{} `json:"fees,omitempty"`
}
