package main

import (
	"fmt"
)

func main() {
	x := 40
	y := 5
	fmt.Printf("x=%v \n y=%v\n", x, y)
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

	for y < 10 {
		fmt.Printf("y is %v\n", y)
		y++
	}

	for {
		fmt.Printf("y is %v\n", y)
		if y > 10 {
			break
		}
		y++
	}
}
