package main

import "fmt"

type square struct {
	side float32
}

func (s square) Area() float32 {
	return s.side * s.side
}

type circle struct {
	radius float32
}
type shaper interface {
	Area() float32
}

func main() {
	var areaIntf shaper
	sq1 := new(square)
	sq1.side = 5

	areaIntf = sq1

	if t, ok := areaIntf.(*square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}

	switch t := areaIntf.(type) {
	case square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	}
}

//func classifier(items ...interface{}) {
//	for i, x := range items {
//
//	}
//
//}
