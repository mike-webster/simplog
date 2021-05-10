package main

import (
	"github.com/mike-webster/simplog/env"
	"github.com/mike-webster/simplog/router"
)

func main() {
	cfg, err := env.GetConfig()
	if err != nil {
		panic(err)
	}
	router.Start(cfg)
}
