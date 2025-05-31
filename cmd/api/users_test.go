package main

import (
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T){
	app := newTestApplication(t)
	mux := app.mount()

	testToken,err := app.authenticator.GenerateToken(nil)
	if err != nil{
		t.Fatal(err)
	}
	t.Run("should not allow unauthenticated user", func(t *testing.T){
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1",nil)
		if err != nil{
			t.Fatal(err)
		}
		rr := executeRequest(req, mux)

		if rr.Code != http.StatusUnauthorized{
			t.Errorf("expected response code to be %d and we got %d",http.StatusUnauthorized, rr.Code)
		}
	})

	t.Run("should  allow authenticated user", func(t *testing.T){
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1",nil)
		if err != nil{
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := executeRequest(req, mux)

		if rr.Code != http.StatusUnauthorized{
			t.Errorf("expected response code to be %d and we got %d",http.StatusUnauthorized, rr.Code)
		}
	})
}