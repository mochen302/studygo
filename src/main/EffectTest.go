package main

import "fmt"

func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println(n)
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}
