package app

import (
	"github.com/KodjoTouglo/bookstore_users/controllers/ping"
	"github.com/KodjoTouglo/bookstore_users/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.PUT("/users/update/:user_id", users.UpdateUser)
	router.PATCH("/users/update/:user_id", users.UpdateUser)
	router.DELETE("/users/delete/:user_id", users.DeleteUser)
}
