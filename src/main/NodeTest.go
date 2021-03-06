package main

import "fmt"

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func main() {
	root := NewNode(nil, nil)
	root.SetData("root node")

	a := NewNode(nil, nil)
	a.SetData("leftNode")

	b := NewNode(nil, nil)
	b.SetData("right node")

	root.le = a
	root.ri = b

	fmt.Printf("%v\n", root)
}
