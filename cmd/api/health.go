package main

import "net/http"

func (app *aplication) healthHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}