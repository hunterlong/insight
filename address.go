package insight

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	insight *Insight
	Address string
	*balancesQuery
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
