package main

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/ostaltsovmichael/todolist/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	server.Server()
}
