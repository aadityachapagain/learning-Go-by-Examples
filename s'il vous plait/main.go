package main

import (
	"fmt"
)

func main() {
	if num := 9; num < 0 {
		fmt.Println(num, "is negative !")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}

	i := 0

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("None")
	}

	sum := 0

	for i < 2000 {
		sum += i
		i++
	}
	fmt.Println(sum, " is total sum")

	sum = 0

	for i := 0; i < 2000; i++ {
		sum += i
	}
	fmt.Println(sum, "is total sum.")
}
