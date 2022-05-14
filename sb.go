// Sleeping barber problem
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Saloon capacity - (excluding the cutting chair), number of clients, number of barbers
var saloonCapacity int = 1
var clientCount int = 10
var barberCount int = 1

// Clients wait group, barbers wait group, saloon channel of set capacity
var cWg sync.WaitGroup
var bWg sync.WaitGroup
var saloon = make(chan int, saloonCapacity)

// Work simulation for output observability
func sleep() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
}

// Barber function
func barber(id int) {
	defer bWg.Done()

	for client := range saloon {
		fmt.Printf("Barber%d: cuts Client%d\n", id, client)
		sleep()
	}
}

// Client function
func client(id int) {
	defer cWg.Done()

	select {
	case saloon <- id:
		fmt.Printf("\t\t\tClient%d: enters saloon\n", id)
	default:
		fmt.Printf("\t\t\tClient%d: saloon is full, leaves\n", id)
	}
}

func main() {

	// Output header
	fmt.Println("BARBER\t\t\tCLIENTS")
	fmt.Println(strings.Repeat("- ", 30))

	// Launch barber goroutines
	for i := 0; i < barberCount; i++ {
		bWg.Add(1)
		go barber(i + 1)
	}

	// Launch client goroutines
	for i := 0; i < clientCount; i++ {
		cWg.Add(1)
		go client(i + 1)
		sleep()
	}

	// Wait for all the client goroutines, close channel, wait for barber
	cWg.Wait()
	close(saloon)
	bWg.Wait()
}
