package main

import "fmt"

func main() {

	const name = "Alex Goodman"
	const openRate = 30.5

	// %s means we are expecting a string and %f for float (%.1f for example means we want only 1 decimal place)
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f", name, openRate)

	fmt.Println(msg)
}
