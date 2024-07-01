package server

import (
	"github.com/gin-gonic/gin"
	db "ordering-server/internal/database"
)

// Server all HTTP requests
type Server struct {
	router *gin.Engine
	store  *db.Store
}

func NewServer(store *db.Store) (*Server, error) {

	server := &Server{
		router: gin.Default(),
		store:  store,
	}
	// Remove from here
	server.RegisterRoutes()

	return server, nil
}

func (s *Server) RegisterRoutes() {
	// Register Routes
	s.userRoutes()
}

func (s *Server) RunServer(address string) error {
	return s.router.Run(address)
}
