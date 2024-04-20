package main

import (
	"fmt"
	"time"
)

func main() {
    // Example booking time as a string
    bookingTimeString := "2024-04-20 15:30:00"

    // Layout for parsing the time (Note: the layout must use the reference time: Mon Jan 2 15:04:05 MST 2006)
    layout := "2006-01-02 15:04:05"

    // Parsing the booking time
    bookingTime, err := time.Parse(layout, bookingTimeString)
    if err != nil {
        fmt.Println("Error parsing booking time:", err)
        return
    }

    // Getting the current time
    currentTime := time.Now()

    // Calculating the difference between the booking time and the current time
    durationUntilBooking := bookingTime.Sub(currentTime)

    // Check if the booking is within the next 48 hours
    if durationUntilBooking <= 48*time.Hour {
        // Printing a reminder
        fmt.Printf("Reminder: You have a booking in %v hours.\n", durationUntilBooking.Hours())
    } else {
        // Printing a message if the booking is more than 48 hours away
        fmt.Println("Your booking is more than 48 hours away.")
    }
}
