package main

import (
	"log"

	"github.com/AlexDillz/Calc_server_yandex/internal/server"
)

func main() {
	app := server.New()
	log.Printf("Starting server on %s", app.Config.Addr)
	if err := app.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
