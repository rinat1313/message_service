package configuration

import (
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Port string
	Ip   string
}

func (conf *Configuration) ReadConfigFile() {
	godotenv.Load()
	ip := os.Getenv("SERVER_IP")
	port := os.Getenv("SERVER_PORT")
	conf.Port = port
	conf.Ip = ip
}
