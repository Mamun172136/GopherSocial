package main

import (
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/social/internal/auth"
	"github.com/social/internal/db"
	"github.com/social/internal/env"
	"github.com/social/internal/ratelimiter"
	"github.com/social/internal/store"
	"github.com/social/internal/store/cache"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or couldn't load it. Using defaults.")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db : dbConfig{
		addr: env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5433/socialnetwork?sslmode=disable"),
		maxOpenCons: env.GetInt("DB_MAX_OPEN_CONNS", 30),
		maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
		maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		redisCfg: redisConfig{
			addr:    env.GetString("REDIS_ADDR", "localhost:6379"),
			pw:      env.GetString("REDIS_PW", ""),
			db:      env.GetInt("REDIS_DB", 0),
			enabled: env.GetBool("REDIS_ENABLED", false),
		},
		auth: authConfig{
			basic : basicConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_PASS", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "example"),
				exp:    time.Hour * 24 * 3, // 3 days
				iss:    "gophersocial",
			},
		},
		rateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUESTS_COUNT", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", true),
		},

	}

	db,err := db.New(cfg.db.addr,
	cfg.db.maxOpenCons,
	cfg.db.maxIdleConns,
	cfg.db.maxIdleTime,

)
if err != nil {
	log.Panic(err)
}

defer db.Close()
log.Println("db connect")

//cache
var rdb *redis.Client
	if cfg.redisCfg.enabled {
		rdb = cache.NewRedisClient(cfg.redisCfg.addr, cfg.redisCfg.pw, cfg.redisCfg.db)
		
		defer rdb.Close()
	}

store  := store.NewStorage(db)

cacheStorage := cache.NewRedisStorage(rdb)

// Rate limiter
	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.rateLimiter.RequestsPerTimeFrame,
		cfg.rateLimiter.TimeFrame,
	)


// Authenticator
	jwtAuthenticator := auth.NewJWTAuthenticator(
		cfg.auth.token.secret,
		cfg.auth.token.iss,
		cfg.auth.token.iss,
	)
	
app := &application{
	config: cfg,
	store: store,
	authenticator: jwtAuthenticator,
	cacheStorage:  cacheStorage,
	rateLimiter:   rateLimiter,
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