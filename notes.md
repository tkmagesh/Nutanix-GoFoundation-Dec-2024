# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com

## Software Requirements
- Go Tools (https://go.dev/dl)
- Visual Studio Code (https://code.visualstudio.com)

## Schedule
- Session-01    : 1:50 hrs
- Tea Break     : 20 mins
- Session-02    : 1:50 hrs

## Methodology
- No powerpoint
- Code & Discussion

## Repository
- https://github.com/tkmagesh/nutanix-gofoundation-dec-2024

## Why Go?
- Simplicity
    - ONLY 25 keywords (package, import, func, var, const, struct, func, return, if, else, switch case, break, continue, fallthrough, return, for, range, select, type, struct, interface, any, go, map, iota etc)
    - No access modifiers (no public/private/protected)
    - No classes (only structs)
    - No inheritence (only composition)
    - No reference types (everything is a value in go)
    - No pointer arithmatic
    - No exceptions (only errors & errors are just values returned from functions)
    - No try-catch-finally construct
    - No implicit type coversion
- Performance
    - equivalent to c++
    - compile to native code
    - built in tooling support for cross compilation
- Managed Concurrency
    - Lightweight concurrecy representation called as goroutines
    - goroutines are cheap (4 KB)
    - Builtin scheduler
    - Concurrency support is built in the language itself
        - "go" keyword
        - "chan" channel data type
        - "<-" channel operator
        - "range" construct
        - "select-case" construct
    - Standard Library API support
        - "sync" package
        - "sync/atomic" package
## To create a build
```shell
go build <file_name.go>
# ex:
go build 01-hello-world.go
```
## To compile and execute
```shell
go run <file_name.go>
# ex:
go run 01-hello-world.go
```
## Cross Compilation
### List the Go tooling system environment variables
```shell
go env
```
### Env variables for cross compilation
```shell
go env GOOS GOARCH
```
### List of supported platforms
```shell
go tool dist list
```
### How to cross compile
```shell
GOOS=windows GOARCH=amd64 go build 01-hello-world.go
```

## Data Types
- bool
- string
- integers
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integers
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- floating points
    - float32
    - float64
- complex numbers
    - complex64 (real[float32] + imaginary[float32])
    - complex128 (real[float64] + imaginary[float64])
- type alias
    - byte
    - rune (unicode code point)

## Variables in function scope & package scope
### function scope
    - Can use :=
    - Cannot have unused variables
### package scope
    - Cannot use :=
    - Can have unsed variables

## Zero values

| type | zero value |
| -------|-------- |
| int, uint, float | 0 |
| string | "" |
| bool | false |
| func | nil |
| struct | struct instance |
| pointer |nil |
| interface | nil | 

## Functions
- Functions can return more than one result
- Variadic functions
- Anonymous functions
- Higher Order functions
    - Assign functions as values to variables
    - Pass a function as an argument
    - Return a function as a return value
- Deferred functions

## Collections
### Array
- Fixed sized typed collection
### Slice
- Varying sized typed collection
- A slice is effectively a pointer to an underlying array
### Map
- Varying sized typed collection of key/value pairs

## Error Handling
- errors are values returned from a function
- errors are values that implement "error" interface
    - error values should have "Error()" method implemented
- Creating an error
    - errors.New()
    - fmt.Errorf()
    - Custom type implementing "error" interface

## Panic & Recovery
### Panic
- Represents the state of the application where the application execution cannot proceed further
- When a panic occurs, all the deferred functions scheduled are executed and the application is shut down
- Use "panic()" to create a panic
### Recovery
- "recover()" returns the error that resulted in the panic
- typically used in the deferred functions

## Modules & Packages
### Module
- Any code that need to versioned and deployed together
- Typically, a folder with go.mod file
#### go.mod file
    - manifest file for the module/application
    - name
        - advisable to include the repo path
    - targetted go runtime version
    - dependencies
#### To create a module
```shell
go mod init <module_name>
```
#### To run a module
```shell
go run .
```
#### To build a module
```shell
go build .
# OR
go build -o <binary_name> .
```

### Package
- internal organization of a module
- typically, a folder with go files
- all the code in all the files of the package are considered to belong to the package
- public entity names MUST start with uppercase