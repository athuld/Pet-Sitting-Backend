package app

import "pet-sitting-backend/controllers/users"

func mapUrls() {
	router.POST("/api/user/register", users.Register)
	router.POST("/api/user/login", users.Login)
	router.GET("/api/user/logout", users.Logout)
}
