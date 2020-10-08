- Goroutines: In go, each concurrently executing block of scope is called `goroutine`. When you execute the program by doing `go run ` then its only goroutine is the one that called `main` function so we call it `main goroutine` while, New go routines can be created by `go` statement. A go statement causes the function to be called in a newly created goroutine. The go steatement itself complete immediately:

```go
f()                // calls f wait for it to return
go f()             // calls f concurrently and immediately goes to next statemtent, while content of f is still executing.
```
