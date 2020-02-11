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

func addAllNumbers(numbers ...int) int {
	total := 0

	for _, number := range numbers { // index, number
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 1; koreanAge < 18 {
		return false
	}
	return true
}

// entry point
func main() {
	fmt.Println(multiply(2, 6))

	length, uppercase := lengthAndUpperCase("hello world")
	fmt.Println(length, uppercase)

	repeatWords("hello", "world", "bye", "world")

	fmt.Println(addAllNumbers(1, 2, 3, 4, 5))

	fmt.Println(canIDrink(17))
}
