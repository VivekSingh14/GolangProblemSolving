package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Insert(data int) {
	newnode := Node{data, nil}

	if ll.head == nil {
		ll.head = &newnode
		return
	}
	ptr := ll.head

	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = &newnode
}

func (ll *LinkedList) Iterate() {
	ptr := ll.head

	for ptr != nil {
		fmt.Print(ptr.data, " ")
		ptr = ptr.next
	}
	fmt.Println()
}

func (ll *LinkedList) MiddleNode() int {
	fast := ll.head.next
	slow := ll.head

	for fast != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow.data
}

func (ll *LinkedList) Reverse() {
	var prev *Node
	curr := ll.head
	var nex *Node

	for curr != nil {
		nex = curr.next
		curr.next = prev
		prev = curr
		curr = nex
	}
	ll.head = prev
}

func main() {

	ll := LinkedList{}

	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(5)
	ll.Insert(7)
	ll.Insert(9)
	ll.Insert(11)
	ll.Insert(14)

	ll.Iterate()
	fmt.Println(ll.MiddleNode())

	ll.Reverse()

	ll.Iterate()

}
