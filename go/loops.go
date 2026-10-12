package main

import "fmt"

func main() {
	for i := 1; i <= 6; i++ {

		if i%2 == 0 {
			continue
		}

		if i == 5 {
			fmt.Println("Stopping at 5")
			break
		}
		fmt.Println(i)
	}

	users := []string{"Az", "Dev", "Wasiu"} //array of strings

	for i, u := range users {
		fmt.Printf("%s is in index %d \n", u, i)
	}
}
