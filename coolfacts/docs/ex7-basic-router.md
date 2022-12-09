# Part 7

As you've probably notices, the handling of the request in `ServeHTTP` method is cumbersome. In this exercise you will implement a basic router for helping us do that.

### The starting point
Be sure to be on branch `v7-basic-router`. Go on and navigate to `go-workshop/coolfacts` directory. You will see a bunch of files and TODOs. We will go over them in the rest of this document.

# **Getting Started**
Up to the last exercise, our server acted as the `http.Handler` to the method `http.ListenAndServe`. The method `server.ServeHTTP` handled the logic for directing the request into the right handler. For every new route you had to add a case in this method.

In real-world application we usually would prefer to use a framework over http that exports a router for abstracting this boilerplate.
- Automatically matches a Method + Path to an `http.Handler`.
- Easy path params handling. For example `/facts/:factID` -> the `factID` param will be given to us by the router.
- Usually we will have more than one domain in the application, for each domain we will have its specific routes. Using a shared router will make it easier.
- Easy to add middlewares across all the APIs.

In this part we will implement a basic router that export a method for registering an API:
```go
func Handle(method, path string, handler http.HandlerFunc)
```

Our server will use it for each route, instead of handling itself the logic inside the `ServeHTTP` method. Meaning, our server won't be an `http.Handler` anymore.

going over the branch, you will notice a couple of changes, and a new package - `coolhttp`. In this package we will implement our router. This package is agnostic to the application, meaning that it doesn't know anything about the server or the coolfacts. In a system of many services, this package could be shared between all of them.

## Step 0 - Notice `coolfacts_server/main.go`
In the `main` function, we are initializing our router and setting its `notFoundHandler`:
```go
router := coolhttp.NewRouter()
router.SetNotFoundHandler(server.HandleNotFound)
```

:Then we are calling a new method in the server for registering the routes (which you will implement)
```go
server.RegisterRouter(router)
```

Finally, we use the router as `http.Handler`, every request will get to `router.ServeHTTP`:
```go
if err := http.ListenAndServe(":9002", router); err != nil {
    //...
}
```

## Step 1 - Implement the router (in coolhttp/router.go)
Heading over to the _coolhttp/router.go_ file, you'll see the new package. You'll need to implement its methods:
```go
func (r *router) Handle(method, path string, handler http.HandlerFunc) {...}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {...}
```
You will need to same some sort of state for handling the logic in order that in `ServeHTTP` you could call to the right `http.HandlerFunc` that was registered to the corresponding http method and path.

If no such handler exist, you should call the `notFoundHandler`.

## Step 2 - coolfacts_server/server.go
The method `ServeHTTP` was deleted, and instead there is a new method:
```go
func (s *server) RegisterRouter(router Router)
```
Implement this method. Reminder, the routes are:
- GET "/ping"
- GET "/facts"
- POST "/facts"

# Test, Build and Run

## Full Walkthrough