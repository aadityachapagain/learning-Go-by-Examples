# GO Basics

This part of the go programming tutorial series covers parallelism, concurrency and channels.

## Concurrency

Go is built with concurrency in mind. Concurrency in Go can be achieved by Go routines which are lightweight threads.

### Go Routine

------------------------

Go Routine are the functions which can run in parallel or concurrently with another function. Creating a Go routine is farily simple task. Simple by adding a Keyword Go in front of a function, we can make it execute in parallel. Go routines are very lightweight, so we can create a thousand of them. Let's look into a simple example:

```go

package main

import (
    "fmt"
    "time"
)

func main(){
    go c()
    fmt.Println("I am main")
    time.Sleep(time.Second * 2)
}

func c(){
    time.Sleep(time.Second * 2)
    fmt.Println("I am concurrent !")
}

```

As you can see in the above example, the function c is a Go routine which executes in parallel with the main Go thread. There are times we want to share resources between multiple threads. Go prefers not sharing the variables of one thread with another because this adds a chance of deadlock and resource waiting. There is another way to share resources between Go routines: via go channels.

### Channels

------------------------


We can pass data between two go rotuines using channels. While creating a channel its necessary to specify which kind of data the channel receive. Lets create a simple channel with string type as follows.

```go

c := make(chan string)

```

With this channel, we can send string type data. We can both send and receive data in this channel:

```go

package main

import "fmt"

func main(){
  c := make(chan string)
  go func(){ c <- "hello" }()
  msg := <-c
  fmt.Println(msg)
}

```

The receiver Channels wait until the sender sends data to the channel.

### One way Channel

------------------------

There are cases where we want a Go routine to receive data via the channel but not send data, and also vice versa. For this, we can also create a **one-way channel**. Let’s look into a simple example:

```go

package main

import (
    "fmt"
)

func main(){
    ch := make(chan string)
    go sc(ch)
    fmt.Println(<-ch)
}

func sc(ch chan<- string) {
    ch <- "hello"
}

```

In the above example, sc is a Go routine which can only send messages to the channel but cannot receive messages.

### Organizing multiple channels for a Go routine using select

---------------------------

There may be multiple channels that a function is waiting on. For this, we can use a select statement. Let us take a look at an example for more clarity:

```go

package main

import (
 "fmt"
 "time"
)

func main() {
 c1 := make(chan string)
 c2 := make(chan string)
 go speed1(c1)
 go speed2(c2)
 fmt.Println("The first to arrive is:")
 select {
 case s1 := <-c1:
  fmt.Println(s1)
 case s2 := <-c2:
  fmt.Println(s2)
 }
}

func speed1(ch chan string) {
 time.Sleep(2 * time.Second)
 ch <- "speed 1"
}

func speed2(ch chan string) {
 time.Sleep(1 * time.Second)
 ch <- "speed 2"
}

```

In the above example, the main is waiting on two channels, c1 and c2. With select case statement the main function prints, the message sends from the channel whichever it receives first.

### Buffered Channel

---------------------

There are cases when we need to send multiple data to a channel. You can create a buffered channel for this. With a buffered channel, the receiver will not get the message until the buffer is full. Let’s take a look at the example:

```go

package main

import "fmt"

func main(){
  ch := make(chan string, 2)
  ch <- "hello"
  ch <- "world"
  fmt.Println(<-ch)
}

```