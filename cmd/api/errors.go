package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal error method: %s url path:%s %s",r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}



func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request  method: %s url path:%s %s",r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusBadRequest, err.Error())
}



func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found  method: %s url path:%s %s",r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusNotFound, "not found")
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found  method: %s url path:%s %s",r.Method, r.URL.Path, err)


	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found  method: %s url path:%s %s",r.Method, r.URL.Path, err)


	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application)forbiddenResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found  method: %s url path:%s %s",r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusForbidden, "forbidden")
}




