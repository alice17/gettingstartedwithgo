package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Scan(&input)
	
	if strings.Contains(input, "i") && strings.Contains(input, "a") && strings.Contains(input, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not found!")
	}
}
