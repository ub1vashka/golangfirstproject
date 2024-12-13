package main

import (
	"fmt"

	"github/golangfirstproject/internal/config"
)

func main() {
	cfg := config.ReadConfig

	fmt.Println(cfg)
}
