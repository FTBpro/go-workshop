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

func main() {
	fmt.Println("Hello, World!")

	cl := NewClient("http://127.0.0.1:9002")

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

		res, err := processCmd(cl, command(cmd), args)
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}

		if res != "" {
			fmt.Println(res)
		}
	}

}

const (
	createFactCommand  command = "createFact"
	commandGetLastFact command = "getLastFact"
)

type command string

func processCmd(cl *client, cmd command, args []string) (string, error) {
	switch cmd {
	case "":
		return "", nil
	case createFactCommand:
		if len(args) < 2 {
			return "", errors.New("invalid arguments")
		}

		fct := coolfact.Fact{
			Image:       args[0],
			Description: strings.Join(args[1:], " "),
		}

		err := cl.CreateFact(fct)
		return "", err
	case commandGetLastFact:
		lastFact, err := cl.GetLastCreatedFact()
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("\timage: %s\n\tdescription: %s\n\tcreatedAt: %s", lastFact.Image, lastFact.Description, lastFact.CreateAt), nil

	default:
		return "", errors.New("unknown command")
	}
}
