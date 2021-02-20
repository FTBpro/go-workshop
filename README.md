# welcome
This is a step-by-step tutorial for creating a simple web server for fetching facts from MentalFloss API, and list them in a simple HTML template

## Feedback
Please help us improve ourselves for the next time with a quick feedback https://forms.gle/fNxcqSBdyZepoyVT6.

## Slides
all relevant slides are available at [go slides]: http://present.minutemediaservices.com/gopresent

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
* [Idiometic Go links](#usefull-links-for-idiometic-go)

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

Create `/facts` endpoint for listing facts in a JSON format by using a repository.

### End result

`http://localhost:9002/facts` will show a JSON of all facts in repository, for now it will be hard coded facts.

### Steps

##### Create fact struct

Create a struct named fact (`type Fact struct {...}`)\
The `fact` struct should have 2 string fields : `Image`, `Description`

##### Create repository struct

Create a struct named repository (or repo)\
The repository can use in memory cache for storing the facts, it can be done by one field `facts` of type `[]fact` (a slice of facts).

Add repository functionality:
* `func (r *repository) getAll() []fact {…}`
  * The method should return all facts in the `repository.facts` field
* `func (r *repository) add(f fact) {…}`
  * The method should add the given fact f to the repository\
  For adding to a slice you can use `repository.facts = append(repository.facts, f)`
  
Init the repository from `main` with some static data. 

##### Register `http.HandleFunc` to `/facts` pattern  

Like in the ping from previous exercise, use an anonymous function as an argument to `http.HandleFunc` function.

In this function you will:
* Get all the facts from the repository
* Write the facts to the `ResponseWriter` in a JSON format
> Use `json.Marshal` to format the struct as json and to write to the `ResponseWriter` 

***

## Exercise 3 - create a new fact <img src="https://github.com/egonelbre/gophers/blob/master/vector/fairy-tale/sage.svg" width="55">

### Goal

Create a new fact by a POST request.

### End result

Create a new fact and add it to the repository by issuing a `POST /facts` request with the next payloiad:
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

Finally, after we have this struct filled, create a new fact from it, and add it to the repository.

For adding it to the repository you should use the `factRepo.Add` from exercise 2.

***

## Exercise 4 - list the facts as HTML <img src="https://github.com/egonelbre/gophers/blob/master/vector/science/power-to-the-linux.svg" width="95">

### Goal

* Using HTML template
* Replace the JSON representation from exercise 2 with an HTML

### End result

`GET /facts` will list the facts in HTML
`http://localhost:9002/facts` will show a all facts in repository in HTML.

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
facts := factsRepo.getAll()
err = tmpl.Execute(w, facts)
```

***

## Exercise 5 - use MentalFloss API <img src="https://github.com/egonelbre/gophers/blob/master/vector/science/jet-pack.svg" width="55">

### Goal

Use MetnalFloss API for fetching the facts and initialize the repository, instead of the static data

### End result

`GET /facts` should show facts from MentalFloss.

> This will be done by sending request to the external provider (MentalFloss) to fetch facts and saving them in the repository
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
- Move the fact entity and repository into `fact` package
- Move mentalfloss functionality into `mentalfloss` package
- Move the http handlers into a into `http` package (which we will create in our project, not `net/http`...)

### Steps

##### Create package `fact`

Create a new folder named `fact` and a new file named `fact.go`. This file will contain the fact entity and the repository.
In the top of the file add

```go
package fact
````

Move the `fact` and the `repository` structs into that file.

> Make sure that the struct `Fact` is exported (capitalized)
  
##### Create package `mentalfloss`

Create a new folder `mentalfloss` and a new file named `mentalfloss.go` and move mentalfloss struct and methods into that folder

Set the package name in that file as `package mentalfloss`

> You will encounter compile error since now the `fact` is in another package.
You will need to import your fact package And replace `fact` with `fact.Fact`.\
(for example in exercise 6 `import "github.com/FTBpro/go-workshop/coolfacts/exercise6/fact"`)

##### Create package `http`

The goal is to separate our application `http.HandlerFunc` logic outside of main.

Create a new folder `http` and some `.go` file. In this package create a struct named `handler` which will hold a field of the fact repository.

Example:

```go
type FactsHandler struct {
	FactRepo *fact.Repository
}
```

Create methods for handling the request
* `func (h *FactsHandler) Ping(w http.ResponseWriter, r *http.Request)`
* `func (h *FactsHandler) Facts(w http.ResponseWriter, r *http.Request)`

> Move the anonymous `http.HandleFunc` from main and put as this struct's methods (these with the signature `func(w http.ResponseWriter, r *http.Request)`)

In main, init `FactsHandler` struct with the `factRepo`.

You may noticed that `http` is already taken as a package name by `net/http`. You're not wrong, we can't use both packageS in one file and still call each package `http`. But we can [rename the import name](https://stackoverflow.com/questions/10408646/how-to-import-and-use-different-packages-of-the-same-name-in-go-language):
```go
package main

import (
	"net/http"
	facthttp "github.com/FTBpro/go-workshop/coolfacts/exercise8/http"
)

func main() {
	handlerer := facthttp.FactsHandler{
		FactRepo: &factsRepo,    
	}
	
	http.HandleFunc("/ping", handlerer.Ping)
}
```

> Creating a package `http` may not be the best idea. If we create a package with an infrastructure name (like sql, mentalfloss...) we need to make sure we import it only from main. 

***

## Exercise 7 - add ticker for updating the facts <img src="https://github.com/egonelbre/gophers/blob/master/vector/superhero/zorro.svg" width="55">

### Goal

Use go channel and ticker for updating the fact inventory

### End result

Every specified time a ticker will send a signal using a `channel` (go built-in) that will trigger fetch new fact from provider (mentalfloss)

### Steps
1. Init a `ctx, cancelFunc := context.WithCancel(context.Background())` (use the cancel func in the end of the run to terminate the go routing execution)
    1. >`context.WithCancel(context.Background())` returns a context which have a done channel on it that can be used as follows : `<-ctx.Done()` to signal termintation.
    2. >`context.WithCancel(context.Background())` also returns a cancel func that can be called at the end of the run - it will send a message to the `ctx.Done()` channel.
    3. > The call to the `cancelFunc()` can be done using `defer` which invokes whatever defined after it at the end of the function that containes the declaration:
    ``` go
    func example(){
      ctx, cancelFunc := context.WithCancel(context.Background())
      defer cancelFunc() //defined but not invoked
      doSomething(ctx)
      //This is the end of the function so defer is invoked
    }
    ```
2. Add a function - func updateFactsWithTicker(ctx context.Context, updateFunc func() error)
    1. (Outside from updateFactsWithTicker) Create the updateFunc from step 7.2. that updates the repository from an external provider
    2. (Within the updateFactsWithTicker) Create a time.NewTicker 
    3. (Within the updateFactsWithTicker) Add a go routine with a function that accepts the context
        1. Inside the function add an endless loop that will select from the channel (the ticker channel and the context one)
            1. If the ticker channel (ticker.C) was selected - use the given updateFunc to update repository
            2. If the context channel (context.Done()) was selected -return (it means the main closed the context)
            
*** 

## Exercise 8 - refactor <img src="https://github.com/ashleymcnamara/gophers/blob/master/SPACEGIRL1.png" width="55">

### Goal
Decouple and break dependancies using interfasces

### Steps

##### Extract repository logic to package `inmem`

We need to encapsulate the process of storing the facts. You can think of this package like some kind of an interactor between the actual caching/persistent layer to the application domain.

In this example we will create a package dedicated for storing the facts in a simple `slice`. The package we'll create will be called `inmem`. (Accordinagly, if we will use SQL, we will create package `sql`)

After creating packe `inmem`, we need to move the `repository` functionality currently in package `fact`.\
For exporting it we will use function NewFactRepository. The consumer (main) will call it via `inmem.NewFactRepositry()`

```go
pkg inmem

type factRepository struct {
	facts []fact.Fact
}

func NewFactRepository() *factRepository {
	return &factRepository{}
}

func (r *factRepository) Add(f fact.Fact) {
	// code
}

func (r *factRepository) GetAll() []fact.Fact {
	// code
}
```

##### Create fact service for abstracting how we update the facts

We need some kind of "service" to handle our update logic. This service will be initialized with the repository and the mentalfloss provider, and have an `Update` method.

We better not limit ourselves to only mentalfloss and only in memory cache. We can do that by using interfaces instead the concrete types.

Example:

In package `fact`, declare two interfaces:

```go
package fact

type Provider interface {
	Facts() ([]Fact, error)
}

type Repository interface {
	Add(f Fact)
	GetAll() []Fact
}
```

The service will be a struct which we will export using NewService function that takes a `Provider` and a `Repository`.
The service will have `UpdateFacts` method that take no parameters.

For example

```go
// continue package fact

type service struct {
	provider Provider
	repository    Repository
}

func NewService(s Repository, r Provider) *service {
	// code
}

func (s *service) UpdateFacts() error {
	// code
}
```

> Although the service is updating the facts, it doesn't know from which provider or what is the persistent layer. That means we could easily replace inmem with a db, switch providers, and add middlewares (decorators) for logging and other stuff.

> Instead of a service, we could just create an exported function `fact.UpdateFacts(p Provider, s Repository) error`, which would achieve the same goal and have some advantages, same as `updateFactsFunc` principle in exercise 7.

##### Add abstraction in local `http` package

By now you can see that we broke our custom `http` package. This is because `*fact.Repository` isn't defined anymore (we moved the repository to `inmem`).

We'll use here the interface principle as well. We'll declare same interface in our `http`:

```go
package http

type FactRepository interface {
	Add(f fact.Fact)
	GetAll() []fact.Fact
}

type factsHandler struct {
	factRepo FactRepository
}

func NewFactsHandler(factRepo FactRepository) *factsHandler {
	return &factsHandler{
		factRepo: factRepo,
	}
}
```

> Same as in the service, we can initialize the handler with a different persistent layer, add middlewares and more

##### Rename mentalfloss struct
For making it more readable (perhaps), we will use the name `provider` instead of the `mentalfloss` struct. This is a naming convention when we are interacting with external services.

```go
package mentalfloss

type provider struct{}

func NewProvider() *provider {
	return &provider{}
}
```

The consumer will now use it by calling `mentalfloss.NewProvider()`

##### Replace calls in main

All left to do is to replcae our way we initialize our dependancies in main.

instead of initializing the `factRepo` like this:
```go
factsRepo := fact.Repository{}
```

we will initialize it with our `inmem` package:
```go
factsRepo := inmem.NewFactRepository()
```

Instead of using the `updateFunc` from exercise 7, we'll use our `fact.NewService`:
```go
	mentalflossProvider := mentalfloss.NewProvider()
	service := fact.NewService(mentalflossProvider, &factsRepo)
```

Now we will just the `service.UpdateFacts` method to update the repository, and to use with the ticker and we're done 🥂

# Usefull links for "idiometic go":

- [Effective Go - a detailed set of the specialities and atyleguides for Go](https://golang.org/doc/effective_go.html)
- [Code review comments - common comments made during reviews of Go code](https://github.com/golang/go/wiki/CodeReviewComments)
- [Idiomatic GO - general guidelines of what not to do](https://about.sourcegraph.com/go/idiomatic-go/)

#### Package Names:
- [Go blog styleguides about naming packages](https://blog.golang.org/package-names)
- [Blog post by Dave Cheney wbout why we shouldn't use `util` or `common`](https://dave.cheney.net/2019/01/08/avoid-package-names-like-base-util-or-common)

#### Go Proverbs
- [Lecture by Rob Pike](https://www.youtube.com/watch?v=PAAkCSZUG1c)
- [The list](https://go-proverbs.github.io)

#### Errors:
- [`pkg/errors` - a popular alternative to the standart `errors` package](https://github.com/pkg/errors)
- [Article by Dave Chaney about the right error handling](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)

#### Solid design in go

- [Lecture of applying the solid principles in Go](https://www.youtube.com/watch?v=zzAdEt3xZ1M)
- [The lecture's script](https://dave.cheney.net/2016/08/20/solid-go-design)
