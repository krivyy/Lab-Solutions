// --------------------------------------------
// Author: Terry Huynh (C00300806)
// Helped by: C00259228, Shadrach (C00298390) and Isabel (C00303465)
// Helped: Shadrach (C00298390) and Isabel (C00303465)
// --------------------------------------------
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, barrier chan bool, release chan bool) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	barrier <- true
	//Rendezvous here
	<-release
	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool)
	release := make(chan bool)
	threadCount := 10 // Changed to 10 because of the comment under

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, barrier, release)
	}
	for range threadCount {
		<-barrier
	}
	for range threadCount {
		release <- true
	}
	wg.Wait() //wait here until everyone (10 go routines) is done

}
