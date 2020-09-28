package main

import (
  "fmt"
)

func main(){
  x := [...]int{0,1,2,3,4}
  fmt.Println("cap(x) :", cap(x))
  fmt.Println("len(x) :", len(x))
}
