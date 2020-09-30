//function defined in the way shown below will have access to the entire lexical environment,
//so inner function can refer to variables from enclosing function

package main

import (
  "fmt"
)

func squares() func() int {
  var x int
  return func() int {
    x++
    return x*x
  }
}

func main() {
  f := squares()
  fmt.Println(f())
  fmt.Println(f())
  fmt.Println(f())
  fmt.Println(f())
}
