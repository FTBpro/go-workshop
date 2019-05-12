package main

// start OMIT

import (
	"log"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Welcome to my website!"))
}

func main() {
	http.HandleFunc("/hello", helloWorldHandler)
	log.Println("Listen and Serve on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// end OMIT
