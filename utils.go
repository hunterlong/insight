package insight

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

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
