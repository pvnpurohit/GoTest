package main

import "fmt"

func init() {
	fmt.Println("INIT function")
}

func main() {
	//array creation; [3] tells its fixed size
	a := [3]int{1, 2, 3}
	fmt.Println(a)

	//array size is auto determined; [...] tells that
	b := [...]string{"Hello", "Praveen"}
	fmt.Println(b)

	//diff way of defining array and assigning values later
	var c [2]int
	c[0] = 4
	c[1] = 5
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)
	fmt.Println(len(a))

	//slice creation; [] denotes size is not fixed
	xs := []string{"Hello", "World"}
	fmt.Println(xs)

	s := []string{"one", "two", "three"}
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	for i, v := range s {
		fmt.Printf("index %v -> value %v\n", i, v)
	}

	//we are ignoring index here denoted by _
	for _, v := range s {
		fmt.Printf("Value %v\n", v)
	}

	//append function
	s = append(s, "Four", "5")
	fmt.Println(s)

	//slicing the slice [inclusive:exclusive]
	fmt.Println(s[0:3])
	//Index starting from 0
	fmt.Println(s[:4])
	//Till the last index
	fmt.Println(s[1:])

	//deleting using slice...deleting 4th element(value at 3rd index)
	s = append(s[:3], s[4:]...)
	fmt.Println(s)

	//slice creation using make function; 10 denotes size and 0 denotes no initial values in the list
	slice := make([]int, 0, 10)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(slice)
	slice = append(slice, 10, 11, 12, 13, 14, 15)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//multidimenstional slices
	jb := []string{"james", "bond", "martini", "chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Vodka", "Vanilla"}
	fmt.Println(jb)
	fmt.Println(jm)
	ms := [][]string{jb, jm}
	fmt.Println(ms)
	fmt.Printf("%T\n", ms)

	//assigning one slice to another; both are referring to same underlying array
	slice1 := slice
	fmt.Println(slice)
	fmt.Println(slice1)
	//This changes both the arrays
	slice[14] = 16
	fmt.Println(slice)
	fmt.Println(slice1)

	//Copying the values to avoid changing both the slices
	slice2 := make([]int, 15)
	copy(slice2, slice)
	slice[14] = 15
	fmt.Println(slice)
	fmt.Println(slice2)

}
