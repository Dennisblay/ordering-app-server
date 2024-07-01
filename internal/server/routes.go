package server

func (s *Server) userRoutes() {
	// User routes
	userGroup := s.router.Group("/users")
	{
		userGroup.GET("/", s.getAllUsersController)
		userGroup.GET("/:id", s.getUserByIdController)
		userGroup.GET("/email/:email", s.getUserByEmailController)
		userGroup.GET("/phone/:phone", s.getUserByPhoneController)
		userGroup.POST("/login", s.getUserByEmailAndPasswordController)
		userGroup.POST("/", s.createUserController)
		userGroup.PUT("/:id/name", s.updateUserNameController)
		userGroup.PUT("/:id/email", s.updateUserEmailController)
		userGroup.PUT("/:id/phone", s.updateUserPhoneController)
		userGroup.PUT("/:id/address", s.updateUserAddressController)
		userGroup.PUT("/:id/password", s.updateUserPasswordController)
		userGroup.DELETE("/:id", s.deleteUserController)
	}
}
