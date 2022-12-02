# Part 4

In this exercise, you will implement a client that we will use for calling our server.
The client will sit in another main package in our `cmd` folder, making our coolfacts Go program to have 2 application. One server and one client. 

The starting point
To get started, run this command to clone the necessary exercise materials in a convenient folder:
```commandline
$ git clone --branch v4-initial-client https://github.com/FTBpro/go-workshop.git
```

# The Goal
After completing the exercise, you will have a client application which will take one command from the terminal, to call the server and get all the facts.

# **Getting Started**
Take a look around the program, you notice a new cmd application - `coolfacts_client`, with two files, `main.go` and `client.go`

## Step 0 - Notice `main.go`
In this file you don't hva any TODO, but let's get over it so you will be familiar with what that's going on

First, we initialize the client with our server endpoint:
```go
cl := NewClient(serverEndpoint)
```

Then we're waiting for an input from the client:
```go
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		// code omitted
		
		res, err := processCmd(cl, cmd, args)
		// code omitted
	}
}

func processCmd(cl *client, cmd string, args []string) (string, error) {
	switch cmd {
	case "":
		return "", nil
	case commandGetAllFacts:
		// code omitted
	default:
		return "", errors.New("unknown command")
	}
}
```

Currently, there is only one command `commandGetAllFacts` which is a const for "getAllFacts".

## Step 1 - client.go - implement the client
Take a look in the file and notice we have client struct and initializer. When the user will send an input `getAllFacts`,
the client's method `GetAllFacts` will be called, and the client will call the server getFacts API. You will implement the method `GetAllFacts`

- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Fill struct `getFactsResponse`
  - This struct represents the JSON response from the client, and we will use it for deserializing the server's response-body into a convenient struct. Add fields corresponding to the response. Reminder, the response looks like this:
  ```json
  {
      "facts": [
          {
              "image": "...",
              "description": "..."
          }
          //...
      ]     
  } 
  ```
  - Example on JSON tags and some json package functionality - For this be sure to add json tags on the struct. (https://gobyexample.com/json)
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `ToCoolFacts()` method. You will use this method for converting the server response to the entity.
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `GetAllFacts`.
  - Notice the start of the method, we are composing the url for get facts which is "127.0.0.1:9002/facts" and calling `http.Get(...)`. The clinet must read and close the response after using it, you can see it in the `defer` block.
  - <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Finish implementing the method, handle the response as specified in the TODO in the code
- <img src="https://user-images.githubusercontent.com/5252381/204141574-767eba62-e9dd-4bc1-9d45-03bef68812aa.jpg" width="18">Implement `readResponseGetFacts` as specified in the TODO in the code.

## Building and Running

If everything is implemented well, this is what the final result should look like when running the program, in one tab we are running the server as we did in the previous exercise, and in another tab we are running the client, and writing the command for getting the facts:

TODO:(oren) add gif

## Full Walkthrough

In the following section you fill find a full walkthrough. Use it in case you are stuck.

## Step 1 - client.go

We will add fields to the struct `getFactsResponse`:
```go
type getFactsResponse struct {
	Facts []struct {
		Image       string `json:"image"`
		Description string `json:"description"`
	} `json:"facts"`
}
```
As been said, this type represent the server JSON response. The json decoding, when receiving a struct, search for the json tags, to know how to format the fields.
Note that the field are exported (public), this is for the json package to be able to see them.

We also implement method `toCoolFacts()`. We will use this method to convert the response to the entity the client needs to return. 
```go
func (r getFactsResponse) toCoolFacts() []coolfact.Fact {
	coolfacts := make([]coolfact.Fact, len(r.Facts))
	for i, fact := range r.Facts {
		coolfacts[i] = coolfact.Fact(fact)
	}

	return coolfacts
}
```

Finish implementing the method `GetAllFacts` - handling the response. In case the `StatusCode` is not `http.StatusOK`, we will read an error from the response and return it. Otherwise, we read the body and convert to our coolfact.Fact slice
```go
func (c *client) GetAllFacts() ([]coolfact.Fact, error) {
        
	// code emitted
	
	if res.StatusCode != http.StatusOK {
		errMessage, err := c.readError(res)
		if err != nil {
			return nil, fmt.Errorf("client.CreateFact: %s", err)
		}

		return nil, fmt.Errorf("client.GetLastCreatedFact got an error from server. status: %d. error: %s", res.StatusCode, errMessage)
	}

	getFactsRes, err := c.readResponseGetFacts(res)
	if err != nil {
		return nil, fmt.Errorf("client.GetLastCreatedFact: %s", err)
	}

	return getFactsRes.toCoolFacts(), nil
}
```
Notice the method `readError` which is already implemented:
```go
func (c *client) readError(res *http.Response) (string, error) {
	var errRes errorResponse
	if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
		return "", fmt.Errorf("readBody failed to read response body: %v. \nbody string is: %s", err)
	}

	return errRes.Error, nil
}
```
We use json package to decode the response. We will do it in the same way in the method `readResponseGetFacts`:
```go
func (c *client) readResponseGetFacts(res *http.Response) (getFactsResponse, error) {
	var factsRes getFactsResponse
	if err := json.NewDecoder(res.Body).Decode(&factsRes); err != nil {
		return getFactsResponse{}, fmt.Errorf("readResponseGetFacts failed to read response body: %v. \nbody string is: %s", err)
	}

	return factsRes, nil
}
```

# Finish!
Congratulation! You've just implemented a client application, and when running alongside our server, you have request-response system!

In the following exercise we will add an API for creating a fact! 