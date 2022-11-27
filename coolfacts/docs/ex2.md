# Part 2

In this exercise, you will create a server that answering to `/ping` and `/facts` to return a list of cool facts.

## The Starting Point
To get the initial application, run the next command:
```commandline
$ git clone --branch v2-initial-server https://github.com/FTBpro/go-workshop.git
```
This will clone the branch `v2-initial-server` which we will use for building and running our web server.

## Your Goal

After implementing all the TODOs, you will have an application that runs a server that listens on 127.0.0.1:9002.
The server will answer to two endpoints:
1. Ping:
```json
GET /ping
```
simply return `PONG`
2. 
```josn
GET /facts
```
Returns JSON
```json
{
  "facts": [
    {
      "image": "...",
      "description": "..."
    }
    //...
  ]
}
```
For every other path, it will return status code 404, with the following JSON response:
```json
{
  "error": "path wasn't found"
}
```

## Getting Started
Take a look around the program. You will notice three folders - `coolfact`, `inmem`, `cmd`
### coolfact
This is a package containing the entity `fact` and the service for implementing the use case (business logic). In this application we won't have much BL, if any. Our application is a very simple CRUD app, and the service will mainly call the repo as we will see.

### inmem
In this package we will implement the facts repository. We will use in memory. In Go, packages names are very important. A package name should say what it provides, not what it contains. We won't have packages names like `models`, `utils`, `common`, etc.

A package’s name provides context for its contents, making it easier for clients to understand what the package is for and how to use it. Go package `time` indicates that it contains functionality for handling times, probably a struct and behaviour for measuring and displaying time.

Package `inmem` imply that it use memory mechanism. If we will have SQL database, we may have package `sql` that implies it lets you speak SQL.

### cmd
In Go programs, `cmd` folder is convention for having more than one binary (e.g application).
Usually the directory name inside cmd matches the binary name, and each folder has package `main`. In this exercise we only have one application, but in the next one we will have client application as well 

Inside the `cmd` we have folder `coolfact_server`. This folder includes `main` package, with two files:
- `main.go`: file that includes our `main()` function
- `server.go`: this is the transport of the application. In this file we will implement the http handlers for the server APIs.

## Step 0 - notice main.go
Let's have a look at `coolfacts_server/main.go`:
```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
	"github.com/FTBpro/go-workshop/coolfacts/inmem"
)

func main() {
	fmt.Println("Hello, Server!")

	factsRepo := inmem.NewFactsRepository()
	service := coolfact.NewService(factsRepo)
	server := NewServer(service)

	log.Println("starting server on port 9002")
	log.Println("You can go to http://127.0.0.1:9002/ping")
	if err := http.ListenAndServe(":9002", server); err != nil {
		panic(fmt.Errorf("server crashed! err: %w", err))
	}
}
```
We first can notice the imports `"fmt"`, `"log"` and `"net/http"` are standard Go libraries.

In the next import lines:
```go
"github.com/FTBpro/go-workshop/coolfacts/coolfact"
"github.com/FTBpro/go-workshop/coolfacts/inmem"
```
we are importing packages from our own module. We see the module path and then the path to the packaghe we want to import. The module path is `github.com/FTBpro/go-workshop/coolfacts`.

In the next lines we initializing the repo, service and the server. A pacage name is only the last param, and each type in Go has a name composed from the package name and the type identifier. For example, we call `inmem.NewFactsRepository()`. The package name is `inmem`, and the type identifier is `inmem.NewFactsRepository()` 

Finally, we use `http.ListenAndServe`. This method receives an address and an `http.Handler`
- address: This is the address it will listen on the TCP network
- `http.Handler`: For every incoming connection, it will use the `http.Handler` to handle the requests
  ```go
  package http
  
  type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
  }
    ```
  The handler should handle the incoming `http.Request` and write the response to the `ResponseWriter`. In out application the http.Handler will be `coolfact.service` initialized with `coolfact.NewService(...)` 

What you will have to complete is:
## Step 1 - package `coolfact`
This package handles the BL of the application.

### <u> file `coolfact/fact.go`: </u>
In here we have the entity of the application. A struct named `Fact`.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Complete the definition of `Fact`. It should have the following fields:
    - Image string
    - Description string

### <u> file `coolfact/service.go`:</u>
In service.go we have the service which will handle the "BL" for the application. In the service you will:
- **Finish `Repository` interface{}**
  - in Go we declare interfaces where we use them, not where we implement them. For the service to operate properly, it requires a `Repository` interface which he defines. It makes sense, since the service knows what it needs to do, and what the dependancy it needs to have. We can note that the name of the interface{} isn't InmemRepo or SQLRepo or something else, since the service is agnostic to the way the repo operates. He doesn't care about the mechanism, only behaviour.
  - <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Add one method for getting facts.
    - `GetFacts` - return slice of `coolfact.Fact` and an error
  - <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Finish definition of service.
  - <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `NewService`. Return instance of `service` initialized with its field.
  - <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement methods `GetFacts`.

## Step 2 - Package `inmem`
### <u> file `inmem/factsrepo.go`:</u>
Here we will implement the facts repository. Currently, only with functionality to return facts.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `NewFactsRepository`
  - Just so we will have initial data, initialize the repo with two facts.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement method `GetFacts`.

## Step 3 - `cmd/server.go`
The server is responsible to respond to the incoming HTTP request. It initialized with the service for handling the BL. Currently our only use-case is getting facts.
In addition, it will have ping route, for answering with PONG.

You can notice that the server has method `ServeHTTP`. This is the entry point for every incoming HTTP request. The server implement `http.Handler`. We have switch case according to the HTTP method and path.

The only recognized path is `GET /facts`. On every path it returns `HandleNotFound` which you will implement. In real world applications, we will usually use a framework wrapping Go http which will provide us with more advanced routing mechanism.

- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Finish definition of `Service` interface{}. 
  - What the server needs in order to complete the request for getting facts?
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Finish definition of `server` and the `NewServer` function.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `HandlePing` method.
  - Response with status 200 (const `http.StatusOK`)
  - Write `PONG` as the response.
  - Clue 1 - check the `http.ResponseWriter` interface. Notice its methods and decide what you can use.
  - Clue 2 - The `http.ResponseWriter` implements `io.Writer` (check why). Therefore, we can use a method in the Go package `fmt` to write the response
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `HandleGetFacts` method. The method for handling the `GET /facts` API
  - Call the service in order to get the facts.
  - Format the response to JSON:
  ```json
  {
    "facts": [
        {
          "image": "...",
          "description": "..."
        },
        //...
    ]
  }
  ```
  - Write status 200.
  - Set "content-type" header to "application/json".
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
In the following section you fill find a full walkthrough. Use it in case you are stuck

## Step 1 - Implement the core entity and the service

In `fact.go` we simply add fiedls to the entity:
```go
type Fact struct {
  Image       string
  Description string
}
```
We can notice that these are exported (public) fields, since it’s an entity and the rest of the application should be aware of.

#### <u> In `service.go` </u>
The `Repository` interface currently only should have one method for getting facts.
The `service` has one field for the `factsRepo`, notice that the service and the field are _private_, we don’t want that someone will set or get it from outside.
We want to require the consumer to use the initializer.
```go
type Repository interface {
	GetFacts() ([]Fact, error)
}

type service struct {
	factsRepo Repository
}

func NewService(factsRepo Repository) *service {
	return &service{
		factsRepo: factsRepo,
	}
}
```

And the `GetFacts` implementation:
```go
func (s *service) GetFacts() ([]Fact, error) {
  facts, err := s.factsRepo.GetFacts()
  if err != nil {
     return nil, fmt.Errorf("factsService.GetFacts: %w", err)
  }

  return facts, nil
}
```
In GetFacts the service calls the repo. Notice that if there is an error, the service wraps it and adding some context, so we will have friendlier message.

## Step 2 - The repo
We initialize the `factsRepo` with a slice including 2 cool facts
```go
func NewFactsRepository() *factsRepo {
  return &factsRepo{
     facts: []coolfact.Fact{
        {
           Image:       "https://images2.minutemediacdn.com/image/upload/v1556645500/shape/cover/entertainment/D5aliXvWsAEcYoK-fe997566220c082b98030508e654948e.jpg",
           Description: "Did you know sonic is a hedgehog?!",
        },
        {
           Image:       "https://images2.minutemediacdn.com/image/upload/v1556641470/shape/cover/entertainment/uncropped-Screen-Shot-2019-04-30-at-122411-PM-3b804f143c543dfab4b75c81833bed1b.jpg",
           Description: "You won't believe what happened to Arya!",  
        },
     },
  }
}

func (r *factsRepo) GetFacts() ([]coolfact.Fact, error) {
  return r.facts, nil
}
```

## Step 3 - The HTTP transport layer
The `server` receives the `Service interface{}` which implements the BL.
```go
type Service interface {
  GetFacts() ([]coolfact.Fact, error)
}

type server struct {
  factsService Service
}

func NewServer(service Service) *server {
  return &server{
     factsService: service,
  }
}
```
Notice the type `coolfact.Fact` which is returned in the `Service` method.
A common anti-pattern in Go is to repeat a word in a type name. For example `http.Handler` and not `http.HttpHandler`.
In case of entities we sometimes encounter such repetition.

`HandlePing` method:
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

`HandleGetFacts` method:

First, we receive the slice of facts from the service and format them to a json which will be returned to the client.
```go
func (s *server) HandleGetFacts(w http.ResponseWriter) {
  facts, err := s.factsService.GetFacts()
  if err != nil {
     s.HandleError(w, fmt.Errorf("server.GetFactsHandler: %w", err))
  }

  // we first format the facts to map[string]interface.
  formattedFacts := make([]map[string]interface{}, len(facts))
  for i, coolFact := range facts {
     formattedFacts[i] = map[string]interface{}{
        "image":       coolFact.Image,
        "description": coolFact.Description,
     }
  }

  response := map[string]interface{}{
     "facts": formattedFacts,
  }
  
  // code omitted
}
```
You may ask yourself why don’t we just return the facts to the client?
This is because we want to keep separation of concerns. The entity should know nothing about the outer layer, for example the client.

The API response and the entity field should be considered different. We don’t want that changes in the entity will have unexpected cascading changes,  
and we don’t want that a requirement to modify the response will trigger a change to the entity.

The rest of the method is to write the status and content type header, and then write the response body to the writer:
```go
func (s *server) HandleGetFacts(w http.ResponseWriter) {
  // code omitted
	
  // write status and content-type
  // status must be written before the body
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")

  // write the body. We use json encoding
  if err := json.NewEncoder(w).Encode(response); err != nil {
  fmt.Printf("HandleGetFacts ERROR writing response: %s", err)
}
```

Since we are returning a JSON, we can use the Go `json` package, which can writes a JSON encoding of the response to the writer.
The `NewEncoder` method receives `io.Writer`. We already can notice the power of Go interfaces, and especially interfaces with one method.
We can send the `http.ResponseWriter` to different packages methods, and use a different encoding, and the packages are totally agnostic to HTTP.

All we left is to implement `HandleNotFound` and `HandleError`:
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
Congratulation! You've just implemented a web server and BL for returning facts.

In the following exercise we will add an API for creating a fact, and a client application for calling our server.
