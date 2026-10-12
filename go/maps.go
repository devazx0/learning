package main

import "fmt"

func checkKeyExist(m map[string]int, key string) {
	i, ok := m[key]

	if !ok {
		fmt.Println("Key doesn't exist")
		return
	}

	fmt.Printf("Key does exist %d", i)
}

type FullName struct {
	First, Last string
}

func main() {

	// ages := make(map[string]int)

	// ages["John"] = 37
	// ages["Sarah"] = 29

	//OR

	ages := map[string]int{
		"John":  37,
		"Sarah": 29,
	}
	fmt.Println(ages) // map[John:37 Sarah:29]

	ages["Alex"] = 21 // Inserts "Alex"

	delete(ages, "Sarah")

	fmt.Println(ages)
	fmt.Println(len(ages))

	checkKeyExist(ages, "Sarah")

	// NESTED MAPS VS STRUCT KEYS

	nestedMap := make(map[string]map[string]int)

	nestedMap["name"] = map[string]int{"Az": 12}

	userScores := make(map[FullName]int)

	userScores[FullName{First: "A", Last: "b"}] = 90
	userScores[FullName{First: "Q", Last: "O"}] = 89

	fmt.Println(userScores)

}
