package insight

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type insight struct {
	Endpoint string
	Timeout  int
	ch       chan struct{}
}

// New insight pointer will query your insight API, requires insight /api endpoint
func New(endpoint string) *insight {
	i := new(insight)
	i.ch = make(chan struct{}, 1)
	i.Endpoint = endpoint
	return i
}

func (i *insight) SetThreads(amount int) {
	i.ch = make(chan struct{}, amount)
}

func httpMethod(url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{Timeout: 120 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
