package main

import (
	"fmt"
	"net/http"
)

// Router is a data type for the array of routes
type Router struct {
	rules map[string]http.HandlerFunc
}

// NewRouter instanciate the router
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello world fron go server")
}
