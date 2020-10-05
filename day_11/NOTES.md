- Type Assertions: Type assertions is an operation applied to the interface value. Syntactically, it looks like x.(T), where x is an expression of an interface type and T is a type, called the "asserted" type. A type  assertion checks that the dynamic type of its operand matches the asserted type.

There are two possibilities. First, if the asserted type T is a concrete type, then type assertions checks whether x's dynamic type to T. if this check succeeds, the result of the type assertions is x's dynamic value, whose type is of course T. In other words, a type assertion to a concrete type extracts the concrete value from its operand. If the check fails, then the operations panics. For example

```go
var w io.Writer

v = os.Stdout
f := w.(*os.File)     // success  : f  == os.Stdout
g := w.(*bytes.Buffer) // panic interface holds *os.File, not *bytes.Buffer
```
Second, if instead the asserted type T is an interface type, then the type assertion checks whether x's dynamic type satisfies T. If this check succeeds, the dynamic value is not extracted; the result is still interface value with same type and value component, but the result has the interface type T. In other words, a type assertion to an interface type changes the type of the expression, making a different (and usually larger ) set of methods accessible, but it preserves the dynamic type and value components inside the interface value. For example

After the first type assertion below, both w and rw hold os.Stdout so each has a dynamic type of *Os.File, but w, an io.Writer, exposes only the files's Write method, whereas rw exposes its Read method too.

```go
var w io.Writer

w = os.Stdout
rw := w.(io.ReadWriter)   // success: *os.File has both Read and Write

w = new(ByteCounter)
rw = w.(io.ReadWriter)   // panic: *ByteCounter has no Read method

```


