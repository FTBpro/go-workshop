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
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">The returned facts from the service should have a new field for the created time. Add this field in the formatted response in the method `HandleGetFacts`
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

## Step 1 - Implement The BL
### coolfact/fact.go
In our entity, we add a new field `CreatedAt` for specifying the time the fact was created.
```go
type Fact struct {
	Image       string
	Description string
	CreatedAt   time.Time // new
}
```

### coolfact/service.go
The service requires another method from the `Repository` interface
```go
type Repository interface {
	GetFacts() ([]Fact, error)
	CreateFact(fct Fact) error // new
}
```
And in the `CreateFact` we just calling the repo:
```go
func (s *service) CreateFact(fact Fact) error {
	if err := s.factsRepo.CreateFact(fact); err != nil {
		return fmt.Errorf("factsService.CreateFact: %w", err)
	}

	return nil
}
```

## Step 2 - The repo
In the `CreateFacts`, before reurning the facts, we will sort them based on their `CreatedAt`
```go
func (r *factsRepo) GetFacts() ([]coolfact.Fact, error) {
	sort.Sort(byCreatedAt(r.facts))

	return r.facts, nil
}
```
We use method `sort.Sort`. This method expects an argument that implements `sort.Interface`:
```go
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
```
Using these methods, it sorts the given slice. Since we can't our `r.facts` slice (because it doesn't implement the `sort.Interface`), we will use the type `byCreatedAt`:
```go
type byCreatedAt []coolfact.Fact

func (s byCreatedAt) Len() int {
	return len(s)
}

func (s byCreatedAt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byCreatedAt) Less(i, j int) bool {
	return s[i].CreatedAt.After(s[j].CreatedAt)
}
```

Note the type conversion: `byCreatedAt(r.facts)`. Type conversion is simply to convert some value to other type.
This is possible since the two types are compatible.

## Step 3 - Server
The service exports a new API for creating the fact. The http method is `POST`, and the path is `"/facts"`. We will add a new case in our `ServeHTTP` method:
```go
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("incoming request", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		// code omitted
	case http.MethodPost:
		switch strings.ToLower(r.URL.Path) {
		case "/facts":
			s.HandleCreateFact(w, r)
		default:
			err := fmt.Errorf("path %q wasn't found", r.URL.Path)
			s.HandleNotFound(w, err)
		}
	default:
		// code omitted
	}
}
```
For decoding the request, we would use a struct `createFactRequest` which is a representation of the request payload:
```go
type createFactRequest struct {
	Image       string `json:"image"`
	Description string `json:"description"`
}

func (r createFactRequest) ToCoolFact() coolfact.Fact {
	return coolfact.Fact{
		Image:       r.Image,
		Description: r.Description,
		CreatedAt:   time.Now(),
	}
}
```
Note that in the method `ToCoolFact` we also set the `CreatedAt` to `time.Now()`.

And then we could implement `HandleCreateFact`:
```go
func (s *server) HandleCreateFact(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling createFact ...")

	var request createFactRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		err = fmt.Errorf("server.HandleCreateFact failed to decode request: %s", err)
		s.HandleError(w, err)
		return
	}

	if err := s.factsService.CreateFact(request.ToCoolFact()); err != nil {
		err = fmt.Errorf("server.HandleCreateFact: %s", err)
		s.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
```

Finally, in `HandleGetFacts` we just need to add the `createdAt` to the response:
```go
func (s *server) HandleGetFacts(w http.ResponseWriter) {
	log.Println("Handling getFact ...")
	facts, err := s.factsService.GetFacts()
	if err != nil {
		formattedFacts[i] = map[string]interface{}{
			"image":       coolFact.Image,
			"description": coolFact.Description,
			"createdAt":   coolFact.CreatedAt, // new
		}
	}

	// code omitted
}

```

## Step 4 - The Client
The client receives the response of the facts from the service, so we'll add the field `createdAt`, so it will be decoded as well:
```go
type getFactsResponse struct {
	Facts []struct {
		Image       string    `json:"image"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"` // new
	} `json:"facts"`
}
```

We will implement the method GetLAstFact using the existing method `GetAllFacts`:
```go
func (c *client) GetLastCreatedFact() (coolfact.Fact, error) {
	allFacts, err := c.GetAllFacts()
	if err != nil {
		return coolfact.Fact{}, fmt.Errorf("GetLastCreatedFact: %w", err)
	}

	if len(allFacts) == 0 {
		return coolfact.Fact{}, fmt.Errorf("fact not found")
	}

	return allFacts[0], nil
}

```

And finally, we implement the method `CreateFact` which call the service
```go
func (c *client) CreateFact(fct coolfact.Fact) error {
	ul := c.endpoint + pathCreateFact
	// First we are preparing the payload
	payload := map[string]interface{}{
		"image":       fct.Image,
		"description": fct.Description,
	}
	// we need io.Reader to create a new http request.
	// we will create bytes.Buffer which implement this interface
	postBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("client.CreateFact failed to marshal payload: %v", err)
	}
	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest(http.MethodPost, ul, responseBody)
	if err != nil {
		return fmt.Errorf("client.CreateFact failed to create request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("client.CreateFact failed to do request: %v", err)
	}

	defer func() {
		if res != nil && res.Body != nil {
			_, _ = io.Copy(ioutil.Discard, res.Body)
			_ = res.Body.Close()
		}
	}()

	if res.StatusCode != http.StatusOK {
		errMessage, err := c.readError(res)
		if err != nil {
			return fmt.Errorf("client.CreateFact: %s", err)
		}

		return fmt.Errorf("client.CreateFact got an error from server. status: %d. error: %s", res.StatusCode, errMessage)
	}

	return nil
}
```

# Finish!