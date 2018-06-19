package main

import (
	"time"
	"fmt"
)

func main() {
	//sendChan := make(chan int)//会死锁
	sendChan := make(chan int)
	receiveChan := make(chan string)
	go processChannel(sendChan, receiveChan) //开启一个协程
	go consume(receiveChan)
	build(sendChan)
	time.Sleep(5 * 1e9)
}
func build(dataChan chan<- int) {
	for i := 0; i < 10000; i++ {
		dataChan <- i
	}
}
func processChannel(in <-chan int, out chan<- string) {
	for inValue := range in {
		out <- string(inValue)
	}
}

func consume(ch <-chan string) {
	for v := range ch {
		fmt.Println(v)
	}
}
