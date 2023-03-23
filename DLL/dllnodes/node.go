package dllnodes

import "fmt"

type DLLNode struct {
	prev *DLLNode
	data int
	next *DLLNode
}

func NewDLLNode(data int) *DLLNode {
	return &DLLNode{nil, data, nil}
}

type DoubleLinkedList struct {
	head *DLLNode
}

func NewDLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{head: nil}
}

func (dll *DoubleLinkedList) Insert(data int) {
	newnode := NewDLLNode(data)

	if dll.head == nil {
		dll.head = newnode
		return
	}
	temp := dll.head

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newnode
	newnode.prev = temp

}

func (dll *DoubleLinkedList) Traverse() {
	temp := dll.head
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()
}
