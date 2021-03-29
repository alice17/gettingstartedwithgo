// 3 points will be given if a program is written.
// 2 points will be given if the program compiles correctly.
// 5 points will be given if the program correctly prints a JSON object with keys ("name", "address") 
// and they are associated with the name and address that was entered.

package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
    Name string
    Address string
}

func main() {
	var name string
	var address string

	fmt.Printf("Enter a name:")
	fmt.Scan(&name)
	fmt.Printf("Enter an address:")
	fmt.Scan(&address)

	p := Person{name, address}
	b, _ := json.Marshal(p)
	
	fmt.Printf(string(b))
}