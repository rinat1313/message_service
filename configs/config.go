package configs

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

var (
	cfg *Config
)

type Config struct {
	DB DB "db"
}

type DB struct {
	DSN             string        "DSN"
	MaxOpenConns    int           "maxOpenConns"
	MaxIdleConns    int           "maxIdleConns"
	ConnMaxIdleTime time.Duration "connMaxIdleTime"
	ConnMaxLifetime time.Duration "connMaxLifetime"
}

func (cfg *Config) GetDSN() string {
	return cfg.DB.DSN
}

func (cfg *Config) GetMaxOpenConns() int {
	return cfg.DB.MaxOpenConns
}

func (cfg *Config) GetMaxIdleConns() int {
	return cfg.DB.MaxIdleConns
}

func (cfg *Config) GetConnMaxIdleTime() time.Duration {
	return cfg.DB.ConnMaxIdleTime
}

func ReadConfigYAML(configYML string) error {
	if cfg != nil {
		return nil
	}
	file, err := os.Open(configYML)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}
	fmt.Println("DSN:", cfg.DB.DSN)
	fmt.Println("MaxOpenConns:", cfg.DB.MaxOpenConns)
	fmt.Println("MaxIdleConns:", cfg.DB.MaxIdleConns)
	fmt.Println("ConnMaxIdleTime:", cfg.DB.ConnMaxIdleTime)
	fmt.Println("ConnMaxLifetime:", cfg.DB.ConnMaxLifetime)
	return nil
}
