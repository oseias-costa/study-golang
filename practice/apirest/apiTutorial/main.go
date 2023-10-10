package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:id`
	Title    string `json:title`
	Author   string `json:author`
	Quantity int    `json:quantity`
}

var books = []book{
	{ID: "1", Title: "Clean Architecture", Author: "Uncle Bob", Quantity: 8},
	{ID: "2", Title: "Programing Go Language", Author: "Other", Quantity: 4},
	{ID: "3", Title: "Clean Code", Author: "Uncle Bob", Quantity: 10},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New(("book not found"))
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missin id query parameter"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATH("/checkout", checkoutBook)
	router.Run("localhost:8080")
}

//curl https://t4nkvl-8080.csb.app/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
