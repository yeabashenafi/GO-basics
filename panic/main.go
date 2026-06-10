package main

import "fmt"

func systemCheck() {
	// defer ensures this function runs at the very end of systemCheck
	// even if the function panics midway!
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("🛡️ Recovered from a critical panic:", r)
			fmt.Println("System stabilized safely")
		}
	}()

	fmt.Println("Starting core systems...")

	// Simulating a catastrophic failure
	panic("Reactor core overheating!")

	fmt.Println("This line will never execute")
}
func main() {
	systemCheck()
	fmt.Println("Program continues running smoothly in main!")
}
