package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, v := range nums {
		total += v
	}

	fmt.Println("total is", total)
	return total
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusplus(1, 2, 3)
	fmt.Println("1+2+3=", res)

	sum(1, 2, 3)

	numSlice := []int{1, 2, 3, 4, 5}
	sum(numSlice...)
}
