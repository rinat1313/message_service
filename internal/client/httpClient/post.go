package httpClient

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func (httpClient *ClientConnector) Post(body string) string {
	req, err := http.NewRequest(http.MethodPost, httpClient.urlServer, strings.NewReader(body))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return string(b)
}
