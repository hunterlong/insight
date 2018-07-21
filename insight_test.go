package insight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	INSIGHT_API = "https://btc.coinapp.io/api"
)

var (
	tester      *insight
	addr        *address
	blk         *block
	block2      *block
	latestBlock *block
)

func TestNew(t *testing.T) {
	tester = New(INSIGHT_API)
	assert.NotNil(t, tester)
}

func TestNewAddress(t *testing.T) {
	addr = tester.NewAddress("37TmaiDdMwdNM9J5iKDcWdPVT1Zi8wNYJH")
	assert.Equal(t, float64(0.10384547), addr.Balance())
}

func TestNewBlockID(t *testing.T) {
	blk = tester.NewBlock(532895)
	assert.Equal(t, 532895, blk.Height)
	assert.Equal(t, "000000000000000000093a6e4e1c993d5cb57ff6b2d6dfccdfe77de48de89cd8", blk.Hash)
	assert.Equal(t, 3, blk.Pages)
}

func TestNewLatestBlock(t *testing.T) {
	current := tester.Sync()
	latestBlock = tester.NewBlock(nil)
	assert.Equal(t, current.Height, latestBlock.Height)
	assert.NotEmpty(t, latestBlock.Hash)
}

func TestNewBlockHash(t *testing.T) {
	block2 = tester.NewBlock("000000000000000000093a6e4e1c993d5cb57ff6b2d6dfccdfe77de48de89cd8")
	assert.Equal(t, 532895, block2.Height)
	assert.Equal(t, "000000000000000000093a6e4e1c993d5cb57ff6b2d6dfccdfe77de48de89cd8", block2.Hash)
	assert.Equal(t, 3, block2.Pages)
}

func TestAddressUTXO(t *testing.T) {
	utxo := addr.UTXO()
	assert.Equal(t, 1, len(utxo))
	assert.Equal(t, "66e541fd4eebc74c48a5bf0b8ee036e5dd909c7f4ffbb2cefe4cb66498bdbd43", utxo[0].Txid)
}

func TestSync(t *testing.T) {
	sync := tester.Sync()
	assert.Equal(t, "finished", sync.Status)
	assert.NotZero(t, sync.Height)
}

func TestPeer(t *testing.T) {
	peers := tester.Peers()
	assert.Equal(t, "127.0.0.1", peers.Host)
}

func TestLatestBlock(t *testing.T) {
	hash := tester.LatestBlock()
	assert.NotZero(t, hash)
}

func TestAddressTransactions(t *testing.T) {
	addressTrx := tester.NewAddress("37TmaiDdMwdNM9J5iKDcWdPVT1Zi8wNYJH")
	transactions := addressTrx.Transactions()
	assert.Equal(t, 3, len(transactions))
}

func TestBTCBlockTxs(t *testing.T) {
	blockTransactions, err := blk.Transactions()
	assert.Nil(t, err)
	assert.Equal(t, 27, len(blockTransactions))
	assert.NotZero(t, len(blk.ToJSON()))
}

func TestBTCBlockTxs2(t *testing.T) {
	blk = tester.NewBlock(532923)
	blk.insight.SetThreads(8)
	blockTransactions, err := blk.Transactions()
	assert.Nil(t, err)
	assert.Equal(t, 985, len(blockTransactions))
	assert.NotZero(t, len(blk.ToJSON()))
}

func TestBTCBlockTxs3(t *testing.T) {
	blk = tester.NewBlock(532921)
	blk.insight.SetThreads(16)
	blockTransactions, err := blk.Transactions()
	assert.Nil(t, err)
	assert.Equal(t, 2457, len(blockTransactions))
	assert.NotZero(t, len(blk.ToJSON()))
}
