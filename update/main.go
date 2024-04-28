package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "yourusername"
    password = "yourpassword"
    dbname   = "yourdbname"
)

func main() {
    // Connection string
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    // Open the connection
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Check the connection
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected!")

    // Update data
    if err := updateUserData(db, 1, "newemail@example.com"); err != nil {
        log.Fatalf("Error updating user data: %v", err)
    }
    fmt.Println("Update was successful")
}

func updateUserData(db *sql.DB, userId int, newEmail string) error {
    // SQL statement
    sqlStatement := `
    UPDATE users
    SET email = $1
    WHERE id = $2;`
    
    // Execute SQL statement
    _, err := db.Exec(sqlStatement, newEmail, userId)
    if err != nil {
        return err
    }
    return nil
}
