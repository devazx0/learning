package main

import "fmt"

func main() {

	// declaring multiple variables in one single line
	averageOpenRate, displayMessage := 0.23, "is the average rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)
}
