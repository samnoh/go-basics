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
	if koreanAge := age + 1; koreanAge < 18 { // variable expression
		return false
	}
	return true
}

type person struct {
	name    string
	age     int
	favFood []string
}

// entry point
func main() {
	fmt.Println(multiply(2, 6))

	length, uppercase := lengthAndUpperCase("hello world")
	fmt.Println(length, uppercase)

	repeatWords("hello", "world", "bye", "world")

	fmt.Println(addAllNumbers(1, 2, 3, 4, 5))

	fmt.Println(canIDrink(17))

	// pointers
	a := 2
	b := a
	a = 10
	fmt.Println(a, b, &a, &b) // 10 2 0xc0000180e8 0xc0000180f0

	c := 2
	d := &c
	c = 5
	*d = 20
	fmt.Println(c, &c, d, &d, *d) // 20 0xc000018108 0xc000018108 0xc00000e030 20

	// arrays
	names := [5]string{"h", "e", "l", "l"}
	names[4] = "o"
	// names[5] = "b" // erorr; out of bound

	// slices
	otherNames := []string{"h", "e", "l", "l"}
	otherNames = append(otherNames, "o")

	me := map[string]string{"name": "sam", "age": "18"}
	for key, value := range me {
		fmt.Println(key, value) // name same; age 18
	}

	favFood := []string{"pizza", "chicken"}
	sam := person{name: "sam", age: 18, favFood: favFood}
	fmt.Println(sam)
}
