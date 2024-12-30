package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("INIT function")
}

func main() {

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	counter := 0 //global counter used by various go routines spawned below

	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex //having a mutex
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() //locking to avoid race condition
			v := counter
			//time.Sleep(time.Second * 2)  //uncomment and remove mutex and you will see counter stays at 1 coz every routine starts, reads old counter and then yields
			runtime.Gosched() //yielding for other routines/threads to run
			v++
			counter = v
			mu.Unlock() //unlocking
			wg.Done()
		}()
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("counter", counter)
}
