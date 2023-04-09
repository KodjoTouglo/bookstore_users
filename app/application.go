package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	MapUrls()
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
