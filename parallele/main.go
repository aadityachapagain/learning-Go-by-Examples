package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	time.Sleep(100 * time.Millisecond)
}

func hello() {
	fmt.Println("hello from go routine ... ")
}
