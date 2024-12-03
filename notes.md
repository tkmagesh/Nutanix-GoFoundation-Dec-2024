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