package routes

import (
	"github.com/gin-gonic/gin"
)

func createMoviesRoutes(server *gin.RouterGroup) {

	Group := server.Group("/public/v1/")
	{
		Group.GET("movies", controller.GetAll)
	}
}
