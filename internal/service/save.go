package service

import (
	"database/sql"
	"errors"
	"fmt"
)

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
	connStr := "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return 0, err
	}

	fmt.Println("Successfully connected to postgres")
	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return 0, err
	}
	fmt.Printf("Версия СУБД: %s", version)

	id := 0

	return id, nil
}
