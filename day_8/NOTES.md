- Functions Declarations:

```go
func name(parameter-list) (result-list) {
  body
}
```

- Return Multiple Values:

```go
func name(links []string)(Node*, int)
  body
  return node(links), value
```

- Errors: Some function always succeed at their task while some other might only succeed only if met certain precondition or might panic while calling the function. Panic is the pure sign of bug in the code and should never happen in well written program.
  While for many other function, even in a well-written program, success is not assured as it depends on factors beyond programmer's control.
  Errors are thus an important part of a packages's API or an application's user interfaces, and failure is just one of several expected behaviors.This is the approch Go takes to error handling.

- Functions Values: Functions are first class values in Go, like other values function values have type and they may be assigned to variables or passed to or return from function

```go
func square(x int) int { return x * x }

f := square

fmt.PrintF("squre of %d is : %d", 3, f(3))
```

Function values can be initialized from function type and may be compared with nil

```go
var f func(int) int
if f != nil{
  f(3)
}
```

- Anonymous Function: A function literal can be written without name which denote anonymous function

```go

strings.Map(func(r rune) rune {return r +1}, "adminx")
```

- Variadic Function: that can be called with varying number of aurguments. The most familiar examples are fmt.Printf and its variants.
  `fmt.Printf` required one fixed arguments at the beginning then accepts any number of subsequent arguments.
  declare variadic function by preceding type of final paramter by an ellipses, `...`. example

```go
func sum(vals ...int) int {
  total := 0
  for _, val := range vals{
    total += val
  }
  return total
}
```

How to invoke variadic function when the arguments are already in slice just place ellipses after final arguments.

```go
values := []int{1,2,3,4,5,6}
fmt.Println(sum(values...))
```

- Deferred Function calls: Syntactically, a defer statement is an ordinary function or method call prefixed by the keyword defer. The function and argument expressions are evaluated when the statements are executed.
  but actual call is deferred until the function that contains the defer statement has finished, whether normally by executing return statement or falling of the end or abnormally by panicking. Any number of calls may be deferred they are executed in the reverse of the order in which they are deferred.

```go
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
  mu.Lock()
  defer mu.Unlock()
  return m[key]
}
```

- Panic: Go’s type system catch es many mistakes at compile time, but others, like an out-of-bounds array access or nil pointer dereference, require checks at run time. When the Go runtime detects these mistakes, it panics.

- Recover: Giving up is usually the right response in a panic, but not always. It might be possible to recover in some ways, or at least cleanup the mess before quitting.
