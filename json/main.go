package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Person struct defines a person with Name, Age, and Email fields
type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func main() {
    // Create an instance of Person
    person := Person{
        Name:  "John Doe",
        Age:   30,
        Email: "john.doe@example.com",
    }

    // Serialize the Person struct to JSON
    jsonData, err := json.Marshal(person)
    if err != nil {
        log.Fatalf("Error occurred during JSON marshaling. Error: %s", err.Error())
    }
    fmt.Println("JSON Output:", string(jsonData))

    // Now let's deserialize the JSON data back into a Person struct
    var personDecoded Person
    err = json.Unmarshal(jsonData, &personDecoded)
    if err != nil {
        log.Fatalf("Error occurred during JSON unmarshaling. Error: %s", err.Error())
    }
    fmt.Printf("Decoded JSON: %+v\n", personDecoded)
}
