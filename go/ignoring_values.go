package main

import "fmt"

func main() {

	// _ will actually ignore the firstName
	_, lastName := getNames()

	fmt.Println(lastName)
}

func getNames() (string, string) {
	return "Az", "Dev"
}
