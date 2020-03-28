package main

import "net/http"

// Server is the basic building block
type Server struct {
	port   string
	router *Router
}

// NewServer create a new instace of the Server struct
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

// Listen incoming request
func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
