package main

import (
	"reflect"
	"fmt"
)

type TagType struct {
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}
func refTag(tagType TagType, i int) {
	ttType := reflect.TypeOf(tagType)
	ixField := ttType.Field(i)
	fmt.Printf("%v\n", ixField.Tag)
}
