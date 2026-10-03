//Barrier.go Template Code
//Copyright (C) 2026 Terry Huynh

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Terry Huynh (C00300806)
// Created: 14/09/2026
// Modified by: Terry
// Issues:
// Helped by: C00259228, Shadrach (C00298390) and Isabel (C00303465)
// Helped: Shadrach (C00298390) and Isabel (C00303465)
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived *int, max int, wg *sync.WaitGroup, sharedLock *sync.Mutex, theChan chan bool) bool {
	for round := 0; round < 5; round++ {
		time.Sleep(time.Second)
		fmt.Println("Part A", goNum)
		sharedLock.Lock()
		*arrived++

		//we wait here until everyone has completed part A
		if *arrived == max { //last to arrive -signal others to go
			sharedLock.Unlock() //unlock before any potentially blocking code
			theChan <- true
			<-theChan
		} else { //not all here yet we wait until signal
			sharedLock.Unlock() //unlock before any potentially blocking code
			<-theChan
			theChan <- true //once we get through send signal to next routine to continue
		} //end of if-else
		sharedLock.Lock()
		*arrived--
		sharedLock.Unlock()
		fmt.Println("PartB", goNum)
	}
	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 10
	arrived := 0
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	var theLock sync.Mutex
	theChan := make(chan bool)     //use unbuffered channel in place of semaphore
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &arrived, totalRoutines, &wg, &theLock, theChan)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
