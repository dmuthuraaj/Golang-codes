package main

import (
	"fmt"
)

func main() {
	num := 2
	fmt.Printf("complement:%v", ^num)
	fmt.Printf("and:%v", 2&^9)
}
