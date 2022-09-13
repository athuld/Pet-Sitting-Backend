package app

import (
	"pet-sitting-backend/controllers/pets"
	sitterreqs "pet-sitting-backend/controllers/sitter_reqs"
	sitterresps "pet-sitting-backend/controllers/sitter_resps"
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

	// Sitter Requests
	router.POST("/api/user/sitter_req/add_request", sitterreqs.AddRequest)
	router.GET("/api/user/sitter_req/get_all/active", sitterreqs.GetActiveRequests)
	router.GET("/api/user/sitter_req/get_all/inactive", sitterreqs.GetInActiveRequests)
	router.DELETE("/api/user/sitter_req/delete_request", sitterreqs.DeleteRequest)

	// Sitters
	router.GET("/api/user/sitter/get_all/by_pincode", users.GetActiveRequestsFromPincode)

    // Response
    router.POST("/api/user/sitter/response",sitterresps.AddResponse)
    router.GET("/api/user/sitter/responses/by_id",sitterresps.GetResponsesById)
    router.PATCH("/api/user/sitter/response/accept",sitterresps.AcceptResponse)
}
