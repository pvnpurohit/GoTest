package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is init func")
}
func main() {
	x := 40
	y := 5
	fmt.Printf("x=%v \ny=%v\n", x, y)
	if x < 42 {
		fmt.Println("less than 42")
	} else if x == 42 {
		fmt.Println("equal to 42")
	} else {
		fmt.Println("greater than 42")
	}

	switch {
	case x < 42:
		fmt.Println("LESS than 42")
	case x == 42:
		fmt.Println("EQUAL to 42")
	case x > 42:
		fmt.Println("GREATER than 42")
	default:
		fmt.Println("DEFAULT case for x")
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("Counting to five: %v\n", i)
	}

	fmt.Println("Second type of for loop")
	for y < 10 {
		fmt.Printf("y is %v\n", y)
		y++
	}
	fmt.Println("Third type of for loop")
	for {
		fmt.Printf("y is %v\n", y)
		if y > 10 {
			break
		}
		y++
	}

	a := rand.Intn(250)
	fmt.Printf("Value of a: %v\n", a)
	if a > 0 && a < 100 {
		fmt.Println("a is between 0 and 100")
	} else if a > 101 && a < 200 {
		fmt.Println("a is between 101 and 200")
	} else if a > 201 && a < 250 {
		fmt.Println("a is between 201 and 250")
	}

	switch {
	case a > 0 && a < 100:
		fmt.Println("a is between 0 and 100")
	case a > 101 && a < 200:
		fmt.Println("a is between 101 and 200")
	case a > 201 && a < 250:
		fmt.Println("a is between 201 and 250")
	default:
		fmt.Println("This is default for a")
	}

	//slice defined...will be covered later
	list := []int{1, 2, 3, 4, 5}
	for index, value := range list {
		fmt.Printf("Value at index %v: %v\n", index, value)
	}

	//map defined....will be covered later
	m := map[string]int{
		"Praveen": 42,
		"Sanjo":   41,
	}
	for name, age := range m {
		fmt.Printf("Name: %v -> Age: %v\n", name, age)
	}

	age := m["Praveen"]
	fmt.Printf("Age of Praveen: %v\n", age)

	age = m["Bond"]
	fmt.Printf("Age of Bond: %v\n", age)

	//, ok Idiom
	if age, ok := m["Bond"]; !ok {
		fmt.Printf("Age of Bond is not available and hence its %v\n", age)
	}
}
