- Method Declarations: method is declared by adding extra parameter before function name. The parameter attach the function to the type of that parameter. eg

```go
//ordinary function declaration

func func_name(name type ) return_type {
  //function body
  ...
}

// method declaration

func (p type) func_name (name type) return_type {
  // function body
  ...
}
```

The extra parameter p is called the method's receiver, a legacy from early object-oriented languages that described calling a method as "sending a message to an object".

- Method with pointer Receiver: Because calling a function makes a copy of each arguments value, if a function needs to update a variable, or if an argument is so large that we wish
  to avoid copying it, we must pass the address of the variable using pointer. The same goes for methods that need to update the receiver variable eg:

```go
func (p *Point) ScaleBy( factor float64 ) {
  p.X *= factor
  p.Y *= factor
}

// the name of this method is (*Point).ScaleBy
// The above method can be called by providing a *Pointer receiver like this

r := &Point{1, 2}

r.ScaleBy(2)

fmt.Println(*r)

//or

p := Point{1,2}
pptr := &p

pptr.ScaleBy(2)

fmt.Println(p)

//or this

p := Point{1,2}
(&p).ScaleBy(1,2)
fmt.Println(p)
```
