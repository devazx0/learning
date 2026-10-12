package main

import "fmt"

type Address struct {
	house   int
	city    string
	country string
}

type ContactInfo struct {
	email  string
	number string
}
type User struct {
	Name    string
	age     int
	address Address //nested struct
	ContactInfo
}

// u User is an example of value receiver
func (u User) Greet() {
	fmt.Printf("Hi, my name is %s and I am %d years old \n", u.Name, u.age)
}

// u *User is an example of pointer receiver
func (u *User) DoBirthday() {
	u.age++
}

func (u User) CompareAge(newAge int) int {
	if u.age > newAge {
		return u.age
	} else {
		return newAge
	}

}

func main() {

	// METHOD 1: EASIER TO READ
	u1 := User{
		Name: "Wale", age: 18, address: Address{house: 2, city: "Lagos", country: "Nigeria"}, ContactInfo: ContactInfo{email: "test@email.com", number: "09084297472"},
	}

	u1.Greet()
	u1.DoBirthday()

	fmt.Println(u1.CompareAge(20))

	fmt.Println(u1)
	// METHOD 2: we use the position to assign the values to fields in the struct
	u2 := User{"Tola", 90, Address{}, ContactInfo{}}
	fmt.Println(u2)

	//METHOD 3: both the name and age are inititaed as "" and 0
	u3 := User{}

	u3.age = 30
	u3.Name = "Coder"
	fmt.Println(u3)

	// anonymouse struct - perfect for one time use
	car := struct {
		Brand string
		year  int
	}{
		Brand: "Toyota", year: 2019,
	}
	fmt.Println(car)

}
