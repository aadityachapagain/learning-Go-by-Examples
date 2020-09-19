// dup1 search for duplicate lines from input stream

package main

import (
  "fmt"
  "os"
  "bufio"
)

func main(){

  counts := make(map[string]int)
  input := bufio.NewScanner(os.Stdin)
  for input.Scan() {
    counts[input.Text()]++
  }
  // Note: ignoring potentials erors from input.Err()
  for line, n:=  range counts {
    if n > 1 {
      fmt.Println("%d\t%s\n", n, line)
    }
  }
}

