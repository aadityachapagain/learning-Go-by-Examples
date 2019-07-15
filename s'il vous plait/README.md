# GO BASICS

This part of the go programming tutorial gives the basic syntax usage of conditional statement in GO.

## Conditional Statement

### if else

------------------
  
For conditional statements, we can use if-else statements as shown in the example below. Make sure that the curly braces are in the same line as the condition is:

```go

if num := 9; num < 0 {
    fmt.Println(num, "is negative !")
} else if num < 10 {
    fmt.Println(num, " has 1 digit")
} else {
    fmt.Println(num, " has multiple digits")
}

```

### switch case

------------------

Switch case helps organize mulitple conditional statements. The following statements shows the simple switch case statements:

```go

i := 0

switch i {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
case 3:
    fmt.Println("Three")
default:
    fmt.Println("None")
}
```

### Looping

------------------

Go has a single keyword for the loop. A signale for loop command helps to achieve different kind of loops:

```go
i := 0

sum := 0

for i < 10 {
    sum += 1
    i++
}

fmt.Println(sum)

```

The above example is similar to a while loop in C. The same for statement can be used for a normal for loop:

```go
sum := 0

for i := 0; i < 2000; i++ {
    sum += i
}
fmt.Println(sum, "is total sum.")

```

Infinite Loop in Go:

```go
for {

}
```
