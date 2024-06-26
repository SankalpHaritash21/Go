package main

import "fmt"

// Public variable
var PublicVar int = 10

// Public function
func PublicFunc() {
    fmt.Println("This is a public function")
}
// Public variable or function always have Capital letter
// Privet variable always have lower letters

func main() {
    fmt.Println("Hello, World!")
    printArrays()
    Initialization()

    // Accessing public variable and function
    fmt.Println("PublicVar:", PublicVar)
    PublicFunc()

    Datatype()
    Print()
}

//? Line 1: In Go, every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package.

//? Line 2: import ("fmt") lets us import files included in the fmt package.

//? Line 3: A blank line. Go ignores white space. Having white spaces in code makes it more readable.

//? Line 4: func main() {} is a function. Any code inside its curly brackets {} will be executed.

//? Line 5: fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!".