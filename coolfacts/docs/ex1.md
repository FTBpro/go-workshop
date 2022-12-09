# Part 1
###Create "Hello, World!" Application

In this exercise, you will create a basic "Hello, World!" application. You will learn how to build and run a Go program and basic terminology.

## Prerequisites
TODO: addddddd

## The Starting Point
To get the initial application, run the next command:
```commandline
$ git clone --branch v1-hello-world https://github.com/FTBpro/go-workshop.git
```
This will clone the first branch which we will use for building and running our very first Go application.
The application will simply print "Hello, World!". Go on and navigate to `go-workshop/coolfacts` directory. You should be on branch `v1-hello-world`.

## Your Goal
Execute a Go program that prints to the screen "Hello, World". Understand terms like packages, module, main, and get to know basic Go commands.

## Step 1
When looking around the program, you will notice two files inside the folder `coolfacts`
- main.go
- go.mod

### main.go
The _main package_ is the starting point for every Go application. A _package_ in go is a way to organize and group functions, and it is nothing but a directory inside your Go workspace, containing one or more Go source files or other packages. Go package name usually is the same as the directory name. package main is almost the only package that usually isn't like the folder name.

In Go, the _main package_ is a special package which is used with the executable commands. The `main()` function is a special type of function and it is the entry point of the executable program. When running the application, Go will automatically call `main()`

In every source file, the first line is the package name. We see here the first line is
```go
package main
```

### go.mod
`go.mod` file is the root of a Go module. Go module is a collection of packages, stored in a file tree. Go code is grouped into packages, and packages are grouped into modules. Your module specifies dependencies needed to run your code, including the Go version and the set of other modules it requires.

The first line in the `go.mod` defines the module's path, which is also the import path.

In this example, you can see the module's path is 
```goregexp
github.com/FTBpro/go-workshop/coolfacts
```
The module's path is the import for consumers of this module. When importing for the first time, Go will fetch the module from the specified path.

Currently, our module has no dependencies.

