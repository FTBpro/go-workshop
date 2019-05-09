# welcome

go workshop - steps
***
## Exercise 0 - build and run
### Goal 
Build and run
### Steps
Write a go program which prints "hello world"
***
## Exercise 1 - Ping
### Goal
First use of `http` package

### Steps
* Register `"/ping"` to an `http.HandleFunc` with a function that writes to ResponseWriter the ‘PONG’
* For printing into `w http.ResponseWriter` you can use `fmt.Fprint`
* In case of an error you can use `http.Error`

### End result
When navigating to `localhost:9002/ping` the browser should show `PONG`\

***
## Exercise 2 - list the facts as JSON
### Goal
Create `/facts` endpoint for listing facts in a JSON format

### Steps

#### Create fact struct
The `fact` struct should have 3 string fields : `Image`, `Url`, `Description`

#### Create store struct
The store can use in memory cache for storing the facts.\
It can be done by one field `facts` of type `[]fact` (a slice of facts).

Add store functionality:
* `func (s store) getAll() []fact {…}`
  * The method should return all facts in the store.facts field
* `func (s store) add(f fact) {…}`
  * The method should add the given fact f to the store\
  For adding to a slice you can use `store.facts = append(store.facts, f)`
  
Init the store from `main` with some static data. 

#### Register `/facts` to to `http.HandleFunc` 
Register "/facts" to http.HandleFunc with a function that:
* Gets all the facts from the store
* Writes the facts to the `ResponseWriter` in JSON format
  * Use `json.Marshal` to format the struct as json and to write to the `ResponseWriter` 

### End result
`GET /facts` will return json of all facts in store

***
## Exercise 3 - create new fact
### Goal
Create POST request for creating a new fact

### Steps

#### Add `POST` functionality to `/facts` handlerFunc
In the handler from the previous exercise check for the request method (GET/POST) add the logic of this step under POST section

Create a json format equivalent in fields (types and names) to the fact struct.
 
```go
var req struct {
    Image       string `json:"image"`
    Url         string `json:"url"`
    Description string `json:"description"`
}
```

#### Parse the encoded request body into `req`
Reading the request body can be done by `ioutil.ReadAll`\
Parsing the data into `req` can be done by `json.Unmarshal` 

#### Finally add the fact to the store\
Use the `factsStore.add` method from Exercise 2.

### End result
`POST /facts` will create a new fact and add it to store

***
## Exercise 4 - list the facts as HTML

### Goal
List the index results you created in exercise 2, using HTML template
### Steps
1. Crate an HTML template using package `text/template` syntax
2. Execute template with `store.getAll` results (that means write to `ResponseWriter` all results in the applied template)
### End result
return the index results (GET /facts) with an HTML template

***

## Exercise 5 - use MentalFloss API

### Goal
Replace the static data with data of external provider (mentalfloss)

### Steps

#### Create a mentalfloss struct 
In a seperate file, create mentalfloss struct which will be the provider for fetching the facts

#### Add `func (mf mentalfloss) Facts() ([]fact, error) {…}`
Function for fetching facts using MetnalFloss API\
* Call `http://mentalfloss.com/api/facts` using `http.Get`
  * This return array of MentalFloss facts in a JSON format
* Parse the response body into a custom struct like in exercise 3 using `json.Unmarshal`
  * The custom struct you can use in here is an array of fact equivalent struct:
  ```go
  var items []struct {
      Url          string `json:"url"`
      FactText     string `json:"fact"`
      PrimaryImage string `json:"primaryImage"`
  }
  ```
  * For convenience you may use private func for parsing the response data and converting to `[]fact` 
    * `func parseFromRawItems(b []byte) ([]fact, error) {…}`
    * (After you read the response body by `ioutil.ReadAll`)   

#### Use mentalfloss instead of static data
In main, get facts from `mentalfloss`, and add these facts to the store.

### End result
`GET /facts` should show facts from MentalFloss.\
(By sending request to external provider (MentalFloss) to fetch facts and saving them in the store)

***

## Exercise 6 - seperate to packages

### Goal
Separate structs into separate packages
### Steps

#### Create package `fact`
Create a new folder fact - move store and fact definition into that folder (change the package name to fact)

Make sure that the struct `Fact` is exported (public)  
#### Create package `mentalfloss`
Create a new folder `mentalfloss` - move mentalfloss struct and methods into that folder (change the package name to mentalfloss)

> You will encounter compile error since the `fact` is now in another package.\
You will need to import your fact package And replace `fact` with `fact.Fact`\
(for example in v6 `import "github.com/FTBpro/go-workshop/coolfacts/v6/fact"`)


#### Create package `http`
The goal is to separate http.HandlerFunc logic outside of main\

Create a new folder `http` and `handler.go` file (can be other file name)\
Create handler struct for handling the request.\
Example 
```go
type FactsHandler struct {
	FactStore *fact.Store
}
```

> Optional: you can declare `type FactStore interface` and use it instead of directly use `*fact.Store`. (Will be used only in exercise 8)

Create methods for handling the request
* `func (h *FactsHandler) Ping(w http.ResponseWriter, r *http.Request)`
* `func (h *FactsHandler) Facts(w http.ResponseWriter, r *http.Request)`

> Move the anonymous `http.HandleFunc` from main and put as this struct's methods (these with the signature `func(w http.ResponseWriter, r *http.Request)`) 


In main, init `FactsHandler` struct with the `factStore`\
for example:
```go
handlerer := facthttp.FactsHandler{
	FactStore: &factsStore,
}
```

Since we already importing `net/http` in main, we need to rename the name we will use for our own http 
```go
import (
	"net/http"
	facthttp "github.com/FTBpro/go-workshop/coolfacts/v8/http"
)
```
Use this handlerer methods for registering to the entpoints, for example:
```go
http.HandleFunc("/ping", handlerer.Ping)
```

***

## Exercise 7 - add ticker for updating the facts

### Goal
Use go channel and ticker for updating the fact inventory
### Steps
1. Init a context.WithCancel (remember to defer its closer…)
2. Add a function - func updateFactsWithTicker(ctx context.Context, updateFunc func() error)
    1. (Outside from updateFactsWithTicker) Create the updateFunc from step 7.2. that updates the store from an external              provider
    2. (Within the updateFactsWithTicker) Create a time.NewTicker 
    3. (Within the updateFactsWithTicker) Add a go routine with a function that accepts the context
        1. Inside the function add an endless loop that will select from the channel (the ticker channel and the context                  one)
            1. If the ticker channel (ticker.C) was selected - use the given updateFunc to update store
            2. If the context channel (context.Done()) was selected -return (it means the main closed the context)
            
### End result
every const time a ticker will send a signal to a `thread` (go built-in) that will fetch new fact from provider (mentalfloss)
*** 

## Exercise 8 - refactor

### Goal
* Enable switching between persistent layers easily
* Enable replacing and adding new providers
### Steps

### Extract store logic to package `inmem`
Create `inmem` package and Move the store currently in `fact` package
    
> `package inmem` is a common name for handling in memory cache. For handling cache in sql for example, you can use `package sql`

### Create fact service for abstracting how we update the facts
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

### _Replace calls in main_
Use `inmem.FactStore{}` for the fact store (instead of `fact.Store{}`)

Init the service with the mentalfloss provider and the inmem store
Example
```go
service := fact.NewService(&factsStore, &mf)
```

Use the `service.UpdateFacts` method to update the store, and to use with the ticker

### End result

We now can add another provider and cache layer easily