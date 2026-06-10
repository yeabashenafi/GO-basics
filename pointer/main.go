package main

import "fmt"

func main() {
	// create a normal variable
	age := 30

	// create a pointer that points to age's memory location using '&'
	var agePointer *int = &age

	fmt.Println("--- Reading Values ---")
	fmt.Println("Value of age", age)
	fmt.Println("Memory address of age", &age)

	fmt.Println("\n --- Reading the Pointer --- ")
	fmt.Println("Value stored in agePointer", agePointer)
	fmt.Println("Value agePointer stores is", *agePointer) // get the value by using '*' from the pointer

	// change the value using the pointer
	*agePointer = 35

	fmt.Println("\n--- After Modifying via Pointer ---")
	fmt.Println("New value of age:", age)
}
