package main

// Book struct represents a book with an ID, title, author, and year of publication.
type Book struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Year   int     `json:"year"`
}

// books simulates a database of books.
var books []Book

// init initializes the books slice with some sample data.
func init() {
    books = append(books, Book{ID: "1", Title: "1984", Author: "George Orwell", Year: 1949})
    books = append(books, Book{ID: "2", Title: "Brave New World", Author: "Aldous Huxley", Year: 1932})
}
