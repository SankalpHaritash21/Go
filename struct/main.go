package main

import "fmt"

// Define the Book struct
type Book struct {
    Title       string
    Author      string
    Genre       string
    Year        int
    ISBN        string
}

func main() {
    // Create a new instance of the Book struct
    book := Book{
        Title:  "The Great Gatsby",
        Author: "F. Scott Fitzgerald",
        Genre:  "Fiction",
        Year:   1925,
        ISBN:   "9780743273565",
    }

    // Print out information about the book
    fmt.Println("Title:", book.Title)
    fmt.Println("Author:", book.Author)
    fmt.Println("Genre:", book.Genre)
    fmt.Println("Year:", book.Year)
    fmt.Println("ISBN:", book.ISBN)
}
