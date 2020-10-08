package main

import (
  "fmt"
  "time"
  "os"
  "strconv"
)

func main(){
  n, err := strconv.Atoi(os.Args[1])
  if err != nil {
    fmt.Println(err)
    os.Exit(3)
  }
  go spinner(100 * time.Millisecond)
  fibN := fib(n)
  fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func fib(x int) int {
  if x < 2 {
    return x
  }
  return fib(x -1) + fib(x-2)
}

func spinner(delay time.Duration) {
  for {
    for _,r := range `-\|/` {
      fmt.Printf("\r%c",r)
      time.Sleep(delay)
    }
  }
}
