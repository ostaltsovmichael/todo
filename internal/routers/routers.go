package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test seccsess"))
}

func GetRouters() *chi.Mux {

	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4000"},
		AllowCredentials: true,

		Debug: true,
	})

	router.Use(c.Handler)
	router.Get("/test", testHandler)

	return router
}
