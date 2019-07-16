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

// method defination
func (p *Person) describe() {
	fmt.Printf("%v is %v years old.", p.name, p.age)
}
func (p *Person) setAge(age int) {
	p.age = age
}

func (p Person) setName(name string) {
	p.name = name
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

	pp := &Person{name: "Bob the swimmer", age: 42, gender: "Male"}
	pp.describe()
	// => Bob is 42 years old
	pp.setAge(45)
	fmt.Println(pp.age)
	//=> 45
	pp.setName("Hari")
	fmt.Println(pp.name)
	//=> Bob the swimmer
}
