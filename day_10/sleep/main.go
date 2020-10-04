package main

import (
  "fmt"
  "flag"
  "time"
)

var period = flag.Duration("period", 1*time.Second, "sleep Duration")

func main(){
  flag.Parse()
  fmt.Printf("Sleeping for %v ...", *period)
  time.Sleep(*period)
  fmt.Println()
}
