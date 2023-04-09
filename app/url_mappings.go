package app

import (
	"github.com/KodjoTouglo/bookstore_users/controllers/ping"
	"github.com/KodjoTouglo/bookstore_users/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
