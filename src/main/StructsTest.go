package main

import "fmt"

type struts1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struts1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)

	ms1 := struts1{10, 15.5, "mochen"}
	fmt.Println(ms1)

	ms2 := &struts1{10, 15.5, "mochen"}
	fmt.Println(ms2)

}
