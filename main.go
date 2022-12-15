package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addBooks(c *gin.Context) {
	// BODY
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
func findBookID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Your book is not found")
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := findBookID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "Your book not found")
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing ID of BOOK"})
		return
	}

	book, err := findBookID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Your book is not found"})
	}
	book.Quantity = book.Quantity - 1
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello World!!!")
	})
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.PATCH("/checkout/", checkoutBook)
	router.POST("/add_book", addBooks)
	router.Run("localhost:8000")
	fmt.Println("Server is running on PORT 8000")
}
