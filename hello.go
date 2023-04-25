package main

import (
	"fmt"
	"log"

	"github.com/liutianq/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("\n")

	fmt.Println(quote.Go())

	fmt.Println("\n")

	var message string
	var err error
	message, err = greetings.Hello("Agnese")
	fmt.Println(message)

	fmt.Println("\n")

	log.SetPrefix("greetings: ")
	log.SetFlags(17)
	message, err = greetings.Hello("Elizabeth")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	fmt.Println("\n")

	names := []string{"A", "Gladys", "C", "Samantha", "Darrin"}

	messages, err, seq := greetings.Hellos(names)
	if err != nil {
		fmt.Println(fmt.Sprintf("The first empty seq#: %v !\n", seq))
		log.Fatal(err)
	}
	//fmt.Println(messages)
	for _, message = range messages {
		fmt.Println(message)
	}
}
