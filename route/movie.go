package routes

import (
	"go-movies/controller"

	"github.com/gin-gonic/gin"
)

func CreateMoviesRoutes(server *gin.RouterGroup) {

	Group := server.Group("/public/v1/")
	{
		// Group.GET("movies", controller.GetAll)
		Group.POST("movies", controller.GlobalMovieSearch)
	}
}
