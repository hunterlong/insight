package insight

type BlockHash struct {
	BlockHash string `json:"blockHash"`
}

type Block struct {
	PagesTotal   int            `json:"pagesTotal"`
	Transactions []*Transaction `json:"txs"`
}

type Transaction struct {
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

type Sync struct {
	Status           string      `json:"status"`
	BlockChainHeight int         `json:"blockChainHeight"`
	SyncPercentage   int         `json:"syncPercentage"`
	Height           int         `json:"height"`
	Error            interface{} `json:"error"`
	Type             string      `json:"type"`
}
