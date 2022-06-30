package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("sum:", total)
}
func main() {
	sum(1, 2, 3)
	no := []int{2, 4, 5}
	sum(no...)
}
