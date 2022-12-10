# Part 1
### Create "Hello, World!" Application

In this exercise, you will create a basic "Hello, World!" application. You will learn how to build and run a Go program and basic terminology - packages, module, go.mod, main, etc.
Feel free to skip this part if you went over it already. 

## Prerequisites
TODO: addddddd

## The Starting Point
Be sure to be on branch `v1-hello-world`. Go on and navigate to `go-workshop/coolfacts` directory, you will see two files: `main.go` and `go.mod`.

### main.go
The _main package_ is the starting point for every Go application. A _package_ in go is a way to organize and group functions, and it is nothing but a directory inside your Go workspace, containing one or more Go source files or other packages. Go package name usually is the same as the directory name. package `main` an example of a package that is different from the folder name, we'll understand why.

In Go, the _main package_ is a special package which is used with the executable commands. The `main()` function is a special type of function and it is the entry point of the executable program. When running the application, Go will automatically call `main()`

In every source file, the first line is the package name. We see here that the first line is
```go
package main
```

### go.mod
`go.mod` file is the root of a Go module. Go _module_ is a collection of packages, stored in a file tree. Go code is grouped into packages, and packages are grouped into modules. Your module specifies dependencies needed to run your code, including the Go version and the set of other modules it requires.

The first line in the `go.mod` defines the module's path, which also will be used as the import path for packages within this module.

In this example, you can see the module's path is 
```goregexp
github.com/FTBpro/go-workshop/coolfacts
```
The module's path is the import for consumers of this module. When importing for the first time, Go will fetch the module from the specified path.

Currently, our module has no dependencies, as you can see our `go.mod` file is empty.

# Print
For printing `Hello, World!` we'll use go package `fmt`. Package `fmt` implements formatted I/O with functions analogous to C's `printf` and ``s`canf. The format 'verbs' are derived from C's but are simpler. You can head over https://pkg.go.dev/fmt to see its documentation.

In here we can use method `fmt.Println` which prints to the stdoud and ends with a new line.

Go on and add the import for `fmt` and the call to the method:
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}
```
Since its not an external dependency, we don't need to add anything in our `go.mod` file.

# Building and Running Go Applications

The Go compiler creates an executable that is run directly on your system.

The `go build` command compiles the packages, along with their dependencies, and output a binary file which you can execute.

```bash
usage: go build [-o output] [-i] [build flags] [packages]
```

This command needs to be called from the root of your module (that it, where the `go.mod` resides) and it receives a path to the main package. 

In this application our main is also in the root, so you can just run `go build`:

```bash
.../coolfacts$ go build 
```

By default, the binary name that is created is in the name of the folder name. In here you can see that a new file created named `coolfacts`. (`coolfacte.exe` on windows)

Now you can run the binary directly on your operating system.

```bash
.../coolfacts$ ./coolfacts
Hello World!
```

You can also build and run the application in one command using `go run`:

```bash
.../coolfacts$ go run .
Hello, world!
```
> `go run` must receive a path to the main package


If you want to know more, the best place to start is the output of

```bash
$ go help build
```

Or https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies
