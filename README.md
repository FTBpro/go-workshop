# welcome

go workshop - steps

0. Hello World!

1. Hello World - http version
    1. Register /ping to http HandleFunc with a function that writes to ResponseWriter the ‘Hello World!’ string

2. End result of this step - GET /facts will return json of all facts in store
    1. Create fact struct
    2. Create store struct
    3. Add method - func (s store) getAll() []fact {…}
    4. Add method - func (s store) add(f fact) {…}
    5. Register /facts to http HandleFunc with a function that writes to ResponseWriter all facts in json format

3. End result of this step - POST /facts will create a new fact and add it to store
    1. Create a fact from req body
    2. Add fact to store

4. End result of this step - return the index results (GET /facts) with an HTML template
    1. Crate an HTML template using package `text/template` syntax
    2. Execute template with store getAll results (that means write to ResponseWriter all results in the applied template)

5. End result of this step - send request to external provider (MentalFloss) to fetch facts parse them and save them to store
    1. Create a mentalfloss struct
    2. Add method - func (mf mentalfloss) Facts() ([]fact, error) {…} - sends request to MentalFloss api and parses response (step 5.3.) to fact struct
    3. Add function - func parseFromRawItems(b []byte) ([]fact, error) {…}
    4. When server starts (in main) add a call to the mentalfloss to get all parsed facts
    5. Adds all facts to facts store

6. End result of this step - Refactor code to separate all structs into separate packages 
    1. Create a new folder `fact` - move store and fact definition into that folder (change the package name to fact) as well as their methods
    2. Create a new folder `mentalfloss` - move mentalfloss struct  into that folder (change the package name to mentalfloss)        as well as its methods
        1. (optional) If we wish to make it even more abstract it is possible to add a provider interface that has a - func (p            provider) Facts() ([]fact, error) {…} functionality that will enable switching between facts providers ar adding              more than 1 provider easily
    3. Refactor main func
        1. Move the anonymous functions used to register the endpoints to the hendleFunc outside of main function
        2. add imports for our new `fact` and `mentalfloss` packages and use them to make the calls to structs and methods                defined outside the main package

7. End result of this step - every const time a ticker will send a signal to a `thread` (go built-in) that will fetch new fact from provider (mentalfloss)
    1. Init a context.WithCancle (remember to defer its closer…)
    2. Add a function - func updateFactsWithTicker(ctx context.Context, updateFunc func() error)
        1. (Outside from updateFactsWithTicker) Create the updateFunc from step 7.2. that updates the store from an external              provider
        2. (Within the updateFactsWithTicker) Create a time.NewTicker 
        3. (Within the updateFactsWithTicker) Add a go routine with a function that accepts the context
            1. Inside the function add an endless loop that will select from the channel (the ticker channel and the context                  one)
                1. If the ticker channel (ticker.C) was selected - use the given updateFunc to update store
                2. If the context channel (context.Done()) was selected -return (it means the main closed the context)
