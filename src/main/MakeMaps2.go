package main

import (
	"fmt"
	"container/list"
)

func main() {

	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		//items[i][1] = 2
	}
	items[0][2] = 4444
	fmt.Printf("Version A: Value of items: %v\n", items)

	list1 := list.New()
	list1.PushFront(1)
	list1.PushFront(2)
	list1.PushFront(3)
	list1.PushFront(4)
	fmt.Println(list1)
}
