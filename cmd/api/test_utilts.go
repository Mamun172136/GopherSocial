package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/social/internal/auth"
	"github.com/social/internal/store"
	"github.com/social/internal/store/cache"
)

func newTestApplication(t *testing.T) *application {
	t.Helper()



	mockStore := store.NewMockStore()
	mockCacheStore := cache.NewMockStore()

	testAuth := &auth.TestAuthenticator{}

	// // Rate limiter
	// rateLimiter := ratelimiter.NewFixedWindowLimiter(
	// 	cfg.rateLimiter.RequestsPerTimeFrame,
	// 	cfg.rateLimiter.TimeFrame,
	// )

	return &application{

		store:         mockStore,
		cacheStorage:  mockCacheStore,
		authenticator: testAuth,
		// config:        cfg,

	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}