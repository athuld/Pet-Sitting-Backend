package app

import (
	"pet-sitting-backend/controllers/pets"
	"pet-sitting-backend/controllers/users"
)

func mapUrls() {

	// User Requests
	router.POST("/api/user/register", users.Register)
	router.POST("/api/user/login", users.Login)
	router.GET("/api/user/logout", users.Logout)
	router.POST("/api/user/add_details", users.AddUserDetails)
	router.GET("/api/user/get_details", users.GetUserDetails)

	// Pets Requests
	router.POST("/api/user/pet/add_pet", pets.AddPet)
	router.DELETE("/api/user/pet/delete_pet", pets.DeletePet)
	router.GET("/api/user/pet/get_all", pets.GetAllPets)
}
