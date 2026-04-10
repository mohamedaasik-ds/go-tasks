package main

import "fmt"

/**
 * No return type
 */
func printMessage() {
	fmt.Println("This function has no return")
}

/**
 * Single return value
 */
func getSquare(number int) int {
	return number * number
}

/**
 * Multiple return values
 */
func getSumAndProduct(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

/**
 * Named return values
 */
func getQuotientAndRemainder(a int, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

/**
 * Function returning boolean
 */
func isEven(number int) bool {
	return number%2 == 0
}

/**
 * Function returning string
 */
func getGreeting(name string) string {
	return "Hello " + name
}

func main() {

	printMessage()

	square := getSquare(4)
	fmt.Println("Square:", square)

	sum, product := getSumAndProduct(3, 5)
	fmt.Println("Sum:", sum, "Product:", product)

	q, r := getQuotientAndRemainder(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	evenCheck := isEven(6)
	fmt.Println("Is Even:", evenCheck)

	greeting := getGreeting("Aasik")
	fmt.Println(greeting)
}
