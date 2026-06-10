package main

import (
	"errors"
	"fmt"
)

// returns a float 64 and an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// using the errors value o create a new error value
		return 0, errors.New("can not divide by a zero denominator")
	}

	// If successfull, we return the result and 'nil' for the error
	return a / b, nil
}

func main() {
	// call the function and catch both the result and the error
	result, err := divide(3, 0)

	// check if error is not nil
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}

	// This line only runs if err was nil
	fmt.Printf("Success! The result is: %.2f\n", result)
}
