package main

import (
	"fmt"
	"os"
)

func main() {
	user := os.Getenv("USER")
	fmt.Println(user)
}
