package main

import "fmt"

func main() {

	messageLen := 10
	maxMessageLen := 20

	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message Sent")
	} else {
		fmt.Println("Message not Sent")
	}

	email := "test@email.com"
	if lenght := len(email); lenght >= 10 {
		fmt.Println("the number of byte is greater than 10", len(email))
	} else {
		fmt.Println("the number of byte is is lesser than 10")
	}
}
