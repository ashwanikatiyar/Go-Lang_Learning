package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Create a buffered channel (the queue) with capacity for 5 emails
	emailQueue := make(chan string, 5)

	// 2. Start a "worker" goroutine to process (print) emails from the queue
	go func() {
		// 'range' will read emails one by one until the channel is closed
		for email := range emailQueue {
			fmt.Printf("Sending email to: %s\n", email)
			// Simulate processing time
			time.Sleep(100 * time.Millisecond) 
		}
		fmt.Println("All emails processed. Worker shutting down.")
	}()

	// 3. Enqueue 10 emails
	for i := 1; i <= 10; i++ {
		email := fmt.Sprintf("%d@gmail.com", i)
		fmt.Printf("Enqueued: %s\n", email)
		emailQueue <- email // This only blocks if the buffer (5) is full
	}

	// 4. Close the channel to signal the worker that no more emails are coming
	close(emailQueue)

	// Wait briefly to allow the worker to finish printing the last few emails
	time.Sleep(2 * time.Second)
}
