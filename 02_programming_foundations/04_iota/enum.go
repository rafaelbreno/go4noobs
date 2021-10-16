package main

import "fmt"

// Our type to define a day
type day int

const (
	// Just some days of the week
	//
	SUNDAY day = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

func (d day) ToString() string {
	switch d {
	case SUNDAY:
		return "sunday"
	case MONDAY:
		return "monday"
	case TUESDAY:
		return "tuesday"
	case WEDNESDAY:
		return "wednesday"
	case THURSDAY:
		return "thursday"
	case FRIDAY:
		return "friday"
	case SATURDAY:
		return "saturday"
	default:
		return ""
	}
}

func main() {
	monday := MONDAY

	fmt.Println(monday.ToString())
}
