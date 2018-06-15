package main

import "fmt"

func f1(in chan int) {
	fmt.Println(<-in)
}
func main() {
	//会死锁!!!
	out := make(chan int)
	out <- 2
	go f1(out)
}
