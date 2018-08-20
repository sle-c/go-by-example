package main

import "fmt"

func main() {
	s := make([]string, 3)

	fmt.Println("Empty:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[1])

	fmt.Println("len:", len(s))

	s = append(s, "d", "e", "f")

	fmt.Println("append:", s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("slice1:", l)

	l = s[:5]
	fmt.Println("slice2:", l)

	l = s[2:]
	fmt.Println("slice3:", l)

	t := []string{"a", "b", "c"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
