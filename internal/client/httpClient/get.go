package httpClient

import (
	"io"
	"log"
	"net/http"
)

func (httpClient *ClientConnector) Get() string {
	req, err := http.NewRequest(http.MethodGet, httpClient.urlServer, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Accepts", "application/json")
	res, err := httpClient.client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return string(body)
}
