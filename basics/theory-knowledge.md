## Go Run vs Go Build

With `go build` we are making a persistent executable file.
With `go run` we compile and run in a single step without creating an executable file.

1. go first compiles go source code into a temp executable file in memory -> not stored in storage but ram
2. once it is created, it is immediately run.
3. after the program is completed running, the go executable is discarded.

it exists only for current session.

## Go Compiler

When we run the program, this comes into the picture.
Translates human readable to machine code (binary) language -> this process is called compilation
executables are machine specific

Runtime: once the code is compiled.
manages the execution of the program, memory allocation, garbage collection and etc

go code -> compiler: translates to instruction -> runtime: execute the instructions and produce final output of program

## Standard Library

Pre written and reusable code modules that adress common prograaming tasks.
Can import using the syntax:
`import "fmt"`

### Import Statement

Integrates external libraries in our code.
While making the program entire package is available: encompasess all functions, types, variables, constants, and other definitions within a package -> but when compiling and running only relevant parts are kept while others are discarded -> "Tree Shaking"

**Tree Shaking**: Technique used to eliminate unused code in the final bundle/code -> reducing size and improving performance.
