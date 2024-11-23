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
}
