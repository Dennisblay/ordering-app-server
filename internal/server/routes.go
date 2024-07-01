package server

func (s *Server) userRoutes() {
	s.router.GET("/user", s.getAllUsersController)
}
