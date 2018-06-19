package main

import "fmt"

func main() {
	primes := sieve0()
	for {
		fmt.Println(<-primes)
	}
}

func generate0() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter0(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve0() chan int {
	out := make(chan int)
	go func() {
		ch := generate0()
		for {
			prime := <-ch
			ch = filter0(ch, prime)
			out <- prime
		}
	}()
	return out
}
