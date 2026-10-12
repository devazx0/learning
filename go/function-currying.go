package main

import "fmt"

func makeMultiplier(factor int) func(int) int {
	return func(val int) int {
		return factor * val
	}
}
func main() {

	double := makeMultiplier(2)
	tripple := makeMultiplier(3)

	fmt.Println(double(10))
	fmt.Println(tripple(10))
}
