# GO BASICS

This part of the lesson will give the fundamental overview of Go programming Language.

---------------

## Getting Started

Go is made of packages. The package main tells the Go compiler that the program is compiled as an executable, rather than a shared library. It is the entry point for an application. The main package is defined as:
  
```go
package main
```
  
Let’s move ahead by writing a simple hello world example by creating a file main.go in the Go workspace.

## Workspace

---------------

A workspace in Go is defined by the environment variable `GOPATH`.

Any code you write is to be written inside the workspace. Go will search for any packages inside the GOPATH directory, or the GOROOT directory, which is set by default when installing Go. GOROOT is the path where the go is installed. [ which I think, I cover in the very first lesson]

Create the file main.go with the following code inside the workspace folder we just created.
  
## Salut beauté

```go
package main

import (
 "fmt"
)

func main(){
  fmt.Println("Salut beauté !")
}
```
  
In the above example, fmt is a built-in package in Go which implements functions for formatting I/O.
  
We import a package in Go by using the import keyword. func main is the main entry point where the code gets executed. Println is a function inside the package fmt which prints `Salut beauté !` for us.
