package main

import "fmt"

type node struct {
	data int
	next *node
}

type LList struct {
	head *node
}

func (ll *LList) insert(data int) {
	newNode := node{data, nil}
	if ll.head == nil {
		ll.head = &newNode
		return
	}
	ptr := ll.head
	for ptr != nil {
		if ptr.next == nil {
			ptr.next = &newNode
			return
		}
		ptr = ptr.next
	}
}

func (l *LList) display() {
	if l.head == nil {
		fmt.Println("0 Elements in the linkedList")
	}
	ptr := l.head
	for ptr != nil {
		fmt.Print(ptr.data, "\t")
		ptr = ptr.next
	}
	fmt.Println()
}

func (l *LList) Reverse() *node {
	var prev *node
	curr := l.head
	var forward *node

	for curr != nil {
		forward = curr.next
		curr.next = prev
		prev = curr
		curr = forward
	}

	return prev
}

func main() {

	ll1 := LList{}
	ll1.insert(1)
	ll1.insert(2)
	ll1.insert(3)
	ll1.insert(4)
	ll1.insert(5)
	ll1.insert(6)

	ll1.display()

	rev := ll1.Reverse()

	ll1.head = rev

	ll1.display()

}
