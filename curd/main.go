package main

import (
	"fmt"
)

// User struct defines the user with ID, Name, and Email
type User struct {
    ID    int
    Name  string
    Email string
}

var users = []User{} // Slice to store users
var nextID = 1       // Variable to generate unique IDs for new users

// Function to create a new user and add to the slice
func createUser(name, email string) User {
    user := User{
        ID:    nextID,
        Name:  name,
        Email: email,
    }
    nextID++
    users = append(users, user)
    return user
}

// Function to list all users
func listUsers() []User {
    return users
}

// Function to update an existing user
func updateUser(id int, name, email string) error {
    for i, user := range users {
        if user.ID == id {
            users[i].Name = name
            users[i].Email = email
            return nil
        }
    }
    return fmt.Errorf("User not found")
}

// Function to delete a user
func deleteUser(id int) error {
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("User not found")
}

// Main function to demonstrate CRUD operations
func main() {
    fmt.Println("Initial Users:", listUsers())

    // Create Users
    user1 := createUser("John Doe", "john@example.com")
    fmt.Println("Added User:", user1)

    user2 := createUser("Jane Doe", "jane@example.com")
    fmt.Println("Added User:", user2)

    // List Users
    fmt.Println("List of Users:", listUsers())

    // Update User
    err := updateUser(user1.ID, "John D.", "john.d@example.com")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Updated User:", listUsers())
    }

    // Delete User
    err = deleteUser(user2.ID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("After Deletion:", listUsers())
    }
}
