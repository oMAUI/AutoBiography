package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewRoute() *chi.Mux{
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			if _, errWrite := w.Write([]byte("hello")); errWrite != nil {
				return
			}
	})

	return router
}