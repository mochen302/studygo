package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 5
	fmt.Println(&num)

	var intP = &num
	fmt.Println(*intP)

	var p1 *int = nil
	*p1 = 0
	fmt.Println(p1)

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Println(s + ";" + *p)

	s1 := 2
	var p2 *int = &s1
	*p2 = 3
	fmt.Println(strconv.Itoa(s1) + strconv.Itoa(*p2))

}
