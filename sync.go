package insight

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

func (i *insight) Sync() *sync {
	url := fmt.Sprintf("%v/sync", i.Endpoint)
	body, err := httpMethod(url, nil)
	if err != nil {
		return nil
	}
	var sync *sync
	err = json.Unmarshal(body, &sync)
	if err != nil {
		panic(err)
	}
	return sync
}

func (i *insight) Peers() *peers {
	url := fmt.Sprintf("%v/peer", i.Endpoint)
	body, err := httpMethod(url, nil)
	if err != nil {
		return nil
	}
	var peers *peers
	err = json.Unmarshal(body, &peers)
	if err != nil {
		panic(err)
	}
	return peers
}

func (i *insight) Ping() error {
	sync := i.Sync()
	if sync == nil {
		return errors.New("could not fetch latest sync url")
	}
	return nil
}

func (i *insight) LatestBlock() int {
	sync := i.Sync()
	if sync == nil {
		return 0
	}
	return sync.Height
}

func (i *insight) Synced() bool {
	sync := i.Sync()
	if sync == nil {
		return false
	}
	if sync.Status == "finished" {
		return true
	}
	return false
}

func (i *insight) SyncPercent() float64 {
	sync := i.Sync()
	if sync == nil {
		return 0
	}
	return float64(sync.SyncPercentage)
}

type peers struct {
	Connected bool   `json:"connected"`
	Host      string `json:"host"`
	Port      *int   `json:"port"`
}
