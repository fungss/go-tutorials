package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	// message, err := greetings.Hello("Ronald")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	names := []string{"Adam", "Ben", "Chris"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
