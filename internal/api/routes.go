package api

func (s *Server) userRoutes() {
	// User routes
	userGroup := s.router.Group("/users")
	{
		userGroup.GET("/", s.getUsersController)
		userGroup.GET("/:id", s.getUserByIDController)
		userGroup.GET("/", s.getUserControllerByEmailOrPhone)
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
