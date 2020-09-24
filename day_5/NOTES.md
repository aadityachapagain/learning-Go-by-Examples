1. append function:
   The builtin append function append items to slices,

```go
var runes []rune
for _, r := range "Hello, world" {
  runes = append(runes, r)
}
fmt.Printf("%q\n", runes)
```

though this specific problem can be solved by using built in conversion `[]rune("hello, world")` 2.
