//append function to append items to array returns new array when array is full
package main

import (
  "fmt"
)

func main(){
  var x, y []int
  for i := 0; i <10; i++ {
    y = append(y, i)
    x = append(x, i)
    fmt.Printf("%d cap=%d\t%v\n", i, cap(y),y)
  }
}
