package main

import (
	"fmt"
)

// func concat(s1 string, s2 string) string {

// s1 and s2 will be considered strings since both of them are strings
func concat(s1, s2 string) string {

	return s1 + " " + s2
}

func fullName(firstName, lastName string) string {
	return firstName + " " + lastName
}

func checkEligibility(nameFormatter func(string, string) string, firstName, lastName string, age int) bool {
	if nameFormatter(firstName, lastName) != "" && age >= 18 {
		return true
	}
	return false

}

func main() {
	fmt.Println(concat("Hello", "world"))

	eligibility := checkEligibility(fullName, "Az", "Dev", 8)
	fmt.Println((eligibility))
}
