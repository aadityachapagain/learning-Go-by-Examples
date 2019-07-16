package main

import (
	"fmt"
)

// Person is  ...
type Person struct {
	name   string
	age    int
	gender string
}

func increment(i *int) {
	*i++
}

// PrintHelloworld is ...
func PrintHelloworld() {
	fmt.Println("Bonjour le monde")
}

func add(a int, b int) int {
	c := a + b
	return c
}

func main() {
	i := 2
	// increment the value of i using defined function
	increment(&i)
	// check the value
	fmt.Println(i)

	PrintHelloworld()

	fmt.Println(add(2, 5))

	bob := Person{name: "Sponge Bob", age: 54, gender: "Male"}
	// bob = Person("Sponge Bob", 24, "Male")

	fmt.Println(bob.name, bob.age, bob.gender)
}
