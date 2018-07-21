package insight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	INSIGHT_API = "https://btc.coinapp.io/api"
)

var (
	tester      *Insight
	address     *Address
	block       *Block
	block2      *Block
	latestBlock *Block
)

func TestNew(t *testing.T) {
	tester = New(INSIGHT_API)
	assert.NotNil(t, tester)
}

func TestNewAddress(t *testing.T) {
	address = tester.NewAddress("1Hz96kJKF2HLPGY15JWLB5m9qGNxvt8tHJ")
	assert.Equal(t, float64(111.068415688), address.Balance())
}

func TestNewBlockID(t *testing.T) {
	block = tester.NewBlock(532895)
	assert.Equal(t, 532895, block.Height)
	assert.Equal(t, "000000000000000000093a6e4e1c993d5cb57ff6b2d6dfccdfe77de48de89cd8", block.Hash)
	assert.Equal(t, 3, block.Pages)
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
	utxo := address.UTXO()
	assert.Equal(t, 11, len(utxo))
	assert.Equal(t, "0ba99fcc7b9e7b5991394320d59293d66c8d787175c1a0166de90737706a51f1", utxo[0].Txid)
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
	addressTrx := tester.NewAddress("1KWaj9LHXyLBzGU1Q5rK5CoZwDBANFksgf")
	transactions := addressTrx.Transactions()
	assert.Equal(t, 2, len(transactions))
}

func TestBTCBlockTxs(t *testing.T) {
	blockTransactions, err := block.Transactions()
	assert.Nil(t, err)
	assert.Equal(t, 27, len(blockTransactions))
	assert.NotZero(t, len(block.ToJSON()))
}
