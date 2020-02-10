package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a + b
}

func lengthAndUpperCase(name string) (length int, uppercase string) {
	defer fmt.Println("I am done")

	length = len(name)
	uppercase = strings.ToUpper(name)
	fmt.Println("I am not done")
	return
}

func repeatWords(words ...string) {
	fmt.Println(words)
}

// entry point
func main() {
	fmt.Println(multiply(2, 6))

	length, uppercase := lengthAndUpperCase("hello world")
	fmt.Println(length, uppercase)

	repeatWords("hello", "world", "bye", "world")
}
