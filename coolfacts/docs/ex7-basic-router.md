# Part 7

In this exercise you will build a simple router for acting as the main `http.Handler` and to direct the request into the appropriate handler

### The starting point
To get started, run this command to clone the necessary exercise materials in a convenient folder:
```commandline
$ git clone --branch v5-create-fact https://github.com/FTBpro/go-workshop.git
```

# **Getting Started**
Until the last exercise, our server acted as the `http.Handler` to the method `http.ListenAndServe`. The method `server.ServeHTTP` handled the logic for directing the request into the right handler. For every new route you'll had to add a case in this method.

In real-world application we usually would prefer to use a framework over http that exports a router for more convinient handling.
- Automatically matches a Method + Path to an `http.Handler`.
- Easy path params handling. For example `/facts/:factID` -> the `factID` param will be given to us by the router.
- Usually we will have more than one domain in the application, for each domain we will have its specific routes. Using a shared router will make it easier.
- Easy to add middlewares across all the APIs.

In this part we will implement a basic router that export a method for registering to a route:
```go
func Handle(method, path string, handler http.HandlerFunc)
```

Our server will use it for each route, instead of handling itself the logic inside the `ServeHTTP` method. Meaning, our server won't be an `http.Handler` anymore.

going over the branch, you will notice a couple of changes, and a new package - `coolhttp`. In this package we will implement our router. This package is agnostic to the application, meaning doesn't know anything about the server or the coolfacts. In a system of many services, this package could be shared between all of them.

## Step 0 - Notice `coolfacts_server/main.go`
We are initializing our router and setting its not found handler:
```go
	router := coolhttp.NewRouter()
	router.SetNotFoundHandler(server.HandleNotFound)
```
Then we are calling a new method in the server:
```go
server.RegisterRouter(router)
```
This server uses the router for registering any route with the corresponding handler. Later you'll implement it.

Finally, we use this router:
```go
	if err := http.ListenAndServe(":9002", router); err != nil {
		//...
	}
```

Every request will get to the `router.ServeHTTP` which you will implement.

## Step 1 - Implement the router (in coolhttp/router.go)
Heading over to the coolhttp/router.go file, you'll see the new package. You'll need to implement its methods:
```go
func (r *router) Handle(method, path string, handler http.HandlerFunc) {...}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {...}
```
The `Handle` method will be called for each route, on registration from the server. You will need to same some sort of state for handling the logic, so that in the method `ServeHTTP` you could call to the right `http.HandlerFunc` that was registered to the corresponding http method and path.

If no such handler exist, you should call the `notFoundHandler`.

## Step 2 - coolfacts_server/server.go
You will see the the method `ServeHTTP` was deleted, and instead there is a new method `func (s *server) RegisterRouter(router Router)`.
Implement this method. Reminder, the routes are:
- GET "/ping"
- GET "/facts"
- POST "/facts"

# Test, Build and Run

## Full Walkthrough