package main

import "fmt"

func square() func() int {
	var x int

	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := square()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

// use of Anonymous function to  persist the variable values across many call to anynymous function
