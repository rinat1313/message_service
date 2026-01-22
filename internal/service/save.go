package service

import "errors"

func SavingMessage(text string) (int, error) {
	message, err := ParseJson(text)
	if err != nil {
		return 0, err
	}
	if len(message.Body) == 0 {
		return 0, errors.New("Message Body is empty")
	}

	// тут мы отправляем запрос в бд на сохранение сообщения
	// метод должен возвращать id сохранённого сообщения в БД
	id := 0

	return id, nil
}
