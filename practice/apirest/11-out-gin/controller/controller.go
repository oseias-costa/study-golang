package controller

import "github.com/gin-gonic/gin"

var albums = []models.albums{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sara Vaughan", Artist: "Sara Vaughan", Price: 36.99},
}

func GetAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOk, albums)
}
