# Part 3

In this exercise, you will add an API to the server for presenting the existing facts.

## The Starting Point

You will need to switch to branch `v3-server-get-facts` for this part. Note that if you have changed in your code from previous exercise, you should clean them. You can use `git reset --hard` for that.
Go on and navigate to `go-workshop/coolfacts` directory. You will see a bunch of files and TODOs. We will go over them in the rest of this document.

## Your Goal

After implementing all the TODOs, the server will export an API for getting the current facts.
```json
GET "/facts"

response: 
{
  "facts": [
    {
        "topic": "...",
        "description": "..."
    }
	//...
  ]
}
```

## Getting Started
Take a look around the program. You will notice new folders - `coolfact` and `inmem`. In addition, you will notice some changes in the main package.

### coolfact
This is a package containing the entity `fact` and the service for implementing the use case (business logic).
In this application we won't have much BL, if any. Our application is a very simple CRUD app, and the service will mainly call the repo as we will see.

### inmem
In this package we will implement the facts-repository. We will use in memory.
In Go, packages names are very important. A package name should say what it provides, not what it contains.
We won't have packages names like `repos`, `models`, `utils`, `common`, etc.

A package’s name provides context for its contents, making it easier for clients to understand what the package is for and how to use it.
For example, Go package `time` indicates that it contains functionality for handling times, probably a struct and behaviour for measuring and displaying time. Go package `http` lets you speak `HTTP`.

In our case, package `inmem` imply that it uses memory mechanism. If we have SQL database, we will have package `sql` that implies it lets you speak SQL.

## Step 0 - notice changes in main.go
Let's have a look at `coolfacts_server/main.go`:
```go
package main

import (
	"fmt"
	"log"
	"net/http"

	// new imports
	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
	"github.com/FTBpro/go-workshop/coolfacts/inmem"
)

func main() {
	fmt.Println("Hello, Server!")

	// new initializations
	factsRepo := inmem.NewFactsRepository(seedFacts()...)
	service := coolfact.NewService(factsRepo)
	server := NewServer(service)

	log.Println("starting server on port 9002")
	log.Println("You can go to http://127.0.0.1:9002/ping")
	if err := http.ListenAndServe(":9002", server); err != nil {
		panic(fmt.Errorf("server crashed! err: %w", err))
	}
}

func seedFacts() []coolfact.Fact {
  return []coolfact.Fact{
    {
      Topic:       "Games",
      Description: "Did you know sonic is a hedgehog?!",
    },
    {
      Topic:       "TV",
      Description: "You won't believe what happened to Arya!",
    },
  }
}
```
We first can notice new imports:
```go
"github.com/FTBpro/go-workshop/coolfacts/coolfact"
"github.com/FTBpro/go-workshop/coolfacts/inmem"
```
we are importing packages from our own module. We have the module path and then the path to the package we want to import. The module path is `github.com/FTBpro/go-workshop/coolfacts`.

In the next lines we're initializing the repo, service and the server. A package name is only the last param, and each type in Go has a name composed of the package name and the type identifier. For example, we call `inmem.NewFactsRepository`. The package name is `inmem`, and the full type name is `inmem.NewFactsRepository()`.

You can note that we are initializing the repo with facts: `inmem.NewFactsRepository(seedFacts()...)`. The `inmem.NewFactsRepository` signature is:
```go
func NewFactsRepository(facts ...coolfact.Fact) *factsRepo {
```
The 3 dots indicates that it is a _variadic_ function. It can be called with any number of trailing arguments. You already used a variadic function in the previous exercises, `fmt.Println`:
```go
func Println(a ...any) (n int, err error) {
```

If you already have multiple args in a slice, you apply them to a variadic function using `func(slice...)`. This is why we call the function like this:
```go
inmem.NewFactsRepository(seedFacts()...)
```


What you will have to complete is:
## Step 1 - package `coolfact`
This package handles the BL of the application.

### <u> file `coolfact/fact.go`: </u>
In here we have the entity of the application. A struct named `Fact`.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Complete the definition of `Fact`. It should have the following fields:
  - Topic - string
  - Description - string

### <u> file `coolfact/service.go`:</u>
In service.go we have the service which will handle the "BL" for the application. In the service you will:
- **Finish `Repository` interface{}**
  - in Go we declare interfaces where we use them, not where we implement them. For the service to operate properly, it requires a `Repository` interface which he defines. It makes sense, since the service knows what it needs to do, and what the dependancy it needs to have. We can note that the name of the `interface{}` isn't `InmemRepo` or `SQLRepo` or something else, since the service is agnostic to the way the repo operates. He doesn't care about the mechanism, only behaviour.
  - <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Add one method for getting facts.
    - `GetFacts` - return slice of `coolfact.Fact` and an error
  - <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Finish definition of service. The service should have a field for the `factsRepo`, which it will receive in the initialization. 
  - <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `NewService`. Return instance of `service` initialized with its field.
  - <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement methods `GetFacts`.

## Step 2 - Package `inmem`
### <u> file `inmem/factsrepo.go`:</u>
Here we will implement the facts repository. Currently, only with functionality to return facts.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `NewFactsRepository`.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement method `GetFacts`.

## Step 3 - `cmd/server.go`
As mentioned before, you will implement a new API for the `server`, what you will have to complete in the server is:
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Finish definition of the `FactsService` interface. Which method does the server needs in order to operate?
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Add field in the `server` for the factsService, and complete the initialization. In function `NewServer` add pass argument. 
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">In method `serveHTTP`, add a case for the new API, that will call method `HandleGetFacts`
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `HandleGetFacts` method. The method for handling the `GET /facts` API
  - Call the service in order to get the facts.
  - Format the response to JSON:
  ```json
  {
    "facts": [
        {
          "topic": "...",
          "description": "..."
        },
        //...
    ]
  }
  ```
  - Write status 200.
  - Set "content-type" header to "application/json".

### Building and Running

If everything is implemented well, this is what the final result should look like when running the program:
![factsgif](https://user-images.githubusercontent.com/5252381/204143457-6eaf59d3-6c52-4fbb-8d2a-19d22436cbd8.gif)

# Full walkthrough
In the following section you fill find a full walkthrough. Use it in case you are stuck.

## Step 1 - Implement the core entity and the service

In `fact.go` we simply add fiedls to the entity:
```go
type Fact struct {
  Topic       string
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
In `GetFacts`, the service calls the repo. Notice that if there is an error, the service wraps it and adding some context, so we will have friendlier message.

## Step 2 - The repo
We initialize the `factsRepo` with the arg slice:
```go
func NewFactsRepository(facts ...coolfact.Fact) *factsRepo {
  return &factsRepo{
     facts: facts,
  }
}

func (r *factsRepo) GetFacts() ([]coolfact.Fact, error) {
  return r.facts, nil
}
```

## Step 3 - The HTTP transport layer
The `server` receives the `Service interface{}` which implements the BL. We're injecting `Service` interface as a dependency to the `server`, for the `server` to operate. This is what the `server` requires.
```go
type FactsService interface {
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
Notice the type `coolfact.Fact`. A common anti-pattern in Go is to repeat a word in a type name. For example `http.Handler` and not `http.HttpHandler`.
In case of entities we sometimes encounter such repetition.

`HandleGetFacts` method:

First, we receive the slice of facts from the service and then format them to a JSON which will be returned to the client.
```go
func (s *server) HandleGetFacts(w http.ResponseWriter) {
  facts, err := s.factsService.GetFacts()
  if err != nil {
     s.HandleError(w, fmt.Errorf("server.GetFactsHandler: %w", err))
  }

  response := s.formatGetFactsResponse(facts)
  
  // code omitted
}

func (s *server) formatGetFactsResponse(facts []coolfact.Fact) map[string]interface{} {
	formattedFacts := make([]map[string]interface{}, len(facts))
	for i, coolFact := range facts {
		formattedFacts[i] = map[string]interface{}{
			"topic":       coolFact.Topic,
			"description": coolFact.Description,
		}
	}
	
	return map[string]interface{}{
		"facts": formattedFacts,
	}
}
```
You may ask yourself why don’t we just return the facts to the client?
This is because we want to keep separation of concerns. The entity should know nothing about the outer layer, for example the http response.

The API response and the entity field should be considered different. We don’t want that changes in the entity will have unexpected cascading changes,  
and we don’t want that a requirement to modify the response will trigger a change to the entity.

NExt, we'll write the status and content type header, and then write the response body to the writer. `w.WriteHeader` must be called before `w.Write`.
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
The `NewEncoder` method receives `io.Writer`. We can pass the `http.ResponseWriter` to different packages methods, and use a different encoding, and the packages are totally agnostic to HTTP.

# Finish!
Congratulation! You've just implemented a new API with a totally cool use-case!
