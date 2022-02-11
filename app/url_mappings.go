package app

import "github.com/khalil-farashiani/microservice_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.GET("/users/search", controllers.SearchUser)
	router.POST("/users/", controllers.CreateUser)
	router.GET("/users/:user_id", controllers.GetUser)
}
