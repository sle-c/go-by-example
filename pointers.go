package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeropoint(point *int) {
	*point = 0
}

func main() {
	i := 1

	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeropoint(&i)
	fmt.Println("zeropoint:", i)

	fmt.Println("pointer:", &i)

}
