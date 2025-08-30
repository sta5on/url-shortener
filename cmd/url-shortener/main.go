package main

import (
	"fmt"
	"url_shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqllite

	// TODO: init router: chi, render

	// TODO: run server:
}
