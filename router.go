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

// FindHandler map our request with a handler
func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello world fron go server")
}
