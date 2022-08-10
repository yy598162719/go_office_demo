package main

import (
	"fmt"
	"github.com/yy598162719/publish_demo"
)

func main() {
	// Get a greeting message and print it.
	message := publish_demo.Hello("Gladys")
	fmt.Println(message)
}
