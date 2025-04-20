package main

import "net/http"

func (app *applicaton) healthCheckerHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("ok"))
}