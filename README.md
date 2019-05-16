# welcome
This is a step-by-step tutorial for creating a simple web server for fetching facts from MentalFloss API, and list them in a simple HTML template

## How to use this repo
The entrypoint described below is intended for setting up your environment and placing you in a ready-to-go folder you can start your project from.

Each exercise continues the previous one by adding or changing functionality. 

If you are encountering issues you can use the steps defined in each exercise for more detailed wolkthrough.\
Furthermore, you can find the implementation for each exercise under folder `coolfacts/exerciseN/...`

Also, you can get a better perspective by running each exercise by

```
cd /path/to/go-workshop/coolfacts
go run ./exerciseN
```

Hope you will have fun and good luck :) <img src="https://github.com/egonelbre/gophers/blob/master/vector/adventure/hiking.svg" width="48">


* [Entrypoint - Hello World](#entrypoint---hello-world-)
  * [Installation and Editor guide](https://github.com/FTBpro/go-workshop/blob/master/INSTALL_GO.md)
  * [Build and Run guide](https://github.com/FTBpro/go-workshop/blob/master/coolfacts/entrypoint/build-and-run.md)
* [Exercise 1 - ping](#exercise-1---ping-)
* [Exercise 2 - list facts as JSON](#exercise-2---list-the-facts-as-json-)
* [Exercise 3 - create a new fact](#exercise-3---create-a-new-fact-)
* [Exercise 4 - list the facts as HTML](#exercise-4---list-the-facts-as-html-)
* [Exercise 5 - use MentalFloss API](#exercise-5---use-mentalfloss-api-)
* [Exercise 6 - separate to packages](#exercise-6---separate-to-packages-)
* [Exercise 7 - add ticker for updating the facts](#exercise-7---add-ticker-for-updating-the-facts-)
* [Exercise 8 - refactor](#exercise-8---refactor-)

***
> By the way, all the gophers images are taken from the wonderfull https://github.com/egonelbre/gophers
## Entrypoint - Hello World <img src="https://github.com/egonelbre/gophers/blob/master/vector/fairy-tale/witch-learning.svg" width="55">

### Goal 

Build and run

### Steps

##### Installation (go + editor)
For install go and editor, see [here](https://github.com/FTBpro/go-workshop/blob/master/INSTALL_GO.md)

##### Clone the project

Clone the project in your favourite terminal, 

```
git clone https://github.com/FTBpro/go-workshop.git
```

##### Run the project

cd into the `entrypoint` folder and run the entrypoint project:

```
cd go-workshop/coolfacts/entrypoint/
go run .
```

For more details on build and run, you can checkout [this readme](https://github.com/FTBpro/go-workshop/blob/master/coolfacts/entrypoint/build-and-run.md)

### End Result

If everything was successfull you should see a lovely welcome msg in your terminal :smile_cat:\
Be sure you know where the code which prints this line is coming from. (hint: you can find it in `entrypoint/main.go`)

For more details on build and run, you can checkout [this readme](https://github.com/FTBpro/go-workshop/blob/master/coolfacts/entrypoint/build-and-run.md)

##### Next exercises:
For all further exercises you can continue to write the code in this folder, (in `main.go` and later in other files)

At any time if you are having any issues, you can use the reference for the exercise implementation under `/exerciseN/...`

***

## Exercise 1 - Ping <img src="https://github.com/egonelbre/gophers/blob/master/vector/projects/go-grpc-web.svg" width="45">

### Goal

First use of `http` package with a simple server

### End result

When navigating to `http://localhost:9002/ping` the browser should show `PONG` string

### Steps

##### Create `/ping` endpoint

From main function, you will need to register to `/ping` pattern.

You can use [`http.HandleFunc`](https://golang.org/pkg/net/http/#HandleFunc) for doing that in a simple way.

For example:

```go
http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// place your code here
}
```

> For printing into `http.ResponseWriter` you can use `fmt.Fprintf`, and in case of an error you can use `http.Error` function


##### Listen on port :9002
Next you will need to have a server listening on port :9002 to get the ping.

We will use the default server in the http package using `http.ListenAndServe`.\
For example:
```go
http.ListenAndServe(":9002", nil)
```
 
***

## Exercise 2 - list the facts as JSON <img src="https://github.com/egonelbre/gophers/blob/master/vector/projects/wwgl.svg" width="45">

### Goal

Create `/facts` endpoint for listing facts in a JSON format by using a store.

### End result

`http://localhost:9002/facts` will show a JSON of all facts in store, for now it will be hard coded facts.

### Steps

##### Create fact struct

Create a struct named fact (`type Fact struct {...}`)\
The `fact` struct should have 2 string fields : `Image`, `Description`

##### Create store struct

Create a struct named store (`type Store struct {...}`)\
The store can use in memory cache for storing the facts, it can be done by one field `facts` of type `[]fact` (a slice of facts).

Add store functionality:
* `func (s *store) getAll() []fact {…}`
  * The method should return all facts in the `store.facts` field
* `func (s *store) add(f fact) {…}`
  * The method should add the given fact f to the store\
  For adding to a slice you can use `store.facts = append(store.facts, f)`
  
Init the store from `main` with some static data. 

##### Register `http.HandleFunc` to `/facts` pattern  

Like in the ping from previous exercise, use an anonymous function as an argument to `http.HandleFunc` function.

In this function you will:
* Get all the facts from the store
* Write the facts to the `ResponseWriter` in a JSON format
> Use `json.Marshal` to format the struct as json and to write to the `ResponseWriter` 

***

## Exercise 3 - create a new fact <img src="https://github.com/egonelbre/gophers/blob/master/vector/fairy-tale/sage.svg" width="55">

### Goal

Create a new fact by a POST request.

### End result

Create a new fact and add it to the store by issuing a `POST /facts` request with the next payloiad:
```json
{
  "image": "image/url",
  "description": "image description"
}
```
> For issuing a POST request you can use the next command from terminal while your server is running:\
curl --header "Content-Type: application/json" --request POST --data '{"image":"\<insertImageURL>", "description": "\<insertDescription>"}' http://localhost:9002/facts

### Steps

##### Add `POST` functionality to `/facts` handlerFunc

In the handler from the previous exercise check for the request method (GET/POST) and add the logic of this exercise under POST section

##### Parse the request body into a `fact`

First, read the request body into a byte stream using `ioutil.ReadAll`:

```go
b, err := ioutil.ReadAll(r.Body)
```

Next, we need to parse this data into some sort of a "request model". We which use a struct, which should be a representation of the request payload.

In this exercise, the expected request payload is:

```json
{
  "image": "image/url",
  "description": "image description"
}
```
So our request model struct can be something like this:

```go
var req struct {
    Image       string `json:"image"`
    Description string `json:"description"`
}
```

Now we need to parse the data into this struct, for this we can use `json.Unmarshal`: 
```go
err = json.Unmarshal(b, &req)
```

Finally, after we have this struct filled, create a new fact from it, and add it to the store.

For adding it to the store you should use the `factStore.Add` from exercise 2.

***

## Exercise 4 - list the facts as HTML <img src="https://github.com/egonelbre/gophers/blob/master/vector/science/power-to-the-linux.svg" width="95">

### Goal

* Using HTML template
* Replace the JSON representation from exercise 2 with an HTML

### End result

`GET /facts` will list the facts in HTML
`http://localhost:9002/facts` will show a all facts in store in HTML.

### Steps

##### Create an HTML template

We will use `html/template` package.

For a very basic use, you can declare this variable outside of main

```go
var newsTemplate = `<!DOCTYPE html>
<html>
  <head><style>/* copy coolfacts/styles.css for some color 🎨*/</style></head>
  <body>
  <h1>Facts List</h1>
  <div>
    {{ range . }}
       <article>
            <h3>{{.Description}}</h3>
            <img src="{{.Image}}" width="100%" />
       </article>
    {{ end }}
  <div>
  </body>
</html>`
```

> This is called a 'package defined variable' (global variable), We can use it from anywhere in the package it defined in

##### Return HTML representation of the facts

This step will replace the JSON you added in exercise 2, so you will need to replace the code in the section under GET method in the handler for `/facts`.

Using `html/template`, create a new template from the `newsTemplate` defined earlier:
```go
tmpl, err := template.New("facts").Parse(newsTemplate)
```

Next, all you need to do is just to execute the template with the facts, and the http.ResponseWriter:
```go
facts := factsStore.getAll()
err = tmpl.Execute(w, facts)
```

***

## Exercise 5 - use MentalFloss API <img src="https://github.com/egonelbre/gophers/blob/master/vector/science/jet-pack.svg" width="55">

### Goal

Use MetnalFloss API for fetching the facts and initialize the store, instead of the static data

### End result

`GET /facts` should show facts from MentalFloss.

> This will be done by sending request to the external provider (MentalFloss) to fetch facts and saving them in the store
You can use this API for fetching the data: ``http://mentalfloss.com/api/facts``

### Steps

##### Create a mentalfloss struct

Open a new file names `mentalfloss.go`, still in the same folder (package main).

In that file, create a struct names `mentalfloss`, for now it will be an empty struct:

```go
type mentalfloss struct{}
```

This struct will used as the provider for fetching the facts.  

##### Add functionality to `metnalfloss`

Attach a method for fetching the facts to `mentalfloss`:

```go
func (mf mentalfloss) Facts() ([]fact, error) {…}
``` 

For fetching the facts, call `http://mentalfloss.com/api/facts` using `http.Get`:

```go
resp, err := http.Get("http://mentalfloss.com/api/facts")
if err != nil {
	...
	return nil, err
}
defer resp.Body.Close()
```

> A `defer` statement defers the execution of a function until the surrounding function returns. This is how we make sure that we close the response body before we exit the function.  

[This API](http://mentalfloss.com/api/facts) returns an array of JSON representation of MentalFloss facts.

The fields which interet us in the API are:
```json
[
  {
    "fact": "fact text",
    "primaryImage": "image/url"
  },
  {
    "fact": "other fact text",
    "primaryImage": "image/url"
   }
]

```

You will need to parse the response body into a custom struct like in exercise 3 using `json.Unmarshal`

Here, a request struct, matching the payload of the response can be:

```go
var items []struct {
    FactText     string `json:"fact"`
    PrimaryImage string `json:"primaryImage"`
 }
``` 

Like in exercise 3, parse the response body into a custom struct using `json.Unmarshal`.   

##### Use `mentalfloss` instead of static data

In `main` function, replace the hard coded facts with facts from `mentalfloss`.

***

## Exercise 6 - separate to packages <img src="https://github.com/egonelbre/gophers/blob/master/vector/superhero/gotham.svg" width="95">

### Goal

Separate structs into packages.
- Move the fact entity and store into `fact` package
- Move mentalfloss functionality into `mentalfloss` package
- Move the http handlers into a into `http` package (which we will create in our project, not `net/http`...)

### Steps

##### Create package `fact`

Create a new folder named `fact` and a new file named `fact.go`. This file will contain the fact entity and the store.
In the top of the file add

```go
package fact
````

Move the `fact` and the `store` structs into that file.

> Make sure that the struct `Fact` is exported (capitalized)
  
##### Create package `mentalfloss`

Create a new folder `mentalfloss` and a new file named `mentalfloss.go` and move mentalfloss struct and methods into that folder

Set the package name in that file as `package mentalfloss`

> You will encounter compile error since now the `fact` is in another package.
You will need to import your fact package And replace `fact` with `fact.Fact`.\
(for example in exercise 6 `import "github.com/FTBpro/go-workshop/coolfacts/exercise6/fact"`)

##### Create package `http`

The goal is to separate `http.HandlerFunc` logic outside of main

Create a new folder `http` and some `.go` file. In this package create a struct named `handler` which will hold a field of the fact store.

Example:

```go
type FactsHandler struct {
	FactStore *fact.Store
}
```

Create methods for handling the request
* `func (h *FactsHandler) Ping(w http.ResponseWriter, r *http.Request)`
* `func (h *FactsHandler) Facts(w http.ResponseWriter, r *http.Request)`

> Move the anonymous `http.HandleFunc` from main and put as this struct's methods (these with the signature `func(w http.ResponseWriter, r *http.Request)`)

In main, init `FactsHandler` struct with the `factStore`.

> Since we are already importing `net/http` in main, we need to rename the name we will use for our own http ([How to import and use different packages of the same name](https://stackoverflow.com/questions/10408646/how-to-import-and-use-different-packages-of-the-same-name-in-go-language))

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

From main, replace the anonymous functions and use this handlerer methods for registering to the entpoints

For example:

```go
http.HandleFunc("/ping", handlerer.Ping)
```

***

## Exercise 7 - add ticker for updating the facts <img src="https://github.com/egonelbre/gophers/blob/master/vector/superhero/zorro.svg" width="55">

### Goal

Use go channel and ticker for updating the fact inventory

### End result

Every const time a ticker will send a signal to a `thread` (go built-in) that will fetch new fact from provider (mentalfloss)

### Steps
1. Init a context.WithCancel (remember to defer its closer…)
2. Add a function - func updateFactsWithTicker(ctx context.Context, updateFunc func() error)
    1. (Outside from updateFactsWithTicker) Create the updateFunc from step 7.2. that updates the store from an external provider
    2. (Within the updateFactsWithTicker) Create a time.NewTicker 
    3. (Within the updateFactsWithTicker) Add a go routine with a function that accepts the context
        1. Inside the function add an endless loop that will select from the channel (the ticker channel and the context one)
            1. If the ticker channel (ticker.C) was selected - use the given updateFunc to update store
            2. If the context channel (context.Done()) was selected -return (it means the main closed the context)
            
*** 

## Exercise 8 - refactor <img src="https://github.com/ashleymcnamara/gophers/blob/master/SPACEGIRL1.png" width="55">

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

Create service struct which will have fields for these as inverted dependencies.

For example

```go
type service struct {
	provider Provider
	store    Store
}
```

Add initializer - `func NewService(s Store, r Provider) *service`

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
