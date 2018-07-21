# Bitcore Insight API - Golang
>A small golang package to help you use your insight API.

[![Go Report Card](https://goreportcard.com/badge/github.com/hunterlong/insight)](https://goreportcard.com/report/github.com/hunterlong/insight)
[![Build Status](https://travis-ci.com/hunterlong/insight.svg?branch=master)](https://travis-ci.com/hunterlong/insight)
[![Godoc](https://godoc.org/github.com/hunterlong/insight?status.svg)](https://godoc.org/github.com/hunterlong/insight)

### Connect to Insight API
```go
// import "github.com/hunterlong/insight"

bitcoin := insight.New("https://insight.bitpay.com/api")

// err := bitcoin.Ping()
```

### Get Block Hash
```go
block := bitcoin.NewBlock(532833) // you can use Block Height
block := bitcoin.NewBlock("0000000000000000003221e76df18226231f8e33694f40318051ec688decd6b0") // or hash
block := bitcoin.NewBlock(nil) // or get the latest block
```

### Fetch Bitcoin Block
```go
transactions := block.Transactions()
// transactions.ToJSON()
```

### Get Balance for Address
```go
address := bitcoin.NewAddress("3KJrsjfg1dD6CrsTeHdHVH3KqMpvL2XWQn")
balance := address.Balance()
```

### Get Address UTXO's
```go
utxos := address.UTXO()
```

### Get Server Sync Status
```go
sync := bitcoin.Sync()
// sync.Status
// sync.BlockChainHeight
```
