package service

import (
	"golang-api/config"
	"log"
)

type Service struct {
	Logger *log.Logger
	Config *config.Config
}

func New() Service {
	return Service{}
}
