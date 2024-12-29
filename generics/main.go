package main

import (
	"fmt"
)

func init() {
	fmt.Println("INIT function")
}

func main() {
	var x myInt = 11
	fmt.Println(addI(10, 20))
	fmt.Println(addF(10.2, 20.2))
	fmt.Println(addT(20, 30))
	fmt.Println(addT(20.3, 30.3))

	fmt.Println(addNums(x, 8))
}

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

//Generic function which takes both types...we are avoiding DRY (Do not Repeat Yourself) by having a generic function

func addT[T int | float64](a, b T) T {
	return a + b
}

// Adding generic types to interface
type numbers interface {
	//int | float64
	~int | ~float64 // ~ covers all aliases for respective types
	//constraints.Integer | constraints.Float    // we are using Type interfaces provided by GoLang from constraints pkg (look in GoLang spec for details about these type interfaces)
}

func addNums[T numbers](a, b T) T { //using the type numbers
	return a + b
}

// alias type
type myInt int
