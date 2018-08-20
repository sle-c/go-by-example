package main

import "fmt"

func fac(num int) int {
	if num == 0 {
		return 1
	}

	return num * fac(num-1)
}

func main() {
	fmt.Println(fac(10))
}
