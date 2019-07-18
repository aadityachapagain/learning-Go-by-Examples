# GO Basics

This part of the go programming tutorial series covers GO packages, many go packages, more go packages, custom packages, Packages Documentation, Error handling, Custom error, Panic and Defer.

## Packages

We write all the code in packages, upto now, we been writing the codes in main packages.  
The main Package is the entry point for the program execution, so probably for huge project we dont want to write there.  
Talking about Packages, there are loads of built in packages in Go. The only packages we have been using for this tutorial is only `fmt` package.


### Installing The Packages

-------------------------

```bash

go get <package-url-github>

go get github.com/satori/go.uuid

```

The packages we installed are saved inside the GOPATH env which is our work directory. You can see the packages by going inside the pkg folder inside our work directory cd $GOPATH/pkg.

### Creating a Custom package

-------------------------

Lets start by creating a folder my_package:

```shell

> mkdir my_package
> cd my_package

```

To create a custom package we need to first create a folder with the package name we need. Let’s say we are building a package person. For that let’s create a folder named person inside custom_package folder:

```bash

> mkdir person ; cd person

```
Now let's create a file person.go inside this folder.

```go

package person

func Description(name string) string {
  return "The person name is: " + name
}
func secretName(name string) string {
  return "Do not share"
}

```
Now, we need to install the package so that it can be imported and used. So, let's install the package:

```bash

> go install

```
Now let’s go back to the custom_package folder and create a main.go file:

```go

package main

import (
    "my_package/person"
    "fmt"
)

func main(){
    p := person.Description("Milap")
    fmt.Println(p)
}

```

Here we can now import the package person we created and use the function Description. Note that the function secretName we created in the package will not be accessible. In Go, **the method name starting without a capital letter will be private**.


### Packages Documentation

-------------------------

Go has built in support for documentation for packages. Run the following command to generate documentation:

```bash

godoc person Description

```

This will generate the Documentation for Description function inside our person packages. To see documentation run a webserver using the following command:

```bash

godoc -http=":8080"

```

Now go to the URL http://localhost:8080/pkg/ and see the documentation of the package we just created.

### Some built-in Go packages

-------------------------

#### fmt

The packages implement I/O operations and functions. We have already used the packages for printing out to stdout.

#### json

Another useful package in Go is the json package. This helps to encode/decode the JSON. Let’s take an example to encode/decode some json:

#### Encode

```go

package main

import (
  "fmt"
  "encoding/json"
)

func main(){
  mapA := map[string]int{"apple": 5, "lettuce": 7}
  mapB, _ := json.Marshal(mapA)
  fmt.Println(string(mapB))
}


```

#### Decode

```go

package main

import (
  "fmt"
  "encoding/json"
)

type response struct {
  PageNumber int `json:"page"`
  Fruits []string `json:"fruits"`
}

func main(){
  str := `{"page": 1, "fruits": ["apple", "peach"]}`
  res := response{}
  json.Unmarshal([]byte(str), &res)
  fmt.Println(res.PageNumber)
}
```

While decoding the json byte using unmarshal, the first argument is the json byte and the second argument is the address of the response type struct where we want the json to be mapped to. Note that the json:”page” maps page key to PageNumber key in the struct.

### Error Handling

----------------------

Errors are Undesired and Unexpected result of the program. Lets say we are making a API CALL to extend our service. The API call may be successful or could fail. An error in a Go Program can be recognized when an error type is present. Let's see the example:

```go

resp, err := http.Get("http://example.com/")

```

Here, API call to the error object may pass or could fail. We can check if the error is null or present and handle the response accordingly.

```go

package main

import (
  "fmt"
  "net/http"
)

func main(){
  resp, err := http.Get("http://goisawsome.com")

  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(resp)
}

```

### Returning Custom Error from Function

-------------------------

When we are writing a function of our own, there are cases when we have errors. These errors can be returned using error object.

```go

func Increment(n int) (int, error) {
  if n < 0 {
    // return error object
    return nil, errors.New("math: cannot process negative number")
  }
  return (n + 1), nil
}
func main() {
  num := 5
 
  if inc, err := Increment(num); err != nil {
    fmt.Printf("Failed Number: %v, error message: %v", num, err)
  }else {
    fmt.Printf("Incremented Number: %v", inc)
  }
}

```

Most of the packages that are built in Go, or external packages we use, have a mechanism for error handling. So any function we call could have possible errors. These errors should never be ignored and always handled gracefully in the place we call these functions, as we have done in the above example.

### Panic

-------------------------

Panic is something that is unhandled and is suddenly encountered during a program execution. In Go, panic is not the ideal way to handle  exceptions in a program.It is recommended  to use a error object instead. When a panic occurs , the program execution get's halted. The thing that get's executed after panic is defer.

```go

package main

import (
  "fmt"
)

func main(){
  f()
  fmt.Println("Returned normally from f.")
}

func f(){
  defer func(){
    if r := recover (); r != nil {
      fmt.Println("Recovered in f", r)
    }
  }()
  fmt.Println("Calling g.")
  g(0)
  fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}

```

### Defer

-----------------------

Defer is something that will always get executed at the end of the function.  
  
In the above example, we panic the execution of the program using panic(). As you notice, there is a defer statement which will make the program execute the line at the end of the execution of the program. Defer can also be used when we need something to be executed at the end of the function, for example closing a file.
