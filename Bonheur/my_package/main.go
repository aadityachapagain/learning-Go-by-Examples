package main

import (
	"fmt"
	"my_package/person"
)

func main() {
	p := person.Description("Milap")
	fmt.Println(p)
}
