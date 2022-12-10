package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

const (
	serverEndpoint = "http://127.0.0.1:9002"

	commandGetFacts    = "getFacts"
	createFactCommand  = "createFact"
	commandGetLastFact = "getLastFact"
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
	case commandGetFacts:
		facts, err := cl.GetFacts()
		if err != nil {
			return "", err
		}

		var msg string
		for i, fact := range facts {
			msg += fmt.Sprintf("\n**************\nFact %d:", i)
			msg += fmt.Sprintf("\tTopic: %s\n\tDescription: %s\n\tCreatedAt: %s", fact.Topic, fact.Description, fact.CreatedAt)
		}

		return msg, nil
	case commandGetLastFact:
		lastFact, err := cl.GetLastCreatedFact()
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("\tTopic: %s\n\tDescription: %s\n\tCreatedAt: %s", lastFact.Topic, lastFact.Description, lastFact.CreatedAt), nil
	case createFactCommand:
		if len(args) < 2 {
			return "", errors.New("invalid arguments")
		}

		fct := coolfact.Fact{
			Topic:       args[0],
			Description: strings.Join(args[1:], " "),
		}

		err := cl.CreateFact(fct)
		return "", err

	default:
		return "", errors.New("unknown command")
	}
}
