package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/social/internal/store"
)

type applicaton struct {
	config config
	store   store.Storage
}

type config struct {
	addr string
}

func (app *applicaton) mount() *chi.Mux{
	// mux := http.NewServeMux()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// mux.HandleFunc("GET /v1/health", app.healthCheckerHandler)
	r.Route("/v1", func(r chi.Router){
		r.Get("/health", app.healthCheckerHandler)
	})

	// posts

	// users

	// auth
	return r
}

func (app *applicaton) run(mux *chi.Mux) error {
	// mux := http.NewServeMux()


	srv := &http.Server{
		Addr : app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second *30,
		ReadTimeout: time.Second*10,
		IdleTimeout: time.Minute,
	}

	log.Printf("serve has started at %s", app.config.addr)

	return srv.ListenAndServe()
}