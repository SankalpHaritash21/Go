package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// getBooks responds with the list of all books as JSON.
func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

// getBookByID locates the book whose ID value matches the id
// parameter sent by the client, then returns that book as a response.
func getBookByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range books {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// postBook adds a new book from JSON received in the request body.
func postBook(c *gin.Context) {
    var newBook Book
    if err := c.BindJSON(&newBook); err != nil {
        return
    }
    books = append(books, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
    router := gin.Default()
    router.GET("/books", getBooks)
    router.GET("/books/:id", getBookByID)
    router.POST("/books", postBook)

    router.Run("localhost:8080")
}
