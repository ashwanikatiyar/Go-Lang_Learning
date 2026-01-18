package main

import (
	"fmt"
	"sync"
)

func producer(intChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Printf("Sending %d to channel\n", i)
		intChan <- i
	}

	close(intChan) // producer owns the channel
}

func consumer(intChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range intChan {
		fmt.Println("Received value:", value)
	}

	fmt.Println("Channel closed, worker exiting")
}

func main() {
	intChan := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(2) // two goroutines

	go producer(intChan, &wg)
	go consumer(intChan, &wg)

	wg.Wait()
	fmt.Println("All goroutines finished")
}
