//--------------------------------------------
// Author: Terry Huynh (C00300806)
// Helped by: C00259228, Shadrach (C00298390) and Isabel (C00303465)
// Helped: Shadrach (C00298390) and Isabel (C00303465)
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

//Global variables shared between functions --A BAD IDEA

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool)

	doStuffOne := func() bool {
		fmt.Println("StuffOne - Part A")
		//wait here
		barrier <- true
		fmt.Println("StuffOne - PartB")
		wg.Done()
		return true
	}
	doStuffTwo := func() bool {
		time.Sleep(time.Second * 5)
		fmt.Println("StuffTwo - Part A")
		//wait here

		<-barrier
		fmt.Println("StuffTwo - PartB")
		wg.Done()
		return true
	}
	wg.Add(2)
	go doStuffOne()
	go doStuffTwo()
	wg.Wait() //wait here until everyone (10 go routines) is done

}
