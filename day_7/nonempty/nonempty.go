// Nonempty is an example of an in-place slice algorithms
package main

import "fmt"

func nonempty(strings []string) []string {
  i := 0
  for _, s := range strings {
    if s!= ""{
      strings[i] = s
      i++
    }
  }
  return strings[:i]
}
