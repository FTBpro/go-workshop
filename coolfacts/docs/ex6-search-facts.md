# Part 6

In this exercise you will add some more interesting use case for the get facts API. The option to search facts by topic, and add a limit. 

### The starting point
To get started, run this command to clone the necessary exercise materials in a convenient folder:
```commandline
$ git clone --branch v5-create-fact https://github.com/FTBpro/go-workshop.git
```

# **Getting Started**
You will add two query params to the get-facts API. `limit` and `topic`. Usually, we don't want that the client will get all resources, so we require they to send a limit. We also add the option to filter facts by a specific topic.

In the client, we've added args to the command `getFacts`
```commandline
$ > getFacts [limit] [topic (optional)]
```
You will implement this new ability at the server side.

The tests were updated for testing the new functionality. After all is implemented they should pass. 

## Step 1 - coolfact/fact.go

For supporting the new use case, you will add new type in the entity package coolfact. The function `GetFacts` will be changed to `SearchFacts` since it will get the filters and the repo will need to search the appropriate facts instead of just returns all of them.

- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Add the type `Filters` that will hold the fields that can be used for filtering the facts.

## Step 2 - fix signatures
As said, the method `GetFacts` should be called `SearchFacts` and should receive arg `coolfact.Filters`
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Go over the project and fix all the signatures of the service and repo methods and in all the interfaces that requires them. There is a TODO next to any method/interface that needs to be fixed.

## Step 3 - inmem/factsrepo.go
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">In here you can notice that the type `factsRepo` now holds `factsByTopic map[string][]coolfact.Fact` instead of just `facts []coolfact.Fact`. Use this field and fix all the implementation inside the repo.
  - `NewFactsRepository` - fix implementation.
  - `GetFacts` - fix method name/signature and implementation.
    - Note that `topic` is optional.
  - `CreateFact` - fix implementation.

## Step 3 - coolfact_service/server.go
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">In `HandleGetFacts`, reaad the query params from the request and use them for inialize `coolfact.Filters` struct to be sent to the service method.
  - Use `r.URL.Query()` for accessing the query params.
  - Note that the limit is mandatory, meaning that if there isn't limit (or the limit isn't string), you should return bad request. In this case call the method `HandleBadRequest` which you will implement.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `HandleBadRequest`. Just like any other error response, but the status should be 400 (`http.StatusBadRequest`)

# Build And Run
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

The method `GetFacts` is changed to `SearchFacts(filters coolfact.Filters) ([]coolfact.Fact, error)`

We'll change in the service:
```go
type Repository interface {
        SearchFacts(filters Filters) ([]Fact, error) // Was changed from `GetFacts`
	CreateFact(fct Fact) error
```
```go
func (s *service) SearchFacts(filters Filters) ([]Fact, error) {...} // Was changed from `GetFacts`
```

In the repo:
```go
func (r *factsRepo) SearchFacts(filters coolfact.Filters) ([]coolfact.Fact, error) {...}
```

And in the server:
```go
type FactsService interface {
	SearchFacts(filters coolfact.Filters) ([]coolfact.Fact, error) // Was changed from `GetFacts`
	CreateFact(fact coolfact.Fact) error
}
```

## Step 3 - the repo
We'll fix the initializer. We'll go over the facts, and add them under the right key
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
We'll fix the `GetFacts` to support the new filters: (Note that it is only one example of implementation, not focusing on performance at all)
```go
func (r *factsRepo) SearchFacts(filters coolfact.Filters) ([]coolfact.Fact, error) {
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

Finally, we'll fix the `CreateFact` method, just add the fact to the facts of its topic:
```go
func (r *factsRepo) CreateFact(fact coolfact.Fact) error {
	r.factsByTopic[fact.Topic] = append(r.factsByTopic[fact.Topic], fact)

	return nil
}
```

## Step 4 - Server
In `HandleGetFacts`, the server needs to call the service's method `SearchFact(filters)`. The filters will be built from the query params of the request. For example, if the client wish to get 10 facts of topic `TV`, it will call:
```commandline
GET /facts?limit=10&topic=TV
```
Both the arguments will be received as strings, but we will need to convert the limit top int. If we can't, we'll return bad request
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