package main

import "fmt"

type TwoInts struct {
	a, b int
}

type TwoInts1 struct {
	a, b int
}

func main() {
	two1 := &TwoInts{1, 2}
	two2 := &TwoInts1{4, 5}

	fmt.Println(two1.AddThem())
	fmt.Println(two2.AddToParam(2))

}
func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts1) AddToParam(param int) int {
	return tn.a + param
}
