package main

import (
	"fmt"
	"time"
)

// A slow function that simulates downloading a large file
func downloadFile(channel chan string) {
	fmt.Println("🛰️  Download started in the background...")

	// simulate taking 3 seconds to download
	time.Sleep(3 * time.Second)

	// send a success message into the channel using the arrow operator '<-'
	channel <- "Download complete! (Movie.mp4)"

}

func main() {
	// Create a channel that can transfer strings
	dataChannel := make(chan string)

	startTime := time.Now() // Track when we started

	// launch the function in the background using 'go'
	go downloadFile(dataChannel)

	// The main program keeps moving instantly
	fmt.Println("🏠 Main program is free to do other work...")

	for i := 1; i <= 5; i++ {
		fmt.Printf("Loading UI animation... Frame %d\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("⏳ OK, now the main program will wait for the channel data...")

	// Receive data from the channel
	// This line "blocks" (pauses main) until the background worker sends something
	result := <-dataChannel

	fmt.Println("Success !", result)
	fmt.Printf("⏱️ Total program execution time: %v\n", time.Since(startTime))
}
