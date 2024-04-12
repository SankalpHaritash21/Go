package main

import "fmt"

type User struct {
    Name   string
    Age    int
    Status string
}

func main() {
    users := []User{
        {"Alice", 30, "active"},
        {"Bob", 20, "inactive"},
        {"Carol", 25, "active"},
        {"Dave", 35, "active"},
        {"Eve", 40, "inactive"},
    }

    for _, user := range users {
        // Using if with an initialization statement
        if ageLimit := 25; user.Age > ageLimit && user.Status == "active" {
            // Processing only users over age 25 and with "active" status
            fmt.Printf("%s meets the criteria.\n", user.Name)
        }
    }
}
