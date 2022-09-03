package main

import "fmt"

type NodeSk struct {
	data int
	next *NodeSk
}

type StackImpl struct {
	head *NodeSk
}

func (s *StackImpl) Push(data int) {
	if s.head == nil {
		s.head = &NodeSk{data: data}
	} else {
		s.head.Push(data)
	}
}

func (n *NodeSk) Push(data int) {

	temp := n
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &NodeSk{data: data}

}

func (s *StackImpl) Traverse() {
	temp := s.head

	for temp.next != nil {
		fmt.Print(temp.data, "\t")
		temp = temp.next
	}
}

func main() {
	var stack1 StackImpl
	stack1.Push(3)
	stack1.Push(4)
	stack1.Push(5)
	stack1.Push(6)

	stack1.Traverse()
}