package main

import "fmt"

func main() {
	const premiumPlanName = "Premium"
	const basicPlanName = "basic"

	fmt.Println(premiumPlanName)
	fmt.Println(basicPlanName)

	// COMPUTED CONSTANTS

	const secondsInMinutes = 60
	const minutesInHour = 60

	const secondsInHour = secondsInMinutes * minutesInHour

	fmt.Println(secondsInHour)
}
