package main

import (
	"fmt"

	"github.com/make-go-great/ioe-go"
)

func main() {
	fmt.Println("Input something:")
	input := ioe.ReadInput()
	fmt.Printf("You input: %s\n", input)

	fmt.Println("Input again:")
	input = ioe.ReadInputEmpty()
	fmt.Printf("You input: %s\n", input)
}
