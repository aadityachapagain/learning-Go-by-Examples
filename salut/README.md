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
  
lets see by running this file. There are two ways we can run a Go command. As we know, Go is a compiled language, so we first need to compile it before executing.

```bash
go run main.go
```

This creates a binary executable file `main` which now we can run:

```bash
$ /.main

salut beaute !
```

There is another, simpler, way to run the program. The go run command helps abstract the compilation step. You can simply run the following command to execute the program.

```bash
$ go run main.go

salut beaute !
```

## Variables

---------------

Variables in Go are declared explicitly. Go is a statically typed language. This means that the variable type is checked at the time of variable declaration. A variable can be declared as:

```go
var a int
```

In this case, the value will be set as 0. Use the following syntax to declare and initialize a variable with a different value:

```go
var a = 1
```

Here the variable is automatically assigned as an int. We can use a shorthand definition for the variable declaration as:

```go
message := "quoi de neuf"
```

We can also declare multiple variables in the same line:

```go
var b, c int = 2, 3
```

Amazing !  Now lets learn about datatypes in go lang

## Data types

### Number, String and Boolean

---------------

The string type stores a sequence of bytes. It is represented and decleared with keyword `string`.

A boolean value is stored using the keyword bool.  

Go also supports complex number type data tyeps, which can be declared with `complex64` and `complex128`.

```go
var a bool = true
var b int = 1
var c string = 'salut que de neuf'
var d float32 = 1.2222
var x complex128 = complex.Sqrt(-5 + 12i)

```

### Arrays, Slices and Maps

---------------

An array is a sequence of elements of the same data type. Arrays have a fixed length dfined at declaration, so it cannot be expanded more than that. An array is declared as:

```go
var a [5] int
```

Arrays can also multidimensional.We can simply create them with the following format:

```go
var b [5][1] int
```

Arrays are limiting for cases when the values of array changes in runtime.Arrays also do not provide the ability to get a subarray.For this,Go has a data type called slices.

Slices store a sequence of elements and can be expanded at any time.Slice declaration is similar to the array declaration - without the capacity defined.

```go
var b [] int
```

This creates a slice with zero capacity and zero length. Slice can also be defined with capacity and length. we can use following syntax for it:

```go
numbers := make([] int , 5, 10)
```

Here, the slice has an initial length of 5 and has a capacity of 10.

Slice are an abstraction to an array. Slice use an array as an underlying structure. A slice contains  three components: capacity, length, and a pointer to the underlying array as shown in the diagram below:  
  

![Slice as a array](https://cdn-images-1.medium.com/max/800/1*P0lNCO0sQwIYHLEX_mfSOQ.png)

The capacity of slice can be increased by using the append or a copy function. An append function adds value to the end of the array and also increses the capacity if needed.

```go
numbers =  append(numbers, 1, 2, 3, 4)
```

Another way to increase the capacity of a slice is to use the copy function.Simply create another slice with a larger capacity and copy the original slice to the newly created slice:

```go
// create a new slice
number2 := make([] int, 15)
// copy the original slice to new slice
copy(number2, number)
```
