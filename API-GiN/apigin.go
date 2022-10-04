package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Book     string `json:"book"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "12", Book: "English", Author: "Ali", Quantity: 33},
	{ID: "5", Book: "Maths", Author: "Ahmed", Quantity: 25},
}

func getbooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookId(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
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
	return nil, errors.New("Error Found")

}

func createbooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func checkoutbook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing THe Parameters"})
		return
	}
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Missing THe Actual 11Output"})
		return

	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book Not available"})

	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}
func main() {
	router := gin.Default()
	router.GET("/books", getbooks)
	router.POST("/createbooks", createbooks)
	router.PATCH("/checkout", checkoutbook)
	router.GET("/books:id", getBookId)
	router.Run("localhost:8081")
}
