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
}

func New(endpoint string) *insight {
	return &insight{Endpoint: endpoint}
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
