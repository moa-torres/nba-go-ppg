package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/moacirtorres/nba-go-ppg/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Printf("Servidor escutando na port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
