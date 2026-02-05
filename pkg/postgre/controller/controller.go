package controller

import (
	"database/sql"
	"fmt"
	"service/internal/domen"
	"time"

	_ "github.com/lib/pq"
)

func connectDB() {
	connStr := "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully connected to postgres")
	//var version string
	//err = db.QueryRow("SELECT version()").Scan(&version)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("Версия СУБД: %s", version)
	insertWithStruct(db)
}

func insertWithStruct(db *sql.DB) {
	msg := domen.Message{
		Timestamp: time.Now(),
		Body:      "Тестовое сообщение",
		IdChat:    int64(1),
	}
	query := `INSERT INTO messages (time, body, id_chat) VALUES ($1, $2, $3) returning id`
	var id int64
	err := db.QueryRow(query, msg.Timestamp, msg.Body, msg.IdChat).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Сохранено сообщение с id = %d\n", id)
}
