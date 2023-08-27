package db

import (
	"go-movies/config"
	models "go-movies/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GlobalMovieSearch(c *gin.Context, input models.Search) ([]models.Movie, error) {
	movies := make([]models.Movie, 0)
	var err error
	db := config.MariaDB

	query := db.Debug()

	if input.SearchText != "" {
		// Search by movie name
		query = query.Where("movie_name LIKE ?", "%"+input.SearchText+"%")
	}

	if input.Language != "" {
		// Search by language
		query = query.Where("LOWER(language) = ?", strings.ToLower(input.Language))
	}

	if input.Rating != 0 {
		// Search by rating
		query = query.Where("rating >= ?", input.Rating)
	}

	err = query.Offset(input.Offset).
		Limit(input.Limit).
		Find(&movies).Error

	if err != nil {
		//  database query error
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Error querying database",
		})
		return nil, err
	}

	return movies, nil
}
