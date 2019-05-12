# welcome
This is a step-by-step tutorial for creating a simple server for fetching facts from MentalFloss API, and list them in a simple HTML template

* [Entry Point - Hello World](#exercise-0---hello-world)
* [Exercise 1 - ping](#exercise-1---ping)
* [Exercise 2 - list facts as JSON](#exercise-2---list-the-facts-as-json)
* [Exercise 3 - create new fact](#exercise-3---create-new-fact)
* [Exercise 4 - list the facts as HTML](#exercise-4---list-the-facts-as-html)
* [Exercise 5 - use MentalFloss API](#exercise-5---use-mentalfloss-api)
* [Exercise 6 - separate to packages](#exercise-6---separate-to-packages)
* [Exercise 7 - add ticker for updating the facts](#exercise-7---add-ticker-for-updating-the-facts)
* [Exercise 8 - refactor](#exercise-8---refactor)

***
## Entry Point - Hello World
### Goal 
Build and run

### Steps
##### Install go
For install go and editor, see [here](https://github.com/FTBpro/go-workshop/blob/master/INSTALL_GO.md)

##### Clone the project

In your favourite terminal, clone the project into your favourite folder

```bash
/your/favourite/folder$ git clone https://github.com/FTBpro/go-workshop.git
```

##### Run the project
Go into the `entrypoint` folder:
```
/your/favourite/folder$ cd go-workshop/coolfacts/entrypoint/
```

From there, run the project:

```
/your/favourite/folder/go-workshop/coolfacts/entrypoint$ go run main.go
```

Now you should see some welcome string in your terminal.\
Be sure you know where the code that prints this line is coming from. (hint: you can finc it in `entrypoint/main.go`)

For more details on build and run, you can checkout [this readme](https://github.com/FTBpro/go-workshop/blob/master/coolfacts/exercise0/build-and-run.md)

##### Next exercises:
For all farther exercises you can continue to write the code in this folder, (in `main.go` and later in other files)

At any time if you are having any issues, you can use the reference for the exercise implementation under `/exerciseN/...`

***
## Exercise 1 - Ping
### Goal
First use of `http` package with a simple server

### End result
When navigating to `localhost:9002/ping` the browser should show `PONG` string

### Steps

##### Create `/ping` endpoint
Register handler function of to `/ping` pattern\
Use `http.HandleFunc` function to register an anonymous function of type `func(http.ResponseWriter, *http.Request)` to `/ping`

> For printing into `http.ResponseWriter` you can use `fmt.Fprintf`\
In case of an error you can use `http.Error` function

##### Listen on port :9002
You can use `http.ListenAndServe` for starting the server

***
## Exercise 2 - list the facts as JSON
### Goal
Create `/facts` endpoint for listing facts in a JSON format by using a static store

### End result
`GET /facts` will return JSON of all facts in store, for now it will be hard coded facts

### Steps

##### Create fact struct
Create a struct named fact (`type Fact struct {...}`)\
The `fact` struct should have 3 string fields : `Image`, `Url`, `Description`

##### Create store struct
Create a struct named store (`type Store struct {...}`)\
The store can use in memory cache for storing the facts.\
It can be done by one field `facts` of type `[]fact` (a slice of facts).

Add store functionality:
* `func (s *store) getAll() []fact {…}`
  * The method should return all facts in the `store.facts` field
* `func (s *store) add(f fact) {…}`
  * The method should add the given fact f to the store\
  For adding to a slice you can use `store.facts = append(store.facts, f)`
  
Init the store from `main` with some static data. 

##### Register `http.HandleFunc` to `/facts` pattern  
Like in the ping from previous exercise, use an anonymous function as argument to `http.HandleFunc` function.

In this function you will:
* Get all the facts from the store
* Write the facts to the `ResponseWriter` in a JSON format
> Use `json.Marshal` to format the struct as json and to write to the `ResponseWriter` 

***
## Exercise 3 - create new fact
### Goal
Create POST request for creating a new fact

### End result
`POST /facts` will create a new fact and add it to store

### Steps

##### Add `POST` functionality to `/facts` handlerFunc
In the handler from the previous exercise check for the request method (GET/POST) add the logic of this step under POST section

Create a json format equivalent in fields (types and names) to the fact struct.

```go
var req struct {
    Image       string `json:"image"`
    Url         string `json:"url"`
    Description string `json:"description"`
}
```

##### Parse the encoded request body into `req`
Reading the request body can be done by `ioutil.ReadAll`\
Parsing the data into `req` can be done by `json.Unmarshal` 

##### Finally add the fact to the store
Use the `factsStore.add` method from Exercise 2.

***
## Exercise 4 - list the facts as HTML

### Goal
* Using HTML template
* Replace the JSON representation from exercise 2 with an HTML

### End result
`GET /facts` will list the facts in HTML

### Steps

##### Create an HTML template
This can be done using package `text/template` syntax

##### Execute the template with facts
Execute template with `store.getAll` results (that means write to `ResponseWriter` all results in the applied template)

***

## Exercise 5 - use MentalFloss API

### Goal
Use MetnalFloss API for fetching the facts and initialize the store, instead of the static data

### End result
`GET /facts` should show facts from MentalFloss.\
> This will be done by sending request to the external provider (MentalFloss) to fetch facts and saving them in the store\
You can use this API for fetching the data: ``http://mentalfloss.com/api/facts``

### Steps

##### Create a mentalfloss struct 
Create `mentalfloss` struct which will act as the provider for fetching the facts
> For convenience you can do this in a separate file, but still in package `main`  

Add `metnalfloss` functionality
* `func (mf mentalfloss) Facts() ([]fact, error) {…}`
  * Function for fetching facts using MetnalFloss API
  * Call `http://mentalfloss.com/api/facts` using `http.Get`, this returns an array of MentalFloss facts in a JSON format
  * Parse the response body into a custom struct like in exercise 3 using `json.Unmarshal`
  * The custom struct you can use in here is an array of fact equivalent struct:
  ```go
  var items []struct {
      Url          string `json:"url"`
      FactText     string `json:"fact"`
      PrimaryImage string `json:"primaryImage"`
  }
  ```
  * For convenience you may use a private func for parsing the response data and converting to `[]fact` 
    * `func parseFromRawItems(b []byte) ([]fact, error) {…}`
      > After you read the response body by `ioutil.ReadAll`   

##### Use `mentalfloss` instead of static data
In main, get facts from `mentalfloss`, and add these facts to the store.

***

## Exercise 6 - separate to packages

### Goal
Separate structs into separate packages

### Steps

##### Create package `fact`
Create a new folder fact - move store and fact definition into that folder (change the package name to fact)

Make sure that the struct `Fact` is exported (public)
  
##### Create package `mentalfloss`
Create a new folder `mentalfloss` - move mentalfloss struct and methods into that folder (change the package name to mentalfloss)

> You will encounter compile error since the `fact` is now in another package.\
You will need to import your fact package And replace `fact` with `fact.Fact`\
(for example in exercise 6 `import "github.com/FTBpro/go-workshop/coolfacts/exercise6/fact"`)


##### Create package `http`
The goal is to separate `http.HandlerFunc` logic outside of main

Create a new folder `http` and `handler.go` file (can be other file name)\
Create handler struct for handling the request.\
Example 
```go
type FactsHandler struct {
	FactStore *fact.Store
}
```

Create methods for handling the request
* `func (h *FactsHandler) Ping(w http.ResponseWriter, r *http.Request)`
* `func (h *FactsHandler) Facts(w http.ResponseWriter, r *http.Request)`

> Move the anonymous `http.HandleFunc` from main and put as this struct's methods (these with the signature `func(w http.ResponseWriter, r *http.Request)`) 


In main, init `FactsHandler` struct with the `factStore`\
Since we already importing `net/http` in main, we need to rename the name we will use for our own http ([How to import and use different packages of the same name](https://stackoverflow.com/questions/10408646/how-to-import-and-use-different-packages-of-the-same-name-in-go-language))\
For example 
```go
import (
	"net/http"
	facthttp "github.com/FTBpro/go-workshop/coolfacts/exercise8/http"
)
```
```go
handlerer := facthttp.FactsHandler{
	FactStore: &factsStore,
}
```
Use this handlerer methods for registering to the entpoints\
For example:
```go
http.HandleFunc("/ping", handlerer.Ping)
```

***

## Exercise 7 - add ticker for updating the facts

### Goal
Use go channel and ticker for updating the fact inventory

### End result
Every const time a ticker will send a signal to a `thread` (go built-in) that will fetch new fact from provider (mentalfloss)

### Steps
1. Init a context.WithCancel (remember to defer its closer…)
2. Add a function - func updateFactsWithTicker(ctx context.Context, updateFunc func() error)
    1. (Outside from updateFactsWithTicker) Create the updateFunc from step 7.2. that updates the store from an external              provider
    2. (Within the updateFactsWithTicker) Create a time.NewTicker 
    3. (Within the updateFactsWithTicker) Add a go routine with a function that accepts the context
        1. Inside the function add an endless loop that will select from the channel (the ticker channel and the context                  one)
            1. If the ticker channel (ticker.C) was selected - use the given updateFunc to update store
            2. If the context channel (context.Done()) was selected -return (it means the main closed the context)
            
*** 

## Exercise 8 - refactor

### Goal
* Enable switching between persistent layers easily
* Enable replacing and adding new providers
### Steps

##### Extract store logic to package `inmem`
Create `inmem` package and Move the store currently in `fact` package
    
> `package inmem` is a common name for handling in memory cache. For handling cache in sql for example, you can use `package sql`

##### Create fact service for abstracting how we update the facts
in package `fact`, declare two interfaces:
```go
type provider interface {
	Facts() ([]Fact, error)
}

type store interface {
	Add(f Fact)
	GetAll() []Fact
}
```

Create service struct which will have fields for these as inverted dependencies.\
For example
```go
type service struct {
	provider Provider
	store    Store
}
```

Add initializer - `func NewService(s Store, r Provider) *service`  \
And a method for updating facts `func (s *service) UpdateFacts() error`
> in `UpdateFacts` you will use the provider to fetch the facts, and the store to save them. Although the service is updating, it doesn't know who is the provider, and what is the persistent layer

##### Add abstraction in local `http` package
In package 'http' declare the FactStore interface:
```go
type FactStore interface {
	Add(f Fact)
	GetAll() []Fact
}
```

This will be used as a field to `FactsHandler`:
```go
type FactsHandler struct {
	FactStore FactStore
}
```
> The FactStore provider is replacing the field type from exercise 7 which is a concrete struct `fact.Store`. This enables us to add another layer for persistent data


##### Replace calls in main
Use `inmem.FactStore{}` for the fact store (instead of `fact.Store{}`)

Init the service with the mentalfloss provider and the inmem store
Example
```go
service := fact.NewService(&factsStore, &mf)
```

Use the `service.UpdateFacts` method to update the store, and to use with the ticker

### End result

We now can add another provider and cache layer easily
