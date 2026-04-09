package main

import "fmt"

func main() {
	// variable to test switch
	number := 2

	// switch checks value of number
	switch number {
	case 1:
		fmt.Println("Number is One")

	case 2:
		fmt.Println("Number is Two")

	case 3:
		fmt.Println("Number is Three")

	default:
		// runs if none of the above cases match
		fmt.Println("Number is something else")
	}
}