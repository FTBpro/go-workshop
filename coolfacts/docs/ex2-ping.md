# Part 2

In this exercise, you will create a simple server that answering to `/ping`.

## The Starting Point
To get the initial application, run the next command:
```commandline
$ git clone --branch v2-initial-server-ping https://github.com/FTBpro/go-workshop.git
```
This will clone the branch `v2-initial-server-ping` which we will use for building and running our web server.

## Your Goal

After implementing all the TODOs, you will have an application that runs a server that answer PONG to ping.

For every other path, the server will return not-found response with the following JSON:
```json
{
  "error": "path wasn't found"
}
```

## Getting Started
Take a look around the program. You will notice one folder - `cmd`

In Go programs, `cmd` folder is convention for having more than one binary (e.g application).
Usually the directory name inside cmd matches the binary name, and each folder has package `main`. In this exercise we only have one application, but in the next ones, we will also have client application.

Inside the `cmd` we have folder `coolfact_server`. This folder includes `main` package, with two files:
- `main.go`: file that includes our `main()` function
- `server.go`: this is the transport of the application. In this file we will implement the http handlers for the server APIs.

## Notice `main.go`

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

Inside of the `main()` function, we're initializing the server, and sending it to `http.ListenAndserve` method.
In Go, a package name is only the last param, and each type in Go has a name composed of the package name and the type identifier. Although the import path is `net/http`, the package name is `http`, and the identifier is `ListenAndServe`. 

The method `http.ListenAndServe` receives an address and an `http.Handler`:
- address: This is the address it will listen on the TCP network.
- `http.Handler`: For every incoming connection, it will use the `http.Handler` to handle the requests, which defined as:
  ```go
  package http
  
  type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
  }
    ```
  The handler should handle the incoming `http.Request` and write the response to the `ResponseWriter`. In our application the http.Handler will be the `server`. 

What you will have to complete is:

## Step 1 - `cmd/server.go`
The server is responsible to respond to the incoming HTTP request, it will have `/ping` route, for answering with PONG.

You can notice that the server has method `ServeHTTP`. This is the entry point for every incoming HTTP request. The server implement `http.Handler`. We have switch case according to the HTTP method and path.

The only recognized path is `GET /ping`. On every other path it returns `HandleNotFound` which you will implement. In real world applications, we will usually use a framework wrapping Go http which will provide us with more advanced routing mechanism.

- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Finish implement `NewServer` function.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `HandlePing` method.
  - Response with status 200 (const `http.StatusOK`)
  - Write `PONG` as the response.
  - Clue 1 - check the `http.ResponseWriter` interface. Notice its methods and decide what you can use.
  - Clue 2 - The `http.ResponseWriter` implements `io.Writer` (check why). Therefore, we can use a method in the Go package `fmt` to write the response
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `HandleNotFound` and `HandleError` methods
  - The JSON response for this methods should be the same:
  ```json
  {
    "error": <error message>
  }
  ```
  - In case of not found return status 404 (Not Found Error), in case of other error, return status 500 (Internal Server Error)

### Building and Running
The main package of our server-application lives inside coolfacts/cmd/coolfacts_service.go, but the go.mod is under coolfacts directory.
The go build must be called from the root of the module, in here it will be called from folder coolfacts. The path we will pass to go build will be the path to the main we want to run:
```commandline
.../coolfacts$ go build ./cmd/coolfacts_server/.
```
This command will create executable name coolfacts_server which we can run by `./coolfacts_server`

Alternatively we can build and run in one single command:
```commandline
.../coolfacts$ go run ./cmd/coolfacts_server/.
```

If everything is implemented, this is what the final result should look like: (TODO:oren): not found path in gif
![factsgif](https://user-images.githubusercontent.com/5252381/204143457-6eaf59d3-6c52-4fbb-8d2a-19d22436cbd8.gif)

# Full walkthrough
In the following section you fill find a full walk through. Use it in case you are stuck

## Step 1 - The HTTP transport layer (server)
We implement `NewServer`, simply by returning a reference to the `service`
```go
type server struct {
}

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

Next, we will write “PONG”, We notice that the `http.ResponseWriter` implements `io.Writer`
Hence we can use all bunch of methods for writing the response. Here we only need to return text, so we will use `fmt.Fprintf` which receives `io.Writer`.

### `HandleNotFound` and `HandleError`:
These methods are almost identical, except of the status. `HandleNotFound` returns status 404 - `http.StatusNotFound`
And `HandleError` returns status 500 - `http.StatusInternalServerError`
```go
func (s *server) HandleNotFound(w http.ResponseWriter, err error) {
  w.WriteHeader(http.StatusNotFound)
  w.Header().Set("Content-Type", "application/json")
 
  response := map[string]interface{}{
     "error": err,
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

# Finish!
<img src=https://user-images.githubusercontent.com/5252381/204150018-b6f1a5af-c557-443a-9301-7c0e98a4a3f7.png width="91">

Congratulation! You've just implemented a web server and BL for returning facts.

In the following exercise we will add an API for creating a fact, and a client application for calling our server.
