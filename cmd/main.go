package main

import (
	"fmt"

	"github.com/ub1vashka/golangfirstproject/internal/config"
)

func main() {
	cfg := config.ReadConfig

	fmt.Println(cfg)
}
