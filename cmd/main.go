package main

import (
	"github.com/ub1vashka/golangfirstproject/internal/config"
	"github.com/ub1vashka/golangfirstproject/internal/logger"
	"github.com/ub1vashka/golangfirstproject/internal/server"
	"github.com/ub1vashka/golangfirstproject/internal/service"
	"github.com/ub1vashka/golangfirstproject/internal/storage"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.ReadConfig()
	log := logger.Get(cfg.Debug)

	err := storage.Migrations("postgres://user:password@localhost:5432/glfirst?sslmode=disable", "mugration")
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	stor := storage.New()
	userService := service.NewUserService(stor)
	bookService := service.NewBookService(stor)
	serve := server.New(cfg, userService, bookService)
	if err := serve.Run(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
