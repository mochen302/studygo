package main

import (
	"fmt"
	"time"
)

func sentData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	for {
		fmt.Println("get Data!")
		input = <-ch
		fmt.Printf("%s \n", input)
		//break
	}
}

func main() {
	ch := make(chan string)

	go sentData(ch)
	go getData(ch)

	time.Sleep(1e9)
}
