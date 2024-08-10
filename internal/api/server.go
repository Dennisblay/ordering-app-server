package api

import (
	db "github.com/Dennisblay/ordering-app-server/internal/database/sqlc"
	"github.com/gin-gonic/gin"
)

// Server all HTTP requests
type Server struct {
	router *gin.Engine
	store  db.Store
}

func NewServer(store db.Store) (*Server, error) {

	server := &Server{
		router: gin.Default(),
		store:  store,
	}
	server.RegisterRoutes()

	return server, nil
}

func (s *Server) RegisterRoutes() {
	// Register Routes
	s.userRoutes()
}

// ConfigureCORS Configure CORS
func (s *Server) ConfigureCORS() {

}

func (s *Server) RunServer(address string) error {
	return s.router.Run(address)
}
