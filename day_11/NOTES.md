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
