package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status: available\n")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "app version: %s\n", version)
}
