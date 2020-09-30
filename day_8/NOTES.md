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

* Functions Values: Functions are first class values in Go, like other values function values have type and they may be assigned to variables or passed to or return from function

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
