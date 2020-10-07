- Type Assertions: Type assertions is an operation applied to the interface value. Syntactically, it looks like x.(T), where x is an expression of an interface type and T is a type, called the "asserted" type. A type assertion checks that the dynamic type of its operand matches the asserted type.

There are two possibilities. First, if the asserted type T is a concrete type, then type assertions checks whether x's dynamic type to T. if this check succeeds, the result of the type assertions is x's dynamic value, whose type is of course T. In other words, a type assertion to a concrete type extracts the concrete value from its operand. If the check fails, then the operations panics. For example

```go
var w io.Writer

v = os.Stdout
f := w.(*os.File)     // success  : f  == os.Stdout
g := w.(*bytes.Buffer) // panic interface holds *os.File, not *bytes.Buffer
```

Second, if instead the asserted type T is an interface type, then the type assertion checks whether x's dynamic type satisfies T. If this check succeeds, the dynamic value is not extracted; the result is still interface value with same type and value component, but the result has the interface type T. In other words, a type assertion to an interface type changes the type of the expression, making a different (and usually larger ) set of methods accessible, but it preserves the dynamic type and value components inside the interface value. For example

After the first type assertion below, both w and rw hold os.Stdout so each has a dynamic type of \*Os.File, but w, an io.Writer, exposes only the files's Write method, whereas rw exposes its Read method too.

```go
var w io.Writer

w = os.Stdout
rw := w.(io.ReadWriter)   // success: *os.File has both Read and Write

w = new(ByteCounter)
rw = w.(io.ReadWriter)   // panic: *ByteCounter has no Read method

```

Often we're not sure of the dynamic type of an interface value, and we'd like to test weather it is some particular type. If the type assertion appears in an assignment in which two results are expected, such as the following declarations, the operation does not panic on failure but instead returns an additional second result, a boolean indicating success:

```go

var w io.Writer = os.Stdout
f, ok := w.(*os.File)              // success: ok, f == os.stdout
b, ok := w.(*bytes.Buffer)         // failure: !ok, b == nil
```

The second result is conventionally assigned to a variable named ok. If the operation failed, ok is false, and the first result is equal to the zero value of the asserted type, which in this example is a nil \*bytes.Buffer.

- Querying Behaviors with Interface Type Assertions: Lets consider a example of web server, where web app is responsible for writing HTTP header, fields such as "Content-type: text/html". The io.Writer w represents the HTTP response; the bytes written to it are ultimately sent to someone's web browser.

```go
func writeHeader(w io.Writer, contentType string) error {
  if _, error := w.Write([]byte("Content-Type: ")); err != nil{
    return err
  }
  if _, error := w.Write([]byte(contentType)); err != nil {
    return err
  }
  // ...
}
```

Because the Write method requires a byte slice, and the value we wish to write is a String, a []byte(...) conversion is required. This conversion allocates memory and makes a copy, but the copy is thrown away almost immediately.So, lets pretend that this is core part of the web server and that our profiling has revealed that this memory allocation is slowing it down.Can we avoid allocating memory here ?

The `io.Writer` interface tells us only one fact about the concrete type that w holds: that bytes may be written to it. If we look behind the curtains of the net/http package, we see that the dynamic type that w holds in this program also has a WriteString method that allows strings to be efficiently written to it, avoiding the need to allocate the temporary copy.

We cannot assume that an arbitrary io.Writer w has the WriteString method. But we can define a new interface that has just this method and use a type assertion to test whether the dynamic type of w satisfies this new interface.

```go

// writeString writes s to w
// if w has a WriteString method, it is invoked instead of w.Write.

func writeString(w io.Writer, s string) (n int, err error) {
  type stringWriter interface {
    WriteString(string) (n int, err error)
  }
  if sw, ok := w.(stringWriter); ok {
    return sw.WriteString(s)  //avoid a copy
  }
  return w.Write([]bytes(s))  //allocate a temporary memory
}

func writeHeader(w io.Writer, contentType string) error {
  if _, error := writeString([]byte("Content-Type: ")); err != nil{
    return err
  }
  if _, error := writeString([]byte(contentType)); err != nil {
    return err
  }
  // ...
}

```

what most exciting thing about above example is that there is no standard interface that defines the `WriteString` method and specifies it required behaviour, Furthermore, whether or not a concrete type satisfies the `stringWriter` interface is determined only by its method, not by any declared relationship between it and the interface type. What this means is that the technique above relies on the assumption that if a type satisfies the interface below, then `WriteString(s)` must have the same effect as `Write([]bytes(s))`.

- Type Switches: Interfaces are used in two distinct styles. First style, exemplified by `io.Reader`, `io.Writer`, `fmt.Stringer`, `sort.Interface`, `http.Handler`, and `error` an interface's methods express the similarities of the concrete types that satisfy the interface but hide the representation details and intrinsic operations of those concrete types. The emphasis is on the methods, not on the concrete types.

The second style exploits the ability of an interface value to hold values of a variety of concrete types and considers the interface to be the union of those types. Type assertions are used to discriminate among these types dynamically and treat each case differently. In this style, the emphasis is on concrete types that satisfy the interface, not on the interface's methods, and there is no hiding of information.

Consider example of Go's API for querying an SQL database , like those of other language.

```go
import "database/sql"
func listTracks(db sql.DB, artist string, minYear, maxYear int) {
  result , err := db.Exec(
    "SELECT * FROM tracks WHERE artist  = ? AND ? <= year AND <= ?",artist, minYear, maxYear
  )
  // ...
}
```

The Exec method replaces each '?' in the query string with an SQL Literal denoting the corresponding argument value, which may be boolean, a number, string or nil. Constructing queries this way help avoid SQL injection attacks, in which an adversary takes control of the query by exploiting improper quotation of input data. Within Exec, we might find a function like the one below, which converts each argument value to its literal SQL notation.

```go
func sqlQuote(x interface{}) string {
  if x == nil {
    return "NULL"
  } else if _, ok := x.(int); ok {
    return fmt.Sprintf("%d",x)
  } else if _, ok := x.(uint); ok {
    return fmt.Sprintf("%d", x)
  } else if _, ok := x.(bool); ok {
    if b {
      return "TRUE"
      }
    return "FALSE"
  } else if s, ok := x.(string); ok{
    return sqlQuoteString(s)
  } else {
  panic(fmt.Sprintf("unexpected  type %T: %v", x,x))
  }
}
```

Above example looks kind of messy with lots of if else chain which can be revamped into simplistic form using switch statements. A type switch looks like an ordinary statements in which the operand is x.(type) - that's literally the keyword type- and each case has one or more types. A type switch enables a multi-way branch based on the interface value's dynamic type. The nil case matches if x==nil, and the default case matches if no other case does. A type switch for sqlQuote would have these cases:

```go

switch x.(type) {
  case nil:     //...
  case int, uint:   //...
  case bool:       // ...
  case string:     // ...
  default:         // ...
}

```

Above in first example of type switches, we can see that string and bool require value extracted by type assertion. Since this is typical, the type switch statement has an extended form that binds the extracted value to a new variable within each case:

```go
switch x := x.(type) { /* ...*/ }
```

Here we've created new variable x which is same as input variable x;as with type assertions, reuse of variable names is common. Like a switch statement, a type switch implicitly creates a lexical block, so the declaration of the new variable called x does not conflict with a variable x in an outer block. Each case also implicitly creates a separate lexical block.

Rewriting original sqlQuote function to use extended form of type switch makes code significantly clearer and manageable:

```go
func sqlQuote(x interface{}) string {
  switch x := x.(type) {
    case nil:
      return "NULL"
    case int, unit:
      return fmt.Sprintf("%d",x)
    case bool:
      if x {
        return "TRUE"
      }
      return "FALSE"
    case string:
      return sqlQuoteString(x)
    default:
      panic(fmt.Sprintf("unexptected type %T: %v", x,x))
  }
}
```
