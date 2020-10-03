- interfaces as Contracts: Interface is abstract type. It dosen't expose the representation or internal structure of its values, or the set of basic operations they support; it revels only some of their basic methods. When you have a value of interface type, you know nothing about what it is, you know only what it can do or what behaviours are provided by methods.

we have been using two similar functions for strings formatting: fmt.Printf, which writes the result to standard output ( a file ) and fmt.Sprintf which return the results as string. I would be unfortunate if the hard part, formatting the result, had to be duplicated because of these superficial differences on how the results is used. Both of these function is wrapper around the third function fmt.Fprintf that is agnostic about what happens to the results it computes:

```go
package fmt

func Fprintf( w io.Writer, format string, args ...interface{}) (int, error)

func Printf(  format string, args ...interface{}) (int, error){
  return Fprintf(os.stdout, format, args...)
}

  func Sprintf(format, args ...interface{}) string {
  var buf bytes.buffer
  Fprintf(&buf, format, args...)
  return buf.String()
}
```

- Interfaces Types : Interfaces types specifies a set of methods that a concrete type must possess to be considered an instance of that interface.

  The `io.Writer` type is one of the most widely used interface because it provides an abstraction of all the types where bytes can be written which includes files, memory buffers, network connections, HTTP clients, archivers, hashers and so on.The io package defines many other useful interfaces. A Reader represents any type from which you can read bytes, and Closer is any value that you can close, such as file and network connections.

```go
package io

type Reader interface {
  func Read (p []bytes ) (n int, err error)
}

type Closer interface {
  func Close () error
}
```

- Interface satisfaction : A type satisfies an interface if it possesses all the methods interface requires. For example `*os.File` Satisfies io.Reader, Writer, Closer, and ReadWriter. A \*bytes.Buffer satisfies Reader, Writer and ReadWriter, but does not satisfies closer.

  The assignability rule for the interfaces is very simple: an expression may be assigned to an interface only if its types satisfies the interface. So:

```go
var w io.Writer
w = os.Stdout              // OK : *os.File has write method
w = new(bytes.Buffer)      // OK : *bytes.Buffer has write method
w = time.Second            // Compile Error: time.Duration lacks write method

var rwc io.ReadWriterCloser
rwc = os.Stdout            // OK : OK: *os.File has Read, Write, Close methods
rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
```
