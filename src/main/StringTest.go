package main

import (
	"strings"
	"fmt"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	s1 := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", s1)

	for index, val := range s1 {
		fmt.Printf("%d - %s", index, val)
	}
	fmt.Println()
	fmt.Println(`----------------------------`)

	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	fmt.Println(`----------------------------`)

	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)


}
