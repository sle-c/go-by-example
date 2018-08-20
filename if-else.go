package main

import "fmt"

func main() {

	if num := 7; num%2 == 0 {
		fmt.Printf("%d is even", num)
	} else {
		fmt.Printf("%d is odd", num)
	}
}
