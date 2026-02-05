package db

import (
	"log"
	"service/configs"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config interface {
	GetDSN() string // Data Source Name
	GetMaxOpenConns() int
	GetMaxIdleConns() int
	GetConnMaxIdleTime() time.Duration
	GetConnMaxLifetime() time.Duration
}

// принимает конфиг и возвращает бд
func ConnectDB(cfg Config) (*sqlx.DB, error) {
	if err := configs.ReadConfigYAML("config.yaml"); err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Open("pgx", cfg.GetDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.GetMaxOpenConns())
	db.SetMaxIdleConns(cfg.GetMaxIdleConns())
	db.SetConnMaxIdleTime(cfg.GetConnMaxIdleTime())
	db.SetConnMaxLifetime(cfg.GetConnMaxLifetime())

	return db, nil
}
