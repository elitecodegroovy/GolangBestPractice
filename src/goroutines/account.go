package goroutines

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func ShowAccount() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)
	fmt.Println("Start Goroutines")
	//launch a goroutine with label "A"
	go printCounts("--A")
	//launch a goroutine with label "B"
	go printCounts("---B")
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

}

func ShowChannelInfo() {
	count := make(chan int)
	// Add a count of two, one for each goroutine.
	wg.Add(2)
	fmt.Println("Start Goroutines")
	//launch a goroutine with label "A"
	go PrintChannel("A", count)
	//launch a goroutine with label "B"
	go PrintChannel("B", count)
	fmt.Println("Channel begin")
	count <- 1
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")

	//channel
	messages := make(chan string, 2)
	messages <- "Golang"
	messages <- "Gopher"
	//Recieve value from buffered channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println("\nTerminating Program")
}

func PrintChannel(label string, count chan int) {
	// Schedule the call to WaitGroup's Done to tell we are done.
	defer wg.Done()
	for {
		//Receives message from Channel
		val, ok := <-count
		if !ok {
			fmt.Println("Channel was closed")
			return
		}
		fmt.Printf("Count: %d received from %s \n", val, label)
		if val == 10 {
			fmt.Printf("Channel Closed from %s \n", label)
			// Close the channel
			close(count)
			return
		}
		val++
		// Send count back to the other goroutine.
		count <- val
	}
}

func printCounts(label string) {
	// Schedule the call to WaitGroup's Done to tell we are done.
	defer wg.Done()
	// Randomly wait
	for count := 1; count <= 10; count++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Count: %d from %s\n", count, label)
	}
}
