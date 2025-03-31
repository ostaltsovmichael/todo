package server

import (
	"log/slog"
	"net/http"
)

// This handler will be delete
func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test seccsess"))

}

func Server() {

	slog.Info("Server started")
	err := http.ListenAndServe("localhost:4000", http.HandlerFunc(testHandler))
	if err != nil {
		panic("Failed to connect to the server")
	}
}
