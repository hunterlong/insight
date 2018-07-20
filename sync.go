package insight

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

func (i *Insight) Sync() *Sync {
	url := fmt.Sprintf("%v/sync", i.Endpoint)
	body, err := httpMethod(url, nil)
	var sync *Sync
	err = json.Unmarshal(body, &sync)
	if err != nil {
		panic(err)
	}
	return sync
}

func (i *Insight) Ping() error {
	sync := i.Sync()
	if sync == nil {
		return errors.New("could not fetch latest sync url")
	}
	return nil
}

func (i *Insight) LatestBlock() int {
	sync := i.Sync()
	if sync == nil {
		return 0
	}
	return sync.Height
}

func (i *Insight) Synced() bool {
	sync := i.Sync()
	if sync == nil {
		return false
	}
	if sync.Status == "finished" {
		return true
	}
	return false
}

func (i *Insight) SyncPercent() float64 {
	sync := i.Sync()
	if sync == nil {
		return 0
	}
	return float64(sync.SyncPercentage)
}
