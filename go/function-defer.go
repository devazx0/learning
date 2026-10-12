package main

import "fmt"

func processOrder() {
	fmt.Println("1. Processing order start")

	// Scheduled to execute when processOrder() finishes
	defer fmt.Println("4. Deferred cleanup: Order fully processed")
	defer fmt.Println("3. Deferred cleanup: Sub-step complete")

	fmt.Println("2. Writing invoice to database")
}

func main() {

	processOrder()

}
