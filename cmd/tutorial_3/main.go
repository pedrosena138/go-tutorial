package main

import (
	"errors"
	"fmt"
)

func greetings(name string) {
	fmt.Println("Greetings", name)
}

func intDivision(numerator, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("error: division by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

func main() {
	greetings("John")

	const num int = 11
	const den int = 0
	var result, remainder, err = intDivision(num, den)
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result is: %v", result)
	default:
		fmt.Printf("Division: %v. Remainder: %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("Exact division")
	case 1, 2:
		fmt.Println("Division was close")
	default:
		fmt.Println("Division was not close")
	}

}
