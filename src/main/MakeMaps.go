package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.111
	mapAssigned["key3"] = 3

	delete(mapAssigned, "key2")
	fmt.Printf("Map literal at \"one\" is: %d\n", mapAssigned["key2"])

	//fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	//fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	//fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	//fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	mf := map[int]func() int{
		1: func() int {
			return 10
		},
		2: func() int {
			return 20
		},
	}
	fmt.Println(mf[1]())

	if _, ok := mf[2]; ok {
		fmt.Println("存在")
	}

}
