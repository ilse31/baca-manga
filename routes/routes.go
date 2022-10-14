package routes

import (
	"baca-manga/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.GET("/ping", controllers.Ping)
	router.GET("/home", controllers.Home)
	router.GET("/ivan", controllers.ControllerIvan)
	router.GET("/pokedex", controllers.BejjoController)
	router.GET("/enzd", controllers.EnzdController)
	router.GET("/wagyuA5", controllers.WagyuController)
}
