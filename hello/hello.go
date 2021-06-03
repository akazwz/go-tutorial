package main

import (
	"fmt"
	"github.com/akazwz/go-tutorial/greetings"
)

func main() {
	message := greetings.Hello("Zhao")
	fmt.Println(message)
}

