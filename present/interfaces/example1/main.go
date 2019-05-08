package main

import (
	"fmt"
)

// start OMIT
type Speaker interface {
	Speak()
}

func Speak(s Speaker) {
	s.Speak()
}

// end OMIT

type speaker struct{}

func (s speaker) Speak() {
	fmt.Println("Hello, this is speaker")
}

func main() {
	s := speaker{}
	Speak(s)
}
