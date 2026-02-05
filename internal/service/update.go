package service

import "errors"

func UpdateMessage(body string) error {
	message, err := ParseJson(body)
	if err != nil {
		return err
	}
	if len(message.Body) == 0 {
		return errors.New("Message Body is empty")
	}

	// тут мы отправляем запрос в бд на изменение сообщения
	// метод возвращает ошибку.

	return nil
}
