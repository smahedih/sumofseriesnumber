package main

import (
	"fmt"
	"os"
)

func main() {

	var startNumber, endNumber int

	fmt.Printf("Enter the starting number: \n")
	fmt.Scan(&startNumber)

	fmt.Printf("Enter the ending number: \n")
	fmt.Scan(&endNumber)

	if startNumber > endNumber {
		fmt.Printf("Invalid input. The starting number should be smaller than the ending number.")
		os.Exit(1)
	}

	fmt.Printf("The series is: ")

	for i := startNumber; i <= endNumber; i++ {
		fmt.Printf("%v ", i)
	}

	sum := 0

	for i := startNumber; i <= endNumber; i++ {
		sum = sum + i

	}

	fmt.Printf("\n")
	fmt.Printf("Sum of the series is: %v", sum)
}
