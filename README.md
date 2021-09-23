# learning-go

Basic introductory programs I wrote while learning go and some helpful notes

## Notes on Go
### Overview
- Go is based mostly on C and influenced by other languages (C++, Java, C#, etc.).
- Go implements many OOP concepts such as encapsulation with types and structs, polymorphism with interfaces etc.
- Go does not support type inheritance.
- Go is a compiled language, and it is statically typed however the Go tool can run a file without precompiling
- There is no external VM like JVM so run time must be included in the compiled code.
- Go has some object-oriented features but there is no inheritance, classes, overloading, structured exception handling (no try, catch etc.), implicit numeric conversions etc. This was done to improve readability.

### Syntax
- Go is case-sensitive
- Methods/fields use initial upper case characters to export. This is similar to the 'public' keyword in Java. Lowercase is similar to the 'private' keyword in Java.
- White space ends a statement as the lexer adds semicolons for you.
- Code blocks are wrapped in braces. The first preceding brace must be on the same line as the preceding line of code due to the white space sensitivity the language has.

### Complex Types
#### Data collections
- Arrays and slices: manage ordered data collections.
- Maps and structs: manage aggregates of values.

#### Language Organisation
- These exist in other languages but are not always considered types in those languages.
- Functions, Interfaces and Channels are all types. This makes it possible to pass a function into another function.

#### Data management
- Go supports pointers which are reference variables that point to an address in memory to refer to another value.

### Memory Allocation
- The ```new()``` function allocates but does not initialise memory. Results in a zeroed storage but returns a memory address.
- The ```make()``` function allocates and initialises memory. Results in non-zeroed storage and returns a memory address.
- Memory is deallocated by the garbage collector. Objects out of scope or set to ```nil``` are eligible.
- Go was designed to have a very low latency garbage collector.