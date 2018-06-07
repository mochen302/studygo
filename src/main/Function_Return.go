package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Add2()(3))
	value := Adder(2)(3)
	fmt.Printf(strconv.Itoa(value))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}
func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
