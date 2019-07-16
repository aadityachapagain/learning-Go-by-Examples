# Go Basics

This part of the go programming tutorial series covers Pointers, Refrences, struct type, Methods and Interfaces  as a object Oriented approches in Go though Go dosen't directly support object Oriented approch.

## Pointers

Go provide Pointers. Pointers are the  are the place to hold the address of the value. A pointer is defined by *. A pointer is defined according to the type of data. Example:

```go
var val *int
```

Here *val* is a pointer to a integer type. The *&* operator van be used to get a address of a variable.

```go
value := 12
// Now & operator can be used to access the address of variable 
val = &value
```

The value pointed by the pointer can be accessed using the * operator:

```go
fmt.Println(*val)
```

Pointers are usually perfered while passing a struct as an argument or while declearing a method for a defined type.

1. While passing value the value if actually copied which means more memory it occupies
2. with the pointer passed to function argument, the value changed by the function is refled back in the method/function caller.

Example:

```go

func increment(a *int){
    *a ++
}

func main(){
    i := 2
    increment(&i)
    fmt.Println(i)
}

```

## Functions

The *main* function defined in the main package is the entry point for a go program to  execute. More functions can be defined and can be used.
Lets look into simple example:

```go
func PrintHelloworld(){
    fmt.Println("Bonjour le monde")
}

func add( a int, b int) int {
    c := a + b
    return c
}

func main(){
    PrintHelloworld()

    fmt.Println(add(2, 5))
}
```

So, Defining function is easy in GO. and if you already know C than its just more than  a easy.

```go

// Function Defination go
func <function_name>(<args> <args_type>, ...) (<return_type>, ...) {
"
    Function Body
"
}

// <args_type> can be simple datatypes or may be Pointers of the datatype like we defined in the previous example
```

The return of a function can be predefined in function as well:

```go

func add(a int, b int) (c int) {
  c = a + b
  return
}
func main() {
  fmt.Println(add(2, 1))
}
//=> 3
```

Here c is defined as the return variable. So the variable c defined would be automatically returned without needing to be defined at the return statement at the end.

You can also return multiple return values from a single function separating return values with a comma.

```go

func add(a int, b int) (int, string) {
    c := a +b
    return c, "successfully added"
}

func main(){
    sum, message = add(3,7)
    fmt.Println(sum, message)
}

```

## Methods, Struct and Interfaces

As I already acknowledged, Go is not completely object-oriented language, but with structs, interfaces, and methods its has a lot  of object-oriented support and feel.

### Struct

-------------------


A struct is typed, collection of different fields. A struct is used to group data together. For example, if we want to group data of Person type, we define a person attributes  and attitude but attitude is subjective things we left off with attributes only, which could include name, age, gander, A struct can be defined using following syntax:

```go
type Person struct {
    name string
    age int
    gender string
}
```

With a Person Struct type now lets create a Person:

```go

// way 1: specifying the attributes and values
p := person{name: "SpongeBob", age: 42, gender: "M"}

p.name

// specifying only name
Person{"SpongeBob square pant" 23, "M"}
```

We can easily access these data with `.` operator.

```go

fmt.Println(p.name, p.age, p.gender)
```

You can also access attributes of a struct directly with its Pointer.

```go

pp := &person{name: "Bob", age: 42, gender: "Male"}
pp.name
//=> Bob

```

### Methods

-------------------

Methods are a special type of function with a receiver. A receiver can be both a value or a pointer. Let’s create a method called describe which has a receiver type person we created in the above example:

```go

// struct defination
type Person struct {
  name   string
  age    int
  gender string
}

// method defination
func (p *Person) describe() {
  fmt.Printf("%v is %v years old.", p.name, p.age)
}
func (p *Person) setAge(age int) {
  p.age = age
}

func (p Person) setName(name string) {
  p.name = name
}

func main() {
  pp := &Person{name: "Bob the swimmer", age: 42, gender: "Male"}
  pp.describe()
  // => Bob is 42 years old
  pp.setAge(45)
  fmt.Println(pp.age)
  //=> 45
  pp.setName("Hari")
  fmt.Println(pp.name)
  //=> Bob the swimmer
}

```

As we can see in the above example, the method now can be called using a dot operator as `pp.describe`. Note that the receiver is a pointer. With the pointer we are passing a reference to the value, so if we make any changes in the method it will be reflected in the receiver pp. It also does not create a new copy of the object, which saves memory.

Note that in the above example the value of age is changed, whereas the value of name is not changed because the method setName is of the receiver type whereas setAge is of type pointer.

### Interfaces

-------------------

Go interfaces are collection of methods. Interfaces help group together the properties of a type.Lets take an example of interface animal:

```go
type animal interface {
  description() string
}

```

Here animal is an interface type. Now let’s create 2 different type of animals which implement the animal interface type:

```go

package main

import (
  "fmt"
)

type animal interface {
  description() string
}

type cat struct {
  Type  string
  Sound string
}

type snake struct {
  Type      string
  Poisonous bool
}

func (s snake) description() string {
  return fmt.Sprintf("Poisonous: %v", s.Poisonous)
}

func (c cat) description() string {
  return fmt.Sprintf("Sound: %v", c.Sound)
}

func main() {
  var a animal
  a = snake{Poisonous: true}
  fmt.Println(a.description())
  a = cat{Sound: "Meow!!!"}
  fmt.Println(a.description())
}

//=> Poisonous: true
//=> Sound: Meow!!!

```

In the main function, we create a variable a of type animal. We assign a snake and a cat type to the animal and use Println to print a.description. Since we have implemented the method describe in both of the types (cat and snake) in a different way we get the description of the animal printed.
