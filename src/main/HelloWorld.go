package main

import (
	fm "fmt"
	"runtime"
	"strconv"
	"os"
	"../util"
)

type IZ int

const (
	a = iota
	b = 4
	c = iota
)

const (
	Unknow = 0
	Female = 1
	Male   = 2
)

func init() {
	fm.Println("init")
}

func main() {
	fm.Println(a)
	fm.Println(b)
	fm.Println(c)

	fm.Println(os.Getenv("HOME"))
	fm.Println(os.Getenv("USER"))
	fm.Println(os.Getenv("GOROOT"))

	fm.Printf("hello%s", runtime.Version()+"world")
	fm.Println("-------------------")
	//fm.Println(rand.Random())

	result, _, code := add(1, 2)
	fm.Println("" + strconv.Itoa(result) + "code:" + code)

	fm.Println("-------------------")
	path1, path2 := getCurrent()
	fm.Println(path1 + "," + path2)

	fm.Println("-------------------")
	num, info := util.Float2Int(35.6)
	fm.Println(info + strconv.Itoa(num))
}

func add(a, b int) (i int, i1 int, string2 string) {
	return a + b, Female, "this is a test!"
}

func getCurrent() (string, string) {
	return runtime.GOOS, os.Getenv("PATH")
}
