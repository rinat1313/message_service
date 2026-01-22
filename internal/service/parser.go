package service

import (
	"encoding/json"
	"service/internal/domen"
)

func ParseJson(text string) (domen.Message, error) {
	result := domen.Message{}
	err := json.Unmarshal([]byte(text), &result)
	if err != nil {
		return domen.Message{}, err
	}
	return result, nil
}
