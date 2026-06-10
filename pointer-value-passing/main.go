package main

import "fmt"

// This takes a COPY. The original variable won't change.
func failToTrain(level int) {
	level = level + 1
}

// This takes a POINTER. It modifies the actual variable directly at its address.
func successfullyTrain(level *int) {
	*level = *level + 1 // Use '*' to modify the value inside the address
}

func main() {
	playerLevel := 10

	failToTrain(playerLevel)
	fmt.Println("Level after failed training:", playerLevel) // Still 10

	// We pass the memory address using '&'
	successfullyTrain(&playerLevel)
	fmt.Println("Level after real training:", playerLevel) // Now it's 11!
}
