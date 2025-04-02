package server

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ostaltsovmichael/todolist/internal/routers"
)

func Server() {

	port := os.Getenv("PORT")
	slog.Info("Server started")
	err := http.ListenAndServe("localhost:"+port, routers.GetRouters())
	if err != nil {
		panic("Failed to connect to the server")
	}
}
