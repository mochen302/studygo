package main

import (
	"time"
	"fmt"
)

func main() {
	//sendChan := make(chan int)//会死锁
	sendChan := make(chan int, 10000)
	receiveChan := make(chan string)
	build(sendChan)
	go processChannel(sendChan, receiveChan)
	go consume(receiveChan)
	time.Sleep(5 * 1e9)
}
func build(dataChan chan<- int) {
	for i := 0; i < 10000; i++ {
		dataChan <- i
	}
}

func processChannel(in <-chan int, out chan<- string) {
	result := 0
	for inValue := range in {
		result += inValue
	}
	out <- string(result)
}

func consume(ch <-chan string) {
	for v := range ch {
		fmt.Println(v)
	}
}
