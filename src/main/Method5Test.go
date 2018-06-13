package main

import "fmt"

type Base struct {
}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

//func (self Voodoo) MoreMagic() {
//	self.Magic()
//	self.Magic()
//}

func main() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
