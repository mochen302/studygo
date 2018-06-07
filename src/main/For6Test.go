package main

import "fmt"

func main() {
	//LABEL1:
	//	for i := 0; i <= 5; i++ {
	//		for j := 0; j <= 5; j++ {
	//			if j == 4 {
	//				continue LABEL1
	//			}
	//			fmt.Printf("i is: %d, and j is: %d\n", i, j)
	//		}
	//	}
	//
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
