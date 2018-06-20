package main

import (
	"fmt"
)

func main() {
	//runtime.GOMAXPROCS(4)

	ch := make(chan string)
	go sendData0(ch)
	getData0(ch)
}

func sendData0(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData0(ch chan string) {
	for input := range ch {
		fmt.Printf("%s \n", input)
	}
	//for {
	//	input, open := <-ch
	//	if !open {
	//		break
	//	}
	//	fmt.Printf("%s \n", input)
	//}
}
