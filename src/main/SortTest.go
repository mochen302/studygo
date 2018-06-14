package main

import (
	"sort"
	"fmt"
)

type IntArray []int

func (p IntArray) Len() int { return len(p) }
func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	data := IntArray{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	sort.Sort(data)
	fmt.Println(data)
}
