package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Bonheur")

	resp, err := http.Get("http://goisawsome.com")

	if err != nil {
		print(err)
		return
	}
	fmt.Println(resp)
}
