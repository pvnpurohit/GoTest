package main

import (
	"fmt"
	"sort"
)

func init() {
	fmt.Println("INIT function")
}

type person struct {
	first string
	age   int
}

func main() {
	si := []int{34, 66, 2, 43, 99, 12, 6}
	fmt.Println(si)
	ss := []string{"Test", "ABC", "Xylo", "Pqrs"}
	fmt.Println(ss)
	sort.Ints(si)
	fmt.Println(si)
	sort.Strings(ss)
	fmt.Println(ss)

	p1 := person{
		"Praveen",
		42,
	}
	p2 := person{
		"Sangeetha",
		41,
	}
	p3 := person{
		"Advika",
		11,
	}

	people := []person{p1, p2, p3}
	fmt.Println(people)

}
