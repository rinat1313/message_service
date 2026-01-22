package httpClient

import (
	"net/http"
	"time"
)

type ClientConnector struct {
	client    *http.Client
	urlServer string
}

func (httpClient *ClientConnector) CreateHttpClient() {
	httpClient.client = &http.Client{
		Transport:     nil,             // Механика выполнения запросов, RoundTripper интерфейс
		CheckRedirect: nil,             // Функция, определяющая политику работы редиректа
		Jar:           nil,             // Общее хранилище выполнения запросов
		Timeout:       2 * time.Second, // Лимит времени выполнения запроса
	}
}

// надо добавить валидацию сайтов
func (httpClient *ClientConnector) SetUrl(url string) error {
	httpClient.urlServer = url
	return nil
}
