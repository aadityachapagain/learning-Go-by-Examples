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

- Unbuffered Channel: A send operation on an unbuffered channel blocks the sending goroutine until another goroutine executes a corresponding receive on the same channel, at which point the value is transmitted and both goroutines may continue. Conversly, if receive operation is executed first then the receiving goroutines is blocked until another goroutines performs a send on the same channel.

  Communication over the unbuffered channel causes the sending and receiving channel to `synchronize`.Because of this unbuffered channels often called `synchronous channel`

- Pipelines: channels can be used to to connect goroutines together so that output of one is the input to another. This is called `pipline`. The program below consists of three goroutine connected by two channels.

  the first goroutine, counter, generate the integers 0,1,2,... and sends them over a channel to the second goroutine, squarer, which receives each value, squares it, and sends the result over another channel to the third goroutine, printer, which receives the squared values and prints them. for clarity of this example, we have intentionally, chosen very simple functions, though of course they are too computationally trivial to warrant their own goroutines in a realistic program.

```go
func main(){
naturals := make(chan int)
squares := make(chan int)

  //counter
  go func(){
    for x := 0; ; x++ {
      naturals <- x
    }
  }()

  //Squarer
  go func(){
    for {
      x := <-naturals
      squares <- x * x
    }
  }

  //printer ( in main goroutine )
  for {
      fmt.Println(<- squares)
  }
}

```

The above program prints infinite series of squares, Pipelines like this may be found in long running servers where channels are used for lifelong communications between goroutines containing infinite loops. But what if we want to run finite number of loops only through pipelines ?

if sender knows that no further values will ever be sent on a channel, it is useful to communicate this fact to the receiver goroutines so that they stop waiting. This can be accomplished by closing the connections using built in `close` function.

```go
close(naturals)
```

After channels has been close, any further send operation in that channels will panic. After the closed channel has been drained, that is, after the last send element has received, all subsequent receive operations will proceed without blocking but will yield a zero value.Closing the naturals channel above would cause the squarer's loop to spin as it receives a never ending stream of zero values, and to send zero values to the printer.

there is no way to test directly whether a channel has been closed, but there is variant of receive operation that produces two results: the received channel element, plus a boolean value which is true for successfully receive operation and false for receive on a close and drained channel.Using this, features we can modify the squarer's loop to stop when the naturals channel is drained and close the squares channel in turn.

```go
// using boolea values
x, ok := <- naturals
if  !ok {
  break
}

//using for range loop
for x := range  naturals {
  squares <- x * x
}

for x: range(squares){
  fmt.Println(x)
}
```

- Unidirectional Channel Types: It is not in the nature of channel to use it as a unidirectional way rather the our intention to use channels as unidirectional purpose by supplying channel as a function parameter.
  To document this intent of using channels and prevent misuse, Go type system provides unidirectional channel types that expose only one or the other of the send and receive operations. The type `chan <- int`, a send-only channel of int, allows send but not receives. Conversely, the type `<-chan int`, a receive only channel of int, allows receive but not sends.Violations of this discipline are detected at compile time.

- Buffered Channel: A buffered channel has queue of elements. The queue maximum size is determined during its creation, by the capacity argument to make.Below statements create a buffered channel capable of holding 3 string values.

```go
ch = make(chan string, 3)
```

A send operations in buffered channels inserts an elements at the back of the queue, and receive operation removes an element from the front .If the channel is full, the send operation blocks its send statement execution inside goroutine until space is freed by executing receive operation in that channels. Conversely, if the channel is empty, a receive operation blocks until a value is sent by another goroutine.

We can send up to three values on previously defined goroutine without blocking the goroutine:

```go
ch <- "A"
ch <- "B"
ch <- "C"

// channel is full and blocks the goroutine
ch <- "D"

//in another go routine
fmt.Println(<- ch)  //"A"

//len to return no of buffer of channel occupied by element
fmt.Println(len(ch))  // 2

//cap to determine the capacity of the channel
fmt.Println(cap(ch))  // 3
```

We can easily get the channel capacity and length by corresponding `cap` and `len` builtin function .
