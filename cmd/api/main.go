package main

import "log"

func main() {

	cfg := config{
		addr: ":8080",
	}

	app := &applicaton{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}

// package main

// import (
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// )

// type config struct {
// 	addr string
// }

// type applicaton struct {
// 	config config
// }

// func (app *applicaton) mount() *chi.Mux {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)

// 	r.Route("/v1", func(r chi.Router) {
// 		r.Get("/health", app.healthCheckerHandler)
// 	})

// 	return r
// }

// func (app *applicaton) run(mux *chi.Mux) error {
// 	srv := &http.Server{
// 		Addr:         app.config.addr,
// 		Handler:      mux,
// 		WriteTimeout: time.Second * 30,
// 		ReadTimeout:  time.Second * 10,
// 		IdleTimeout:  time.Minute,
// 	}

// 	log.Printf("serve has started at %s", app.config.addr)

// 	return srv.ListenAndServe()
// }

// func main() {
// 	cfg := config{
// 		addr: ":8080",
// 	}

// 	app := &applicaton{
// 		config: cfg,
// 	}

// 	mux := app.mount()
// 	log.Fatal(app.run(mux))
// }