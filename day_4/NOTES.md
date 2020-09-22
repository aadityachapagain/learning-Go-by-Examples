1. Conversion between String and Numbers.
   can be done with package _strconv_ (with `fmt.Sprintf` or `strconv.Itoa`. Examples,

```go
x := 123
y := fmt.Sprintf("%d", x)
fmt.Println(y, strconv.Itoa(x))
```

2. FormatInt and FormatUint can be used to format numbers in different base.

```go
fmt.Println(strconv.FormatInt(Int64(X),2))
```

3. The constant Generator iota.
   In a const declaration,the value of iota begins at zero and increments by one for each item in the sequence. Examples

```go
  type Weekday int
  const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
  )
```

This declares sunday to be 0 , Monday to be 1 and so on.
