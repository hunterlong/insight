package insight

import (
	"encoding/json"
	"fmt"
)

type block struct {
	Hash    string
	Height  int
	Pages   int
	insight *insight
}

func (a *insight) NewBlock(id interface{}) *block {
	b := new(block)
	b.insight = a
	switch v := id.(type) {
	case int:
		b.Height = int(v)
		b.hash()
		b.pages()
		b.info()
		return b
	case string:
		b.Hash = string(v)
		b.pages()
		b.info()
		return b
	case nil:
		return b.latestBlock()
	}
	return nil
}

func (b *block) IsLatest() bool {
	sync := b.insight.Sync()
	if sync.Height == b.Height {
		return true
	}
	return false
}

func (b *block) latestBlock() *block {
	sync := b.insight.Sync()
	b.Height = sync.Height
	b.hash()
	b.info()

	return b
}

func (b *block) Transactions() ([]*blockTx, error) {
	var trxs []*blockTx
	for p := 0; p <= b.Pages; p++ {
		bTrax, _ := b.blockTransactions(b.Hash, p)
		for _, tx := range bTrax.BlockTxs {
			trxs = append(trxs, tx)
		}
	}
	return trxs, nil
}

func (b *block) blockTransactions(hash string, page int) (*blockTransactions, error) {
	url := fmt.Sprintf("%v/txs/?block=%v&pageNum=%v", b.insight.Endpoint, hash, page)
	body, err := httpMethod(url, nil)
	if err != nil {
		return nil, err
	}
	var block *blockTransactions
	err = json.Unmarshal(body, &block)
	if err != nil {
		panic(err)
	}
	return block, err
}

func (b *block) pages() int {
	url := fmt.Sprintf("%v/txs/?block=%v", b.insight.Endpoint, b.Hash)
	body, err := httpMethod(url, nil)
	if err != nil {
		return 0
	}
	var blkjson *blockJson
	json.Unmarshal(body, &blkjson)
	b.Pages = blkjson.Pages
	return blkjson.Pages
}

func (b *block) hash() (string, error) {
	url := fmt.Sprintf("%v/block-index/%v", b.insight.Endpoint, b.Height)
	body, err := httpMethod(url, nil)
	var hash *blockHash
	err = json.Unmarshal(body, &hash)
	b.Hash = hash.BlockHash
	return hash.BlockHash, err
}

func (b *block) ToJSON() string {
	data, _ := json.Marshal(b)
	return string(data)
}

func (b *block) info() (*blockInformation, error) {
	url := fmt.Sprintf("%v/block/%v", b.insight.Endpoint, b.Hash)
	body, err := httpMethod(url, nil)
	var info *blockInformation
	json.Unmarshal(body, &info)

	b.Hash = info.Hash
	b.Height = info.Height

	return info, err
}

type blockInformation struct {
	Hash              string   `json:"hash"`
	Size              int      `json:"size"`
	Height            int      `json:"height"`
	Version           int      `json:"version"`
	Merkleroot        string   `json:"merkleroot"`
	Tx                []string `json:"tx"`
	Time              int      `json:"time"`
	Nonce             int64    `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	Chainwork         string   `json:"chainwork"`
	Confirmations     int      `json:"confirmations"`
	Previousblockhash string   `json:"previousblockhash"`
	Nextblockhash     string   `json:"nextblockhash"`
	Reward            float64  `json:"reward"`
	IsMainChain       bool     `json:"isMainChain"`
	PoolInfo          struct {
	} `json:"poolInfo"`
}
