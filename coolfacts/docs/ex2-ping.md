# Part 2

In this exercise, you will create a simple server that answering to `/ping`.

## The Starting Point

Be sure to be on branch `v2-initial-server-ping`. Go on and navigate to `go-workshop/coolfacts` directory. You will see a bunch of files and TODOs. We will go over them in the rest of this document.

Try to go over this doc and implement by yourself. If you are stuck there is a full walkthrough in the end, and for reference, you can look inside the branch `v2-initial-server-ping`. All the implementations for the TODOs can be found there, or simply by looking at the pull-request: https://github.com/FTBpro/go-workshop/pull/28/files 


## Your Goal

After implementing all the TODOs, you will have an application that runs a server that answer "PONG" to "ping".

For every other path, the server will return not-found response with the following JSON:
```json
{
  "error": "path wasn't found"
}
```

## Getting Started
Take a look around the program. You will notice one folder - `cmd`

In Go programs, `cmd` folder is a convention when a program has more than one binary (e.g application).
Usually the directory name inside the `cmd` matches the binary name, which on default will be the name of the binary. Each directory in the `cmd` (e.g application) has package `main`. In this exercise we only have one application, but in the next ones, we will also have a client application.

In here, we have inside the `cmd`, a folder named `coolfact_server`. This represent our server application, and includes a `main` package, with two files:
- `main.go`: file that includes our `main()` function
- `server.go`: this is the transport of the application, its responsibility is to receive the http request and to apply the corresponding business-logic. In this file we will implement the http handlers for the server APIs.

## Step 0 - Notice `main.go`

Let's have a look at `coolfacts_server/main.go`:
```go
package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  fmt.Println("Hello, Server!")

  server := NewServer()

  log.Println("starting server on port 9002")
  log.Println("You can go to http://127.0.0.1:9002/ping")
  if err := http.ListenAndServe(":9002", server); err != nil {
    panic(fmt.Errorf("server crashed! err: %w", err))
  }
}
```
We first can notice the imports `"fmt"`, `"log"` and `"net/http"`. These are standard Go libraries.

Inside the `main()` function, we're initializing the server, and passing it to `http.ListenAndserve` method.
In Go, a package name is only the last param of the import path (usually), and each type in Go has a name composed of the package name and the type identifier.

For example, let's look at the method `http.ListenAndServe` - although the import path is `net/http`, the package name is `http`, and the identifier is `ListenAndServe`.

This method receives an address and an `http.Handler`:
- address: This is the address it will listen on the TCP network.
- `http.Handler`: For every incoming connection, it will use the `http.Handler` to handle the requests, which is defined as:
  ```go
  package http
  
  type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
  }
    ```
  The handler should handle the incoming `http.Request` and write the response to the `ResponseWriter`. In our application the `http.Handler` will be our `server`, which we initialize by `NewServer()`

The type `http.Handler` is a very central pillar of Go. It makes abstraction over http real easy. Besides this type, there is also the type `http.HandlerFunc`:
```go
package http

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
  f(w, r)
}
```
The `HandlerFunc` is an adapter to allow the use of ordinary functions as HTTP handlers. 
Due to the method `ServeHTTP`, this type also implements `http.Handler`. We will see it in farther exercises. 


## Step 1 - `cmd/server.go`
The server is responsible to respond to the incoming HTTP request, it will have `/ping` route, for answering with PONG.

You can notice that the server has method `ServeHTTP`. This is the entry point for every incoming HTTP request. The server implements `http.Handler`. We have a switch case according to the HTTP method and path, and then we applies the right method. You can already notice that the methods `HandlePing` and `HandleNotFound` are `http.HandlerFunc`. This will help us in exercise 7 :) 

For now, the only valid path is `GET /ping`. On every other path it calls `HandleNotFound` which you will implement. In real world applications, we will usually use a framework wrapping Go http which will provide us with more advanced routing mechanism. 

Specifically, what you will have todo is:

- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Finish implement `NewServer` function.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `HandlePing` method.
  - Response with status 200 (const `http.StatusOK`)
  - Write `PONG` as the response.
  - Clue 1 - check the `http.ResponseWriter` interface. Notice its methods and decide what you can use for setting the status code and the text.
  - Clue 2 - The `http.ResponseWriter` implements `io.Writer` (check why). Therefore, we can use a method in the Go package `fmt` to write the response
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `HandleNotFound`
  - Return status code 404 and set header "Content-Type" to "application/json" (check the interface `ResponseWriter` for how to set the header)
  - The JSON response for a not found path should be:
  ```json
  {
    "error": "path <http method + path> not found"
  }
  ```
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `HandleError`
  - Like not found, but return status 500, and the JSON response:
  ```json
  {
    "error": <the error message>
  }
  ```

## Building and Running
The main package of our server-application lives inside coolfacts/cmd/coolfacts_service.go, but the go.mod is under coolfacts directory.
The go build must be called from the root of the module, in here it will be called from folder _coolfacts_. The path we will pass to `go build` will be the path to the `main` we want to run:
```commandline
.../coolfacts$ go build ./cmd/coolfacts_server/.
```
This command will create executable name _coolfacts_server_ which we can run by `./coolfacts_server`

Alternatively we can build and run in one single command:
```commandline
.../coolfacts$ go run ./cmd/coolfacts_server/.
```

If everything is implemented, this is what the final result should look like: (TODO:oren): not found path in gif
![factsgif](https://user-images.githubusercontent.com/5252381/204143457-6eaf59d3-6c52-4fbb-8d2a-19d22436cbd8.gif)

# Full walkthrough
In the following section you fill find a full walk through. Use it in case you are stuck. 

## Step 1 - The HTTP transport layer (server)
We implement `NewServer`, simply by returning a reference to the `service`
```go
type server struct {}

func NewServer(service Service) *server {
    return &server{}
}
```

###`HandlePing` method:
```go
func (s *server) HandlePing(w http.ResponseWriter) {
    w.WriteHeader(http.StatusOK)

    if _, err := fmt.Fprint(w, "PONG"); err != nil {
        fmt.Printf("ERROR writing to ResponseWriter: %s\n", err)
        return
    }
}
```
First, we need to write the status. For this we use `w.WriteHeader(http.StatusOK)`

Next, we will write “PONG”. We can notice that the `http.ResponseWriter` implements `io.Writer`.
Hence, we can use all bunch of methods for writing the response. Here we only need to return text, so we will use `fmt.Fprintf` which receives `io.Writer`.
We already see the power of go interfaces, and especially interfaces with one method. `fmt.Fprintf` has no clue about `http`, and `ResponseWriter` has no clue about `fmt`.

### `HandleNotFound` and `HandleError`:
These methods are almost identical, except of the error message and the status. `HandleNotFound` returns status 404 - `http.StatusNotFound`
And `HandleError` returns status 500 - `http.StatusInternalServerError`
```go
func (s *server) HandleNotFound(w http.ResponseWriter, err error) {
    w.WriteHeader(http.StatusNotFound)
    w.Header().Set("Content-Type", "application/json")

    response := map[string]interface{}{
        "error": "path wasn't found",
    }

    if err := json.NewEncoder(w).Encode(response); err != nil {
        fmt.Printf("HandleGetFacts ERROR writing response: %s", err)
    }
}

func (s *server) HandleError(w http.ResponseWriter, err error) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Header().Set("Content-Type", "application/json")

    response := map[string]interface{}{
        "error": err,
    }

    if err := json.NewEncoder(w).Encode(response); err != nil {
        fmt.Printf("HandleGetFacts ERROR writing response: %s", err)
    }
}
```

Since we are returning a JSON, we can use the Go `json` package, which can write a JSON encoding of the response to the writer.
The `NewEncoder` method receives `io.Writer`. Again we can notice how helpfull is `io.Writer` interface.
We can send the `http.ResponseWriter` to different packages methods, and use a different encoding, and the packages are totally agnostic to HTTP.


# Finish!
<img src=https://user-images.githubusercontent.com/5252381/204150018-b6f1a5af-c557-443a-9301-7c0e98a4a3f7.png width="91">

Congratulation! You've just implemented a web server!

In the following exercise we will add an API for getting facts, create fact, client and introduce other packages.
