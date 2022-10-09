package main

import (
	"log"

	routes "baca-manga/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := InitRoute()
	log.Fatal(router.Run(":8080"))
}

func InitRoute() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	RouteApi := router.Group("/api")
	routes.Routes(RouteApi)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
	return router
}
