// if the program correctly prints the sorted slice after entering three distinct integers. 
// will be given for the second test execution, if the program correctly prints the sorted slice after entering 
// four distinct integers. *

package main

import (
	"fmt"
	"sort"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var arr []int
	var num int

	fmt.Print("Enter the integers: ")

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
        if input.Text() == "" {
            break
        } 
		
		num, _ = strconv.Atoi(input.Text())
		arr = append(arr, num)
		sort.Ints(arr)
        fmt.Println(arr)
    }
}