package main

import (
	"fmt"
	"net/http"
)

// HandleRoot handle "/" route
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from home")
}

// HandleHome handle "/api" route
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from api")
}
