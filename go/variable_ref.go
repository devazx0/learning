package main

import "fmt"

func main() {
	sendSoFar := 430
	const sendToAdd int = 25
	sendSoFar = incrementSends(sendSoFar, sendToAdd)

	fmt.Println(sendSoFar)
}

func incrementSends(sendSoFar, sendToAdd int) int {
	return sendSoFar + sendToAdd
}
