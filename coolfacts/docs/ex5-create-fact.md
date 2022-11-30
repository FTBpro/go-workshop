# Part 5

In this exercise, you will implement a new API in the server for creating a fact.
You will also implement in the client the method that calls this API with an input from the user.
Another command that the client will export is getting the last created fact. 

The starting point
To get started, run this command to clone the necessary exercise materials in a convenient folder:
```commandline
$ git clone --branch v5-create-fact https://github.com/FTBpro/go-workshop.git
```

# **Getting Started**
Take a look around the program, there are a bunch of new TODOs and functionality. We will go over them in the next section.

## Step 0 - Notice `coolfacts_clinet/main.go`
We added two new commands:
- `"createFact"` for creating a new fact in the server. 
- `"getLastFact` for returning the last created fact.
In the `client`, you will implement the methods for supporting these command.

In addition, for returning the last created fact, we need a way to which fact was created last. For supporting this, you will add a new field in the entity `coolfact.Fact` - `CreatedAt`.
For this, you will get to know the go package `time` - This package Package time provides functionality for measuring and displaying time.
Everytime we want to use field that representing a time, we will use types from this package, instead of string or int or other types. (Not include of course the http response payload).  
The basic type is `time.Time` which represents an instant in time with nanosecond precision.

For sorting, you will learn and use the go package `sort`.

## Step 1 - Implement The BL
### coolfact/fact.go
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">For supporting the sorting, add the field for the createdAt.

### coolfact/service.go
The service has new functionality for creating a fact. 
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">In the `Repository` interface, add the method that the service requires.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement the method `CreateFact`.

## Step 2 - The repo
The repository now have new method for creating a fact. In addition you need to change the method for getting facts and sort them by their created time. 
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">In `GetFacts` add the sorting. For the sort, you will use the type `byCreatedAt`. To learn about sorting using `sort.Sort`, read the TODO in the code in the `GetFacts` method.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement the method `CreateFact`

## Step 3 - Server
The service has a new API:
```json
POST "/facts"

Request: 
{
  "image": "...",
  "description": "..."
}

Response:
Success: 200
Failure:
- Missing field: 404
- Error in the server: 500
```
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">The returned facts from the service should have a new field for the created time. Add this fieldin the formatted response in the method `HandleGetFacts`
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Add the required method in the `FactsService` interface.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Add a new case in the method `ServeHTTP`. Note that the http method is POST request, and the path is "/path". In case of a such a request, call the server method `HandleCreateFact`
  - <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Add  struct `factService` for decoding the request into. This struct should be the representation of the request body.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement method `HandleCreateFact`.

## Step 4 - The Client
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement method `GetLastCreatedFact`. Note that the server doesn't have a dedicated API for this, so use the method `GetAllFacts`
  - Of course in real world application it's not a good practice, but this exercise is not a real world application.
  - <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Since the server returns a new field for the created time, add this field in the struct `getFactsResponse` which is the representation of the response.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement the method `CreateFact`.

## Building and Running

If everything is implemented well, this is what the final result should look like when running the program, in one tab we are running the server as we did in the previous exercise, and in another tab we are running the client, and writing the command for getting the facts:

TODO:(oren) add gif

## Full Walkthrough

