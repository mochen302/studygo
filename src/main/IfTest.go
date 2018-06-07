package main

import "fmt"

func main() {
	if value := 2; value < 5 {
		fmt.Println(value)
	} else if (value > 4) {
		fmt.Println("error")
	}
}
