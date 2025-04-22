package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moazan-waqar-hpe/gin-gorm-rest/config"
	"github.com/moazan-waqar-hpe/gin-gorm-rest/routes"
)

func main() {
	router := gin.New()
	routes.UserRoute(router)
	config.Connect()
	router.Run(":8080")
}
