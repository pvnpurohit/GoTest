package main

import (
	"fmt"
	"io"
	"os"
)

func init() {
	fmt.Println("INIT function")
}

func main() {
	fmt.Println("Writer interface")
	fmt.Fprintln(os.Stdout, "Writer interface")   //since fmt.Println calls fmt.Fprintln, we can print using fmt.Fprintln
	io.WriteString(os.Stdout, "Writer interface") //another way of writing/printing the same msg
}
