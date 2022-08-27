// Sleeping barber problem
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// saloonCapacity >= 1, clientCount >= 1, barberCount >= 1
var saloonCapacity int = 1
var clientCount int = 10
var barberCount int = 1

var cWg sync.WaitGroup
var bWg sync.WaitGroup
var saloon = make(chan int, saloonCapacity)

// sleep is for the output to be nicer, it is not a vital part of the algorithm
func sleep() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
}

func barber(id int) {
	defer bWg.Done()

	for client := range saloon {
		fmt.Printf("Barber%d: cuts Client%d\n", id, client)
		sleep()
	}
}

func client(id int) {
	defer cWg.Done()

	select {
	case saloon <- id:
		fmt.Printf("Client%d: enters saloon\n", id)
	default:
		fmt.Printf("Client%d: saloon is full, leaves\n", id)
	}
}

func main() {
	// launch barber goroutines
	for i := 0; i < barberCount; i++ {
		bWg.Add(1)
		go barber(i + 1)
	}

	// launch client goroutines
	for i := 0; i < clientCount; i++ {
		cWg.Add(1)
		go client(i + 1)
		sleep()
	}

	// wait for all the client goroutines, close channel, wait for barber goroutines
	cWg.Wait()
	close(saloon)
	bWg.Wait()
}
