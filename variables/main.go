package main

import "fmt"

func main() {
	var age int = 30 // explicit variable definition
	name := "Alice"  // implicit variable definition

	isOnline := true

	const maxPlayers = 500

	fmt.Println("Age is", age)
	fmt.Println("Name is", name)
	fmt.Println("IsOnline is", isOnline)
	fmt.Println("Max Players is", maxPlayers)
}
