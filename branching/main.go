package main

import (
	"fmt"
	"time"
)

func main() {
	//day := "Tuesday"

	now := time.Now()

	day := now.Weekday().String()
	//Weekday()
	// func (t Time) Weekday() Weekday
	// type Weekday int
	/*
		A Weekday specifies a day of the week (Sunday = 0, ...).

		const (
			Sunday Weekday = iota
			Monday
			Tuesday
			Wednesday
			Thursday
			Friday
			Saturday
		)

	*/

	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	case "Thursday":
		fmt.Println("Today is Thursday")
	case "Friday":
		fmt.Println("Today is Friday")
	case "Saturday":
		fmt.Println("Today is Saturday")
	case "Sunday":
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Unknown day")
	}
}
