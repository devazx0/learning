package main

import "fmt"

func aggregator() func(int) int {
	total := 0

	return func(amount int) int {
		total += amount
		return total
	}

}

func main() {

	trackA := aggregator()
	trackB := aggregator()

	fmt.Println("A:", trackA(10)) // 10
	fmt.Println("A:", trackA(15)) // 25 (remembers previous 10)

	// trackB has its own isolated instance of 'total'
	fmt.Println("B:", trackB(50)) // 50
	fmt.Println("A:", trackA(5))  // 30
}
