package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "CD 1 J&M", Artist: "Jorge e Mateus", Price: 20.56},
	{ID: "2", Title: "CD 2 B&M", Artist: "Bruno e Marroni", Price: 9.59},
	{ID: "3", Title: "CD 3 GL", Artist: "Gustavo Lima", Price: 19.99},
	{ID: "4", Title: "CD 4 MM", Artist: "Maraia e Maraisa", Price: 1.96},
}

func main() {
	r := gin.Default()

	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumByID)
	r.POST("/albums", addAlbum)
	r.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Printf("Error type Json invalid %s", err.Error())
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			// chage here
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}
