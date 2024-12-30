package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("INIT function")
}

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1) //saying the runtime to wait for the next item which will run in separte Go routine
	go foo()  //runs or launches another go routine;hence you will not see foo printing in the output unless you use sync.WaitGroup
	bar()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait() //Waits for any launched Go routines to finish
}

func foo() {
	for i := 0; i < 10; i++ {

		fmt.Println("foo:", i)
	}
	wg.Done() //Indicating that the task for which something is waited on is completed
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
