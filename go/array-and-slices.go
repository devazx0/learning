package main

import "fmt"

func modifyFirst(s []string) {
	s[0] = "MUTATED"
}

func main() {

	var numbers [3]int

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	evens := [6]int{2, 4, 6, 8, 10, 12}

	s := evens[1:4] //slicing from 1 to 4 (4 exclusive)

	fmt.Println(len(evens))
	fmt.Println(s)

	names := []string{"Alice", "Alice", "Gen"}

	modifyFirst(names) //sclices are reference to their backing array, so if we change something in the slice it changes in the array too

	fmt.Println(names)

	newNumbers := make([]int, 3, 10)
	fmt.Println(newNumbers)
	fmt.Println(len(newNumbers)) //4 because we've appended one more item
	newNumbers = append(newNumbers, 1)

	//even though 3 is the length of this slight but the cap is 10 so we can use append function to grow it
	fmt.Println(newNumbers)
	fmt.Println(cap(newNumbers))
	fmt.Println(len(newNumbers)) //4 because we've appended one more item

}
