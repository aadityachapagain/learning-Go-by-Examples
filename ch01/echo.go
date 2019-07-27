package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for id, arg := range os.Args[1:] {
		s += string(id) + " " + sep + arg
		sep = " "
	}
	fmt.Println(s)
}
