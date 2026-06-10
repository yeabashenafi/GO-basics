package main

import "fmt"

// Define the blueprint using a struct
type Gamer struct {
	Username string
	Level    int
}

// attach a method to the struct
func (g Gamer) Greet() {
	fmt.Printf("Player %s logged in! (Level %d)\n", g.Username, g.Level)
}

func main() {
	player1 := Gamer{
		Username: "Player One",
		Level:    2,
	}

	// call the method
	player1.Greet()
}
