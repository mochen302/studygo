package main

import "fmt"

func main() {
	var num int = 100

	switch {
	case num >= 90:
		fmt.Println(num)
		fallthrough
	case num < 90:
		fmt.Println(-num)
	}

	switch result := calculate(); {
	case result < 0:
		fmt.Println()
	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

}
func calculate() (int) {
	return 1
}
