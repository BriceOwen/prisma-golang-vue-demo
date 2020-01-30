package app

import (
	"log"
	"net/http"
	"prisma-golang-vue/src/system/router"
)

type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}

// Init all vals
func (s *Server) Init(port string) {
	log.Println("Initializing server...")
	s.port = ":" + port
}

// Start the server
func (s *Server) Start() {
	log.Println("Starting server on port " + s.port)

	r := router.NewRouter()

	r.Init()

	http.ListenAndServe(s.port, r.Router)
}
