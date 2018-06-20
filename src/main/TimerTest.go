package main

import (
	"time"
	"fmt"
	"sync"
)

func main() {

	lock := new(sync.Mutex)

	lock.Lock()
	tick := time.Tick(1e9)
	boom := time.After(5e9)
	lock.Unlock()

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom.")
		default:
			//fmt.Println("    .")
			//time.Sleep(5e7)
		}
	}

}
