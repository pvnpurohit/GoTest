package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func init() {
	fmt.Println("INIT function")
}

func main() {
	foo()
	bar("Praveen")
	s := aloha("Purohit")
	fmt.Println(s)
	s1, n := dogYears("Praveen", 42)
	fmt.Println(s1, n)

	total1 := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Sum is:", total1)
	xi := []int{11, 12, 13, 14, 15, 16, 17, 18}
	total2 := sum(xi...) // slice unfurling
	fmt.Println("Sum is:", total2)

	defer foo1() //DEFER example
	bar1()

	//METHODS example
	p1 := person{
		first: "Praveen",
	}
	p2 := person{
		first: "Sangeetha",
	}
	p1.speak()
	p2.speak()

	sa1 := secretAgent{
		person: person{
			first: "Advika",
		},
		ltk: true,
	}
	sa1.speak()

	//calling interface func; note that we are passing types which are attached to speak method and speak method is added to interface
	saySomething(sa1)
	saySomething(p1)

	//STRINGER interface; just prints the type differently
	b := book{
		title: "I am confused",
	}

	var y count = 42

	fmt.Println(b)
	fmt.Println(y)

	//using log pkg instead of fmt
	log.Println(b)
	log.Println(y)

	//WRAPPER function call
	logInfo(b)
	logInfo(y)

	//WRITING to a file
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s2 := []byte("Hello gophers!")
	_, err = f.Write(s2)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	//WRITING to a buffer
	byts := bytes.NewBufferString("Hello ")
	fmt.Println(byts)
	fmt.Println(byts.String()) //calling the stringer interface
	byts.WriteString("gophers!")
	//fmt.Println(byts)
	fmt.Println(byts.String())
	byts.Reset()
	byts.WriteString("This is very confusing ")
	fmt.Println(byts.String())
	byts.WriteString("Going MAD here!")
	fmt.Println(byts.String())

	//WRITING to either file or buffer
	var b1 bytes.Buffer
	p1.writeOut(f)
	p1.writeOut(&b1) //address has to be passed here

	//anonymous funcs
	func() {
		fmt.Println("I am anonymous function")
	}()

	func(s string) {
		fmt.Println("My name is : ", s)
	}("Praveen")

	//assigning funcs to variables; in this case we are assigning anonymous funcs as an example

	f1 := func() {
		fmt.Println("I am anonymous function 2")
	}

	f2 := func(s string) {
		fmt.Println("My name is : ", s)
	}
	f1()
	f2("Sangeetha")

	//calling function foo2 (returns an Int) and bar2 (returns a function)
	one := foo2()
	fmt.Println(one)

	two := bar2()
	fmt.Println(two)   //see the output; this prints the address of the function since bar2 returns a function
	fmt.Println(two()) //two() will call the function which is returned by bar2 and that function returns 2

	fmt.Printf("%T\n", foo2)
	fmt.Printf("%T\n", bar2)
	fmt.Printf("%T\n", two)

	//passing a func as argument
	f3 := doMath(12, 23, add)
	fmt.Println(f3)

	f4 := doMath(45, 23, sub)
	fmt.Println(f4)
}

// no parameter, no return
func foo() {
	fmt.Println("I am from foo")
}

// 1 parameter, no return
func bar(s string) {
	fmt.Println("My name is", s)
}

// 1 param, 1 return
func aloha(s string) string {
	return fmt.Sprint("They call me Mr. ", s)
}

// 2 params, 2 returns
func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}

// variatic parameter....it has to be the last parameter; it can be passed no values as well
func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)
	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

func foo1() {
	fmt.Println("foo")
}

func bar1() {
	fmt.Println("bar")
}

// METHODS: Attaching a type to a function
type person struct {
	first string
}

func (p person) speak() { //attaching person type to speak function
	fmt.Println("I am", p.first)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() { //attaching secretAgent type to speak function
	fmt.Println("I am a secret Agent", sa.first)
}

//INTERFACE example

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

// STRINGER interface; just prints the type differently
type book struct {
	title string
}

func (b book) String() string { //attaching book type to String function
	return fmt.Sprint("The title of the book is:", b.title)
}

type count int //count is another type which is of type int

func (c count) String() string { //attaching count type to String function
	return fmt.Sprint("The count is:", strconv.Itoa(int(c))) //coz c is of type count, we need to first convert to int so that strconv.Itoa can be called on that
}

//WRAPPER function

func logInfo(s fmt.Stringer) {
	log.Println("LOGGING:", s.String())
}

// method for person type
func (p person) writeOut(w io.Writer) error { //person type is at line 144
	_, err := w.Write([]byte(p.first))
	return err
}

// returing a function
func foo2() int {
	return 1
}

func bar2() func() int {
	return func() int {
		return 2
	}
}

// functions to be passed as arguments
func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
