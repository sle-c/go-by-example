package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("person", person{"Bob", 20})
	fmt.Println("person", person{name: "Gother", age: 22})
	fmt.Println("person", person{name: "Francis"})

	p := person{name: "gogo", age: 33}

	fmt.Println("name", p.name)
	fmt.Println("age", p.age)
}
