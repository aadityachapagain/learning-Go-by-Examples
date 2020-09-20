1. If entity is decleared within a function, it is local to that function.
2. The Case of first letter of name determines its visibility across package boundries.If the names begin with upper-case letter, it is exported, which means it is visible and accessible outside of its own package and maybe refer to by other parts of program, as with Printf in the fmt package. Package names themselves are always in lower case.
3. variable declarations `var name type = expression`.Either type or `= expression` can be ommited but not both. So, if type is ommited,var type is identified by expression and expression initial value is attached to variable. or if `= expression` is ommited variable type is defined by `type` and initial default value will be assigned to variables 0 for int, float, false for bool, "" for string, nil for interfaces.
   Short Variable declarations `name := expression` and type of name is defined by type of declarations like `t := 0.0` is float, `t := bufio.NewScanner(os.stdin)` is bufio object, `p := "string"` is string and many more.
4. Pointers, A pointer value is the address of the variable.Thus, pointer is the location at which value is stored.Not every value has an address and but every variable does.If a variable is declared as `var x int` the expression &x ("address of x") yields a pointer to an integer variable, that is, a value of type *int , which is pronounced "pointer to int". If this value is p, we say "p points to x" or equivalently "p contains adress of x". The variable to which p points is written as *p
5. The New Function: creating unnamed variable with new , return new variable with address. eg `p:= new(int)` return p of type \*int , points to an unnamed int variable.
6. Type Declarations most often appear at package level, where the named type is visible through out the package, and if the name is exported its accessible from other package as well.
   type name underlying-type

7 Package Initialization: begins by initializing package level variables in the order in which they are declared, except that dependencies are resolved first.
if the package have multiple .go files, they are initialized in the order in which the files are given to the compiler.Each variable declared at package level starts life with the value of its initializer expression, if any, but for some variables, like tables of data, an initializer expression may not be the simplest
way to set its initial value. In that case, the init function mechanism may be simpler. Any file may contain any number of functions whose declaration is just

```go
func init () { */...*/ }
```

such init function cannot be called or referenced. otherwise they are just a normal functions. Within each file, init functions are automatically executed when the program starts, in the order in which they are declared.
