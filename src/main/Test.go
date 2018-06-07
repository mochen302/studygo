package main

import (
	"fmt"
	"strings"
)

var a1 = "G"

func main() {
	n()
	m()
	n()

	fmt.Println(strings.Replace("this is a test \n es", "t", "a", -1))
}

func n() { print(a1) }

func m() {
	a1 := "O"
	print(a1)
}
