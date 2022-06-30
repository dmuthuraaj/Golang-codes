package main

import (
	"fmt"
)

func main() {
	var myMap = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	for _key, _value := range myMap {
		fmt.Println(_key, ":", _value)
	}
}
