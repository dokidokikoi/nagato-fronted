package tools

import (
	"fmt"
	"io"
	"net/http"
)

func SendReq(method, url string, body io.Reader, header map[string]string) ([]byte, int, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("http://localhost:10000/api/%s", url), body)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	fmt.Printf("-----------------%s-----------------\n", fmt.Sprintf("http://localhost:10000/api/%s", url))
	if err != nil {
		return nil, -1, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	return data, resp.StatusCode, err
}
