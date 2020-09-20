1. If entity is decleared within a function, it is local to that function.
2. The Case of first letter of name determines its visibility across package boundries.If the names begin with upper-case letter, it is exported, which means it is visible and accessible outside of its own package and maybe refer to by other parts of program, as with Printf in the fmt package. Package names themselves are always in lower case.
3.
