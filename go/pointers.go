package main

import "fmt"

func incrementVal(val int) {
	val++
}

func incrementPointer(val *int) {
	*val++
}

//pointer safety
func printMessage(msg *string) {

	//checking if the variable is pointing to a valid address
	if msg == nil {
		return
	}
	fmt.Println(msg)
}

type Car struct {
	Color string
}

func (c Car) PaintWrong(color string) {
	c.Color = color
}

func (c *Car) Paint(color string) {
	c.Color = color
}

func main() {

	// x := 5
	// y := &x //z is a pointer storing the address of y

	// fmt.Println(y)

	//DEREFERCE OPERATOR

	x := 5          //declearing a variable
	var z *int = &x //specifying the type of the underlying data in the pointer *int or *string and &z stores the mem addres of var x

	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(*z) //by doing *z we can print the underlying value of the data stored in the address the pointer is pointing to

	*z = 6 // now since z is pointing to x, now x is now 6

	fmt.Println(x) //6

	count := 5

	incrementVal(count)
	fmt.Println(count) //still 5

	incrementPointer(&count)
	fmt.Println(count) //now 6

	myCar := Car{Color: "white"}
	myCar.PaintWrong("blue")

	fmt.Println("After PaintWrong:", myCar.Color) // White

	myCar.Paint("red")
	fmt.Println("After Paint:", myCar.Color) // red

}
