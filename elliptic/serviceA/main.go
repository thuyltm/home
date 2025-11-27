package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/servicea", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome Go Chi"))
	})
	r.Get("/servicea/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi Go Chi"))
	})
	http.ListenAndServe(":3000", r)
}
