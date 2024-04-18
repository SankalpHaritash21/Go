package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// User defines a user type for JSON parsing
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    // Example 1: Basic Type Conversion
    var i int = 42
    var f float64 = float64(i)
    var s string = strconv.Itoa(i)
    fmt.Println(f, s) // Outputs: 42 42

    // Example 2: Error Handling in Type Conversion
    var str string = "100"
    var num int
    var err error
    num, err = strconv.Atoi(str)
    if err != nil {
        fmt.Println("Error converting string to int:", err)
    } else {
        fmt.Println("Converted integer:", num) // Outputs: Converted integer: 100
    }

    // Example 3: JSON to Struct and Struct to JSON
    jsonStr := `{"id":1, "name":"John Doe"}`
    var user User
    err = json.Unmarshal([]byte(jsonStr), &user)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
    } else {
        fmt.Println("JSON parsed into struct:", user) // Outputs: JSON parsed into struct: {1 John Doe}
    }

    newUser := User{ID: 2, Name: "Jane Doe"}
    var jsonData []byte
    jsonData, err = json.Marshal(newUser)
    if err != nil {
        fmt.Println("Error converting struct to JSON:", err)
    } else {
        fmt.Println("Struct converted to JSON:", string(jsonData)) // Outputs: Struct converted to JSON: {"id":2,"name":"Jane Doe"}
    }
}
