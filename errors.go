package main

import (
	"errors"
	"fmt"
)

// Custom error using struct
type NegativeNumberError struct {
	Number int
}

// Implement the Error() method to satisfy the error interface
func (e NegativeNumberError) Error() string {
	return fmt.Sprintf("negative number not allowed: %d", e.Number)
}

// Function that returns a standard error
func divide(a, b int) (int, error) {
	if b == 0 {
		// Using built-in error creation
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Function that returns a custom error
func squareRoot(n int) (float64, error) {
	if n < 0 {
		// Returning custom error
		return 0, NegativeNumberError{Number: n}
	}
	return float64(n) * 0.5, nil // simplified logic (not actual sqrt)
}

func main() {

	// Example 1: Standard error handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Example 2: Custom error handling
	value, err := squareRoot(-5)
	if err != nil {
		// Type assertion to check custom error
		if customErr, ok := err.(NegativeNumberError); ok {
			fmt.Println("Custom Error:", customErr)
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Value:", value)
	}
}
