package main

import (
	"sort"
	"fmt"
)

type Element interface{}

type Vector struct {
	a []Element
	b int
}

func (p *Vector) At(i int) Element {
	return p.a[i]
}

func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}

func (p Vector) Len() int { return len(p.a) }
func (p Vector) Less(i, j int) bool {
	ei := p.a[i]
	ej := p.a[j]
	if pi, ri := ei.(int); ri {
		if pj, rj := ej.(int); rj {
			return pi < pj
		}
	}
	return false
}
func (p Vector) Swap(i, j int) {
	p.a[i], p.a[j] = p.a[j], p.a[i]
}

func main() {
	vector := Vector{make([]Element, 8), 2}
	//vector := Vector{[]Element{6, 7, 7}}
	fmt.Println(vector)
	vector.Set(0, 6)
	vector.Set(1, 4)
	vector.Set(2, 5)
	sort.Sort(vector)
	fmt.Println(vector)
}
