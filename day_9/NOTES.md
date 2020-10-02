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

- Composing Types by struct embedding: Consider the type ColoredPoint:

```go
import "image/color"

type Point struct{X, Y float64}

type ColoredPoint struct {
  Point
  Color color.RGBA
}
```

and previously we see that embedding let us take a syntatic shortcut to defining a ColoredPoint that contains all the fields of Point, plus some more. A similar mechanism applies to the methods of Point. We can call methods of the embedded Point field using a receiver of type ColoredPoint, even though ColoredPoint have no declared methods:

```go
red := color.RGBA{255,0,0,255}
blue := color.RGBA{0,0,255,255}
var p = ColoredPoint{Point{1, 1}, red}
var q = ColoredPoint{Point{5, 4}, blue}

fmt.Println(p.Distance(q.Point)) // "5"

p.ScaleBy(2)
q.ScaleBy(2)

fmt.Println(p.Distance(q.Point)) // "10"
```

The methods of the Point has been promoted to ColordPoint. In this way embedding allows complex types with many methods to be built up by composition of several fields, each providing few methods.
People familiar with object-oriented programming may be tempted to view Point as a base class and ColoredPoint as a subclass or derived class. But that would be mistake. Notice the calls to Distance above.
Distance has parameter of type Point, and q is not a Point, so although q does have an embedded field of that type, we must explictly select it.Attempting to pass q would be an error:
`p.Distance(q) // compile error: cannot use q (ColoredPoint) as Point`
A ColoredPoint is not a Point but it has a Point, and it has two additional methods Distance and ScaleBy promoted from Point.

- Method values and expressions : Usually we select and call method in the same expressions as in p.Distance(), but it's possible to separate these two operations. The selector p.Distance yields a method value, a function that binds a method (Point.Distance) to a specific receiver value p. This function then be invoked without receiver value; it needs only non-receiver arguments.

```go
p := Point{1,2}
q := Point{4, 6}

distanceFromP := p.Distance
fmt.Println(distanceFromP(q))
var origin Point
fmt.Println(distanceFromP(origin))

scaleP := p.ScaleBy

scaleP(2) // p becomes (2, 4)
scaleP(3) //           (6, 12)
scale(10) //           (60, 120)
```

Related to the method value in the method expression. When calling a method, as opposed to an ordinary function, we must supply receiver in a special way using the selector syntax, A method expression, written T.f where T is a type, yields a function, value with a regular first parameter, taking, the place of the receiver, so it can be called in the usual way.

```go
p := Point{1,2}
q := Point{2,3}

distance := Point.Distance   //method expression
fmt.Println(distance(p,q))
fmt.Printf("%T\n", distance)

scale := Point.ScaleBy
scale(&q, 2)
fmt.Println(p)
fmt.Printf("%T\n", scale)
```
