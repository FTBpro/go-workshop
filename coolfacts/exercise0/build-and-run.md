# Building and Running Go Applications

The Go compiler creates an executable that is run directly on your system.

## Go Build

```bash
usage: go build [-o output] [-i] [build flags] [packages]
```

For example, if you want to build `exercise0`: 

```bash
$ cd coolfacts/exercise0
$ go build .
```

This will create a binary called `exercise0` in the current directory.
(`exercise0.exe` on windows)

```bash
$ ls -l
-rwxr-xr-x  [...] exercise0
-rw-r--r--  [...] main.go
```

And now run the binary directly on your operating system

```bash
$ ./exercise0
Hello World!
```

For development, there's an easier way
for locally running a Go application

```bash
$ go run .
Hello World!
```

## Advanced

If you want to know more, the best place to start is the output of

```bash
$ go help build
```
