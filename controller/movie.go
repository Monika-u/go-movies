package controller

import (
	"go-movies/db"
	models "go-movies/model"
	"go-movies/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

func GlobalMovieSearch(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "GlobalSearch", "controller")
	defer span.End()

	ctx = utils.SetContext(c, ctx)

	// Get the search input
	var input models.Search
	input.Offset = 0
	input.Limit = 20

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Error Invalid ReqBody",
		})
		return
	}

	// Perform the global movie search
	output, err := db.GlobalMovieSearch(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Error querying database",
		})
		return
	}
	if len(output) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Oops! No data with given input",
		})
		return
	}

	c.JSON(http.StatusOK, output)
}
