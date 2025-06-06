package main

import "net/http"

func (app *application) healthCheckerHandler(w http.ResponseWriter, r *http.Request){

	data:= map[string]string{
		"status": "ok",
		"env" : "",
		"version": "version",
	}
	if err := writeJSON(w, http.StatusOK, data); err != nil{
		writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
	
}