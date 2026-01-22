package controller

import (
	"log"
	"net/http"
	"service/internal/configuration"
	"service/internal/server/httpServer"
)

type Controller struct {
	conf configuration.Configuration
}

func (c *Controller) StartWork() {
	c.conf.ReadConfigFile()
	handler := httpServer.NewHttpHandler()
	err := http.ListenAndServe(c.conf.Ip+":"+c.conf.Port, handler)
	if err != nil {
		log.Println(err)
	}
}
