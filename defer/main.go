package main

import (
	"fmt"
)

/*
What is defer?
defer is a statement in Go that allows you to schedule a function call to be executed right before the enclosing function returns. The deferred function calls are executed in a last-in, first-out (LIFO) order, meaning the most recently deferred function is executed first, followed by the next, and so on. This mechanism is useful for ensuring cleanup actions, resource release, or finalization tasks are performed before exiting a function, regardless of the function's execution path (normal return, panic, or runtime error).

Why Use defer?
1. Resource Cleanup: Defer is commonly used for releasing resources like closing files, unlocking mutexes, or closing database connections. It helps prevent resource leaks by ensuring cleanup tasks are always executed.
2. Readability and Maintainability: By deferring cleanup actions next to setup code, the flow of execution becomes clearer and more readable. It keeps the cleanup logic close to where resources are acquired.
3. Error Handling: Defer ensures that even if an error occurs during a function's execution, necessary cleanup operations will still take place before the function exits.*/

func main() {
	// This statement will be executed last
	defer fmt.Println("Deferred statement 1")

	// These statements will be executed in order
	fmt.Println("Regular statement 1")
	fmt.Println("Regular statement 2")

	// Another deferred statement
	defer fmt.Println("Deferred statement 2")

	// These will execute before the deferred statements
	fmt.Println("Regular statement 3")
	fmt.Println("Regular statement 4")
}
