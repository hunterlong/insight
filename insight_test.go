package insight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	tester *Insight
)

func TestNew(t *testing.T) {
	tester = New("https://btc.elementchains.com/api")
	assert.NotNil(t, tester)
}

func TestNewAddress(t *testing.T) {
	address := tester.NewAddress("1DCWizK9oeMCB2nEj1brZ4aGEvWjpS1KuU")
	assert.Equal(t, float64(6.408), address.Balance())
}

func TestLatestBlock(t *testing.T) {
	hash := tester.LatestBlock()
	assert.NotZero(t, hash)
}

func TestBTCBlockHash(t *testing.T) {
	hash, err := tester.BlockHash(532830)
	assert.Nil(t, err)
	assert.Equal(t, "00000000000000000009119366319b9ec9e3b1349d76b21bf73dbcb0f0528c91", hash)
}

func TestBTCBlockPages(t *testing.T) {
	pages, err := tester.blockPages("00000000000000000009119366319b9ec9e3b1349d76b21bf73dbcb0f0528c91")
	assert.Nil(t, err)
	assert.Equal(t, 283, pages)
}

func TestBTCBlockTxs(t *testing.T) {
	block, err := tester.BlockTransactions("00000000000000000009119366319b9ec9e3b1349d76b21bf73dbcb0f0528c91")
	assert.Nil(t, err)
	assert.Equal(t, 2826, len(block.Transactions))
	assert.Equal(t, 6331123, len(block.ToJSON()))
}

func TestBTCBlockTxs2(t *testing.T) {
	block, err := tester.BlockTransactions("00000000000000000024197edf087521335c9f66447580bc40e0a053d23341da")
	assert.Nil(t, err)
	assert.Equal(t, 2433, len(block.Transactions))
	assert.Equal(t, 6183899, len(block.ToJSON()))
}
