package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	serverEndpoint = "http://127.0.0.1:9002"

	commandGetAllFacts = "getAllFact"
)

func main() {
	fmt.Println("Hello, Client!")

	cl := NewClient(serverEndpoint)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.Trim(input, "\n ")
		tokens := regexp.MustCompile("[ ]+").Split(input, -1)

		cmd, args := tokens[0], tokens[1:]
		if cmd == "exit" {
			fmt.Println("Bye, bye!")
			return
		}

		res, err := processCmd(cl, cmd, args)
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}

		if res != "" {
			fmt.Println(res)
		}
	}

}

func processCmd(cl *client, cmd string, args []string) (string, error) {
	switch cmd {
	case "":
		return "", nil
	case commandGetAllFacts:
		facts, err := cl.GetAllFacts()
		if err != nil {
			return "", err
		}

		var msg string
		for i, fact := range facts {
			msg += fmt.Sprintf("\n**************\nFact %d:", i)
			msg += fmt.Sprintf("\timage: %s\n\tdescription: %s\n\tcreatedAt: %s", fact.Image, fact.Description)
		}

		return msg, nil
	default:
		return "", errors.New("unknown command")
	}
}
