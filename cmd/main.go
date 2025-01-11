package main

import (
	"github.com/ub1vashka/golangfirstproject/internal/config"
	"github.com/ub1vashka/golangfirstproject/internal/logger"
	"github.com/ub1vashka/golangfirstproject/internal/server"
	"github.com/ub1vashka/golangfirstproject/internal/service"
	"github.com/ub1vashka/golangfirstproject/internal/storage"
)

func main() {
	cfg := config.ReadConfig()
	log := logger.Get(cfg.Debug)

	stor := storage.New()
	userService := service.NewUserService(stor)
	bookService := service.NewBookService(stor)
	serve := server.New(cfg, userService, bookService)
	if err := serve.Run(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
