package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/social/internal/store"
)

type application struct {
	config config
	store   store.Storage
}

type config struct {
	addr string
	db 	dbConfig
	auth authConfig
}
type authConfig struct{
	basic basicConfig
}

type basicConfig struct{
	user string
	pass string
}

type dbConfig struct {
	addr 			string
	maxOpenCons		int
	maxIdleConns	int
	maxIdleTime		string
}

func (app *application) mount() *chi.Mux{
	// mux := http.NewServeMux()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// mux.HandleFunc("GET /v1/health", app.healthCheckerHandler)
	r.Route("/v1", func(r chi.Router){
		r.Get("/health", app.healthCheckerHandler)

		r.Route("/posts", func(r chi.Router){
			r.Post("/", app.createPostHandler)
			
			r.Route("/{postId}", func (r chi.Router){
				r.Use(app.postsContextMiddleware)
				r.Get("/",app.getPostHandler)
				r.Delete("/", app.deletePostHandler)
				r.Patch("/",app.updatePostHandler)
			})
		})

		r.Route("/users", func(r chi.Router){
			r.Route("/{userId}", func(r chi.Router){
				r.Use(app.userContextMiddleware)
				r.Get("/", app.getUserHandler)
				r.Put("/", app.followUserHandler)
				r.Put("/", app.unfollowUserHandler)
			})

			r.Group(func(r chi.Router) {
			
				r.Get("/feed", app.getUserFeedHandler)
			})
		})
	})

	// posts

	// users

	// auth
	return r
}

func (app *application) run(mux *chi.Mux) error {
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