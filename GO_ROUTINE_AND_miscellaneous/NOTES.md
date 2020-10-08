- Goroutines: In go, each concurrently executing block of scope is called `goroutine`. When you execute the program by doing `go run ` then its only goroutine is the one that called `main` function so we call it `main goroutine` while, New go routines can be created by `go` statement. A go statement causes the function to be called in a newly created goroutine. The go steatement itself complete immediately:

```go
f()                // calls f wait for it to return
go f()             // calls f concurrently and immediately goes to next statemtent, while content of f is still executing.
```

- channels: If goroutines are the activities of a concurrent Go program, channels are the connections between them. A channel is a communication mechanism that lets one goroutine send values to another go routine. Each channel is a conduit for values of particular type int is written chan int.

  To create a channel we use default builtin make function:

```go
ch := make( chan int )
```

As shown above, a channel is the reference to the data structure created by make. When we copy a channel or pass one as an argument to a function, we are copying a reference, so caller and callee refer to the same data structure. As with other reference types, the zero value of a channel is nil.

Two channel of same type may be compared using `==`. The comparison is true if both are references to same data channel structure. A channel may also be compared to nil.

A channel has two principle operations, send and receive, collectively known as communications. A send statement transmit a value from one goroutine, through the channel, to another goroutine executing a corresponding receive expression. Both operations are written using the `<-` operator.

```go
ch <- x   // a send statement
x =  <-ch //a receive expression in an assignment statement
<-ch      // a receive statement; result is discarded
```

Channel also support the third operation, close, which sets a flag indicating that no more values will ever be send to this channel;subsequent attempts to send will panic.

To close the channel we call builtin function `close`:

```go
close(ch)
```

A channel created with a simple call to make is called an unbuffered channel, but make accepts an optional second argument, an integer called the channel's capacity.If the capacity is non zero, creates a buffered channel.

```go
ch := make(chan int)          // unbuffered channel
ch := make(chan int, 0)       // unbuffered channel
ch := make(chan int, 3)       // buffered channel with capacity 3
```
