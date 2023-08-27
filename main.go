package main

import (
	"go-movies/config"
	routes "go-movies/route"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// r := mux.NewRouter()
	config.InitializeDB()

	// Staring the http server using GIN
	// with port 8019
	// and creating the required routes
	// binding.Validator = new(config.CustomValidator)
	server := gin.New()
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch},
		AllowHeaders: []string{
			"Accept",
			"Content-Type",
			"contentType",
			"Content-Length",
			"Accept-Encoding",
			"Client-Security-Token",
			"X-CSRF-Token",
			"X-Auth-Token",
			"processData",
			"Authorization",
			"Access-Control-Request-Headers",
			"Access-Control-Request-Method",
			"Connection",
			"Host",
			"Origin",
		},
		MaxAge: 12 * time.Hour,
	}))

	group := server.Group("/")
	routes.CreateMoviesRoutes(group)

	// Public Health Check URL
	group.GET("/public/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	if err := server.Run(":8019"); err != nil {
		panic(err)
	}

}
