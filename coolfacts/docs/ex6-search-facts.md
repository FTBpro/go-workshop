# Part 6

In this exercise, you'll add functionality to the `getFacts` API. The server will be able to return facts for a specific topic. In addition, the API will require the client to add a limit to the request.

### The starting point
Be sure to be on branch `v6-search-facts`. Go on and navigate to `go-workshop/coolfacts` directory. You will see a bunch of files and TODOs. We will go over them in the rest of this document.

# **Getting Started**
In the client, the command `getFacts` now receives two new argument:
```commandline
$ > getFacts [limit] [topic (optional)]
```
You will implement this new functionality on the server side. Limit will be mandatory, and topic will be optional.

## Added Tests

You can notice that we've added tests for the server under file `.../cmd/coolfacts_server/server_test.go`. The structure of the tests is the same as we saw earlier, but there is a couple of new concepts that we can go over. Let's look in the `t.run` scope in `Test_server_GetFacts`:
```go
t.Run(tt.name, func(t *testing.T) {
    mockService := mockFactsService{
        factsToReturn: tt.want,
    }

    srv := server.NewServer(&mockService)
    ts := httptest.NewServer(srv)

    res, err := http.Get(ts.URL + "/facts" + tt.queryParamsToSend)
    require.NoError(t, err)
    require.Equal(t, tt.expectedHTTPStatus, res.StatusCode)

    if tt.wantErr {
        return
    }

    gotFacts, err := factsFromResponse(t, res)
    require.NoError(t, err)

    require.Equal(t, tt.expectedFilters, mockService.filtersGot)
    expectEqualFacts(t, tt.want, gotFacts)
})
```

### `mockFactsService`
Our server depends on `FactsService` interface. Since we don't wish to use a real service, we use a mock. our type `mockFactsService` implements this interface, so we can use it to initialize the server
```go
    mockService := mockFactsService{
        factsToReturn: tt.want,
    }

    srv := server.NewServer(&mockService)
```

We initialize our mock with the facts we defined in our `testCase`. Go over the implementation of this mock which is a really simple Go code.

Next you can notice this line of code:
```go
ts := httptest.NewServer(srv)
```
We use a Go package `httptest` which provides utilities for HTTP testing. The method `httptest.NewServer` starts and returns a new `httptest.Server`. This server has a field `URL` which is a base URL of form "http://ipaddr:port". When we will issue a call to this URL, the request will be directed to our `srv` that we've passed.  

So we can use this URL to issue a "real" http request:
```go
res, err := http.Get(ts.URL + "/facts" + tc.queryParamsToSend)
```

## Step 1 - coolfact/fact.go

For supporting the new use case, you will add a new type in the entity package `coolfact`. This type will represent the filters that the service supports.

- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Add the type `Filters` that will hold the fields that can be used for filtering the facts.

## Step 2 - fix signatures
As said, the method `GetFacts` should receive two new arguments.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Go over the project and fix all the signatures of the service and repo methods and in all the interfaces that requires them.

## Step 3 - inmem/factsrepo.go
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">In here you can notice that the type `factsRepo` now holds `factsByTopic map[string][]coolfact.Fact` instead of just `facts []coolfact.Fact`. Use this field and fix all the implementation inside the repo.
    - `NewFactsRepository` - fix implementation.
    - `GetFacts` - fix method signature and implementation.
        - Note that `topic` is optional.
    - `CreateFact` - fix implementation.

## Step 4 - coolfact_service/server.go
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">In `HandleGetFacts`, read the query params from the request and use them to initialize `coolfact.Filters` struct to be sent to the service method.
    - Use `r.URL.Query()` for accessing the query params.
    - Note that the limit is mandatory, meaning that if there isn't limit (or the limit isn't string), you should return bad request. In this case you can call the method `HandleBadRequest`.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `HandleBadRequest`. Just like any other error response, but the status should be 400 (`http.StatusBadRequest`)

# Test, Build, And Run
If all is implemented, you should be able to run tests and see them pass.
And you should see this:

TODO: add gif

# full Walkthrough

## Step 1 - coolfact/fact.go
We'll add the new type for the available filters:
```go
type Filters struct {
	Topic string
	Limit int
}
```

## Step 2 -fix signature

The method `GetFacts` receive two new arguments:

We'll change in the service:
```go
type Repository interface {
        GetFacts(filters Filters) ([]Fact, error)
        CreateFact(fct Fact) error
}

// code omitted

func (s *service) GetFacts(filters Filters) ([]Fact, error) {...}
```

In the repo:
```go
func (r *factsRepo) GetFacts(filters coolfact.Filters) ([]coolfact.Fact, error) {...}
```

And in the server:
```go
type FactsService interface {
	GetFacts(filters coolfact.Filters) ([]coolfact.Fact, error)
	CreateFact(fact coolfact.Fact) error
}
```

## Step 3 - the repo
We'll fix the initializer. We'll go over the facts, and add them under the right key.
```go
func NewFactsRepository(facts ...coolfact.Fact) *factsRepo {
	factsByTopic := map[string][]coolfact.Fact{}
	for _, fact := range facts {
		factsByTopic[fact.Topic] = append(factsByTopic[fact.Topic], fact)
	}

	return &factsRepo{
		factsByTopic: factsByTopic,
	}
}
```
We'll fix the `GetFacts` to support the new filters:
```go
func (r *factsRepo) GetFacts(filters coolfact.Filters) ([]coolfact.Fact, error) {
	var facts []coolfact.Fact
	if filters.Topic != "" {
		facts = r.factsByTopic[filters.Topic]
	} else {
		facts = r.allFacts()
	}

	sort.Sort(byCreatedAt(facts))

	if filters.Limit < len(facts) {
		facts = facts[:filters.Limit]
	}
	
	return facts, nil
}

func (s *factsRepo) allFacts() []coolfact.Fact {
	var allFacts []coolfact.Fact
	for _, facts := range s.factsByTopic {
		allFacts = append(allFacts, facts...)
	}

	return allFacts
}
```
>
>Note how we used append: `append(allFacts, facts...)`. `append` is a variadic function:
>```go
>func append(slice []Type, elems ...Type) []Type
>```
 
Finally, we'll fix the `CreateFact` method, just adding the fact to the facts of its topic:
```go
func (r *factsRepo) CreateFact(fact coolfact.Fact) error {
	r.factsByTopic[fact.Topic] = append(r.factsByTopic[fact.Topic], fact)

	return nil
}
```

## Step 4 - Server
In `HandleGetFacts`, the server needs to call the service's method `GetFacts(filters)`. The filters will be represented by query params. For example, if the client wants to get 10 facts of topic `TV`, it will call:
```json
GET "/facts?limit=10&topic=TV"
```
Query params are received as strings. We will convert the limit to int using go package `strconv`. If we can't convert, we will return bad request:
```go
func (s *server) HandleGetFacts(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling getFact ...")
	
	limitString := r.URL.Query().Get("limit")
	if limitString == "" || limitString == "0" {
		err := fmt.Errorf("limit is mandatory")
		s.HandleBadRequest(w, err)
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		err = fmt.Errorf("limit isn't int")
		s.HandleBadRequest(w, err)
	}

	filters := coolfact.Filters{
		Topic: r.URL.Query().Get("topic"),
		Limit: limit,
	}

	facts, err := s.factsService.GetFacts()
	// code omitted
```

# Finish :boom: