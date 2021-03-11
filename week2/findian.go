package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Scan(&input)
	
	if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not found!")
	}
}
