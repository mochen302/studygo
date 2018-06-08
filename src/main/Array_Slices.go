package main

import (
	"fmt"
	"regexp"
	"io/ioutil"
	"strconv"
)

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	//fmt.Printf(make([]int, 50, 100))
	//new([100]int)[0:50]

	fmt.Println(FindFileDigits("resources/text.txt"))

}

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindFileDigits(filename string) []int {
	fileBytes, _ := ioutil.ReadFile(filename)
	b := digitRegexp.FindAll(fileBytes, len(fileBytes))
	c := make([]int, 0)
	for _, bytes := range b {
		num, _ := strconv.Atoi(string(bytes))
		c = append(c, num)
	}
	return c
}
