package main

import (
	"fmt"
	"net/http"

	"github.com/FTBpro/go-workshop/coolfacts/cfhttp"
)

func main() {
	fmt.Println("Hello, World!")

	server := cfhttp.NewServer()
	if err := http.ListenAndServe(":9002", server); err != nil {
		panic(err)
	}
}
