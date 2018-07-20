package insight

import (
	"encoding/json"
	"fmt"
)

func (i *Insight) BlockTransactions(hash string) (*Block, error) {
	pages, _ := i.blockPages(hash)
	txs := new(Block)
	txs.PagesTotal = pages
	txs.Transactions = nil
	for p := 0; p <= pages; p++ {
		block, _ := i.blockTransactions(hash, p)
		for _, tx := range block.Transactions {
			txs.Transactions = append(txs.Transactions, tx)
		}
	}
	return txs, nil
}

func (i *Insight) blockTransactions(hash string, page int) (*Block, error) {
	url := fmt.Sprintf("%v/txs/?block=%v&pageNum=%v", i.Endpoint, hash, page)
	body, err := httpMethod(url, nil)
	var block *Block
	err = json.Unmarshal(body, &block)
	if err != nil {
		panic(err)
	}
	return block, err
}

func (i *Insight) blockPages(hash string) (int, error) {
	url := fmt.Sprintf("%v/txs/?block=%v", i.Endpoint, hash)
	body, err := httpMethod(url, nil)
	var block *Block
	json.Unmarshal(body, &block)
	return block.PagesTotal, err
}

func (i *Insight) BlockHash(id int64) (string, error) {
	url := fmt.Sprintf("%v/block-index/%v", i.Endpoint, id)
	body, err := httpMethod(url, nil)
	var hash *BlockHash
	err = json.Unmarshal(body, &hash)
	return hash.BlockHash, err
}

func (b *Block) ToJSON() string {
	data, _ := json.Marshal(b)
	return string(data)
}
