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
hash := bitcoin.BlockHash(532833)
// 0000000000000000003221e76df18226231f8e33694f40318051ec688decd6b0
```

### Fetch Bitcoin Block
```go
hash := "0000000000000000003221e76df18226231f8e33694f40318051ec688decd6b0"
block := bitcoin.BlockTransactions(hash)

// block.ToJSON()
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

