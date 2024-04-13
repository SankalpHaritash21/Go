package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current day of the week
	today := time.Now().Weekday()
	fmt.Print("current time is: ",today)
	// Switch-case statement with an expression
	switch today {
	case time.Monday:
		fmt.Println("Today is Monday.")
	case time.Tuesday:
		fmt.Println("Today is Tuesday.")
	case time.Wednesday, time.Thursday:
		fmt.Println("Today is either Wednesday or Thursday.")
	case time.Friday:
		fmt.Println("Today is Friday.")
	default:
		fmt.Println("It's a weekend day.")
	}

	// Switch-case with expressions without a condition
	switch {
	case today == time.Saturday || today == time.Sunday:
		fmt.Println("It's the weekend!")
	case today == time.Monday:
		fmt.Println("Back to work!")
	default:
		fmt.Println("It's a weekday.")
	}

	// Switch-case with fallthrough
	switch today {
	case time.Saturday:
		fmt.Println("It's Saturday.")
		fallthrough
	case time.Sunday:
		fmt.Println("It's either Saturday or Sunday.")
	default:
		fmt.Println("It's a weekday.")
	}
}
