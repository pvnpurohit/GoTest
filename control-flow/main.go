package main

import (
	"fmt"
)

func main() {
	x := 40
	y := 5
	fmt.Printf("z=%v \n y=%v\n", x, y)
	if x < 42 {
		fmt.Println("less than 42")
	} else {
		fmt.Println("equal to or greater than 42")
	}
}
