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
    - ONLY 25 keywords
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