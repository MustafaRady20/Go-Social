package main

import "net/http"

func (app *aplication) healthHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"Status": "Ok",
		"Env":    app.config.Env,
	}
	if err := WriteJson(w, http.StatusOK, data); err != nil {
		WriteJsonError(w, http.StatusInternalServerError, err.Error())
	}

}
