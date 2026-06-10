package main

import "fmt"

func main() {
	// A standard count-to-five loop
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// while loop style loop
	count := 1
	for count <= 4 {
		fmt.Println("Count is", count)
		count++
	}
}
