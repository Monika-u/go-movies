package db

import (
	"go-movies/config"
	models "go-movies/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// func GlobalMovieSearch(c *gin.Context, input models.Search) ([]models.Movie, error) {
// 	movies := make([]models.Movie, 0)
// 	var err error
// 	db := config.MariaDB

// 	// If neither SearchText nor Language is provided, return all data (default value)
// 	if input.SearchText == "" && input.Language == "" {
// 		err = db.Debug().Offset(input.Offset).
// 			Limit(input.Limit).
// 			Find(&movies).Error
// 	} else if input.SearchText != "" {
// 		// If SearchText is provided, search based on movie_name
// 		err = db.Debug().Where("movie_name LIKE ?", "%"+input.SearchText+"%").Offset(input.Offset).
// 			Limit(input.Limit).
// 			Find(&movies).Error
// 	} else {
// 		// If only Language is provided, search based on language
// 		err = db.Debug().Where("LOWER(language) = ?", strings.ToLower(input.Language)).
// 			Offset(input.Offset).
// 			Limit(input.Limit).
// 			Find(&movies).Error
// 	}

// 	if err != nil {
// 		// database query error
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"msg": "Error querying database",
// 		})
// 		return nil, err
// 	}

// 	// format the release_date to the desired format(if needed)
// 	for i := range movies {
// 		movies[i].ReleaseDateFormatted = movies[i].ReleaseDate.Format("01-02-2006")
// 	}

// 	return movies, nil
// }

func SearchAllMovies(c *gin.Context, input models.Search) ([]models.Movie, error) {
	movies := make([]models.Movie, 0)
	db := config.MariaDB

	err := db.Debug().Offset(input.Offset).
		Limit(input.Limit).
		Find(&movies).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Error querying database",
		})
		return nil, err
	}

	formatReleaseDates(&movies)

	return movies, nil
}

func SearchByMovieName(c *gin.Context, input models.Search) ([]models.Movie, error) {
	movies := make([]models.Movie, 0)
	db := config.MariaDB

	err := db.Debug().Where("movie_name LIKE ?", "%"+input.SearchText+"%").Offset(input.Offset).
		Limit(input.Limit).
		Find(&movies).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Error querying database",
		})
		return nil, err
	}

	formatReleaseDates(&movies)

	return movies, nil
}

func SearchByLanguage(c *gin.Context, input models.Search) ([]models.Movie, error) {
	movies := make([]models.Movie, 0)
	db := config.MariaDB

	err := db.Debug().Where("LOWER(language) = ?", strings.ToLower(input.Language)).
		Offset(input.Offset).
		Limit(input.Limit).
		Find(&movies).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Error querying database",
		})
		return nil, err
	}

	formatReleaseDates(&movies)

	return movies, nil
}

// format the release_date to the desired format(if needed)
func formatReleaseDates(movies *[]models.Movie) {
	for i := range *movies {
		(*movies)[i].ReleaseDateFormatted = (*movies)[i].ReleaseDate.Format("01-02-2006")
	}
}
