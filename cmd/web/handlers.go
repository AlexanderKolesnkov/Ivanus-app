package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

	w.Write([]byte("Home of Ivan Shron"))
}
