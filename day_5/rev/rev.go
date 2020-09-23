// reverse reverses a slice of ints in place

package main

import (
  "fmt"
)

func main(){
  s := [...]int{1,2,3,4,5,6}
  reverse(s[:])
  // reverse a slice
  fmt.Println(s)
  // now rotate a slice by n position
  n := 2
  a := [...]int{1,2,3,4,5,6,7,8,9}
  fmt.Println("before rotating : ", a)
  reverse(a[:n])
  reverse(a[n:])
  reverse(a[:])
  fmt.Println("After rotating ",n," positions :", a)
}

func reverse( s []int ){
  for i,j := 0, len(s) -1; i < j; i,j = i+1, j - 1{
    s[i], s[j] = s[j], s[i]
  }
}


