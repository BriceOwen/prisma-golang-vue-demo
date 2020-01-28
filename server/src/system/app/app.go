package app

import (
	"log"
	"net/http"
)

type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}

// Init all vals
func (s *Server) Init() {
	log.Println("Initializing server...")
	s.port = ":4000"
}

// Start the server
func (s *Server) Start() {
	log.Println("Starting server!")
	http.ListenAndServe(s.port, nil)
}
