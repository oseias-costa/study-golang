package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "errors"
)

type book struct {
	ID		 string	`json:id`
	Title	 string	`json:title`
	Author   string	`json:author`
	Quantity int 	`json:quantity`
}

var books = []book {
	{ID: "1", Title: "Clean Architecture", Author: "Uncle Bob", Quantity: 8},
	{ID: "2", Title: "Programing Go Language", Author: "Other", Quantity: 4},
	{ID: "3", Title: "Clean Code", Author: "Uncle Bob", Quantity: 10},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context){
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main(){
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}

//curl https://t4nkvl-8080.csb.app/books --include --header "Content-Type: application/json" -d @body.json --request "POST"      