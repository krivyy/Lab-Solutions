//atomic.go Template Code
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

// --------------------------------------------
// Author: Terry Huynh (C00300806)
// Created: 14/09/2026
// Modified by: Terry
// Issues:
// Helped by: C00259228, Shadrach (C00298390) and Isabel (C00303465)
// Helped: Shadrach (C00298390) and Isabel (C00303465)
// --------------------------------------------
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Global variables shared between functions --A BAD IDEA
var wg sync.WaitGroup

func addsAtomic(n int, total *atomic.Int64) bool {
	for i := 0; i < n; i++ {
		total.Add(1)
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

func main() {

	var total atomic.Int64

	//for loop using range option
	for i := range 10 {
		//the waitgroup is used as a barrier
		// init it to number of go routines
		wg.Add(1)
		fmt.Println("go Routine ", i)
		go addsAtomic(1000, &total)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println(total.Load())

}
