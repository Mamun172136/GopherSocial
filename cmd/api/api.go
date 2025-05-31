package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/social/internal/auth"
	"github.com/social/internal/store"
	"github.com/social/internal/store/cache"
)

type application struct {
	config config
	store   store.Storage
	authenticator auth.Authenticator
	cacheStorage  cache.Storage
}

type config struct {
	addr string
	db 	dbConfig
	auth authConfig
	redisCfg    redisConfig
}

type redisConfig struct {
	addr    string
	pw      string
	db      int
	enabled bool
}

type authConfig struct{
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
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
			r.Use(app.AuthTokenMiddleware)
			r.Post("/", app.createPostHandler)
			
			r.Route("/{postID}", func (r chi.Router){
				r.Use(app.postsContextMiddleware)
				r.Get("/",app.getPostHandler)
				r.Delete("/", app.deletePostHandler)
				r.Patch("/",app.updatePostHandler)
			})
		})

		r.Route("/users", func(r chi.Router){
			r.Put("/create", app.registerUserHandler)
			r.Route("/{userID}", func(r chi.Router){
				r.Use(app.AuthTokenMiddleware)
				r.Use(app.userContextMiddleware)
				r.Get("/", app.getUserHandler)
				r.Put("/follow", app.followUserHandler)
				r.Put("/unfollow", app.unfollowUserHandler)
			})

			r.Group(func(r chi.Router) {
			
				r.Get("/feed", app.getUserFeedHandler)
			})
		})

		r.Route("/authentication", func(r chi.Router) {
			r.Post("/token", app.createTokenHandler)
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