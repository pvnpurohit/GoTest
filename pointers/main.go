package main

import "fmt"

func init() {
	fmt.Println("INIT function")
}

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%v\t%T\n", &x, &x)

	s := "Praveen"
	fmt.Printf("%v\t%T\n", &s, &s)

	y := &x
	fmt.Printf("%v\t%T\n", y, y)   // see the output as just y prints the address as it is assigned &x
	fmt.Printf("%v\t%T\n", &y, &y) // see the output as &y prints address where address of x (&x) is stored and see the type **int
	fmt.Println(y)
	fmt.Println(*y)  //prints the value stored at address (dereferencing a pointer)
	fmt.Println(*&x) //same as just using x

	*y = 43
	fmt.Println(x) //changing value at address; so x is changed as well

	//pass by value, pointers/ref types and mutability
	intDelta(&x)
	fmt.Println(x) //value of x is changed in the function intDelta

	//value semantics
	fmt.Println(addOne(x)) //adds 1 to x; x will be 45
	fmt.Println(x)         // old value of x=44

	//pointer semantics
	addOneP(&x)
	fmt.Println(x) // x will be 45

	//method sets
	d1 := dog{"DOGGY"}
	d1.walk()
	d1.run()
	//youngRun(d1)     //This fails as run can be called only by d2 as its pointer type "&dog"

	d2 := &dog{"NAAYI"} //this is same as dog{"NAAYI"}
	d2.walk()
	d2.run()

	youngRun(d2)

}

func intDelta(n *int) {
	*n = 44
}

func addOne(v int) int {
	return v + 1
}

func addOneP(v *int) {
	*v += 1
}

//method sets
type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is ", d.first, " and I am walking.")
}

func (d *dog) run() {
	d.first = "GOOD DOGGY"
	fmt.Println("My name is ", d.first, " and I am running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}
