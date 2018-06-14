package main

import "fmt"

type myPrintInterface interface {
	print()
}
type myInterface struct {
}

func (my *myInterface) print() {
	fmt.Println("this is my interface!")
}
func main() {
	var x interface{}
	x = &myInterface{}
	x.(myPrintInterface).print()
}
