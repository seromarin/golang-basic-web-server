package main

import "net/http"

// Server is the basic building block
type Server struct {
	port string
}

// NewServer create a new instace of the Server struct
func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

// Listen incoming request
func (s *Server) Listen() error {
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
