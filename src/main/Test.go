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

func n() { consume(a1) }

func m() {
	a1 := "O"
	consume(a1)
}
