package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("hell start")
	case time.Tuesday:
		fmt.Println("Crossed hell but lot more to go")
	case time.Wednesday:
		fmt.Println("center for who have leave on saturday")
	case time.Thursday:
		fmt.Println("Thinking just a day to go")
	case time.Friday:
		fmt.Println("Just a day to go")
	case time.Saturday:
		fmt.Println("At last its weekend")
	case time.Sunday:
		fmt.Println("Hell its so fast")
	default:
		fmt.Println("Whatever it is")
	}
}
