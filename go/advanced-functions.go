package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

//HIGHER ORDER FUNCTION
// 1. ACCEPT ONE OR MORE FUNCTION AS A PARAM
// 2. RETURNS A FUNCTION

func double(n int) int {
	return n * 2
}

func processNumbers(numbers []int, transformer func(int) int) []int {

	result := make([]int, len(numbers))

	for i, n := range numbers {
		result[i] = transformer(n)
	}

	return result

}

func main() {
	myOperation := add
	fmt.Println(myOperation(10, 20))

	// anon function

	// 1. Assigning an anonymous function directly to a variable
	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Println(multiply(10, 20))

	// 2. Immediately Invoked Function Expression (IIFE)
	// The trailing parentheses () call the function on the spot

	func(msg string) {
		fmt.Println("Message", msg)
	}("Hello, GO!!")

	nums := []int{1, 2, 3, 4}
	doubled := processNumbers(nums, double)
	fmt.Println(doubled)

	squared := processNumbers(nums, func(n int) int {
		return n * n
	})

	fmt.Println("squared", squared)

}
