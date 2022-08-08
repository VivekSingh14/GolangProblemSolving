package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(data int) {
	n := Node{}
	n.data = data
	if l.head == nil {
		l.head = &n
		//l.len++
		return
	}
	ptr := l.head
	for ptr != nil {
		if ptr.next == nil {
			ptr.next = &n
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) Display() {
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

func (l *LinkedList) DeleteAtStart() {
	if l.head == nil {
		fmt.Println("No elements to be deleted.")
	}

	ptr := l.head

	fmt.Println(ptr.data, " will be deleted.")
	ptr = ptr.next
	l.head = ptr
}

// func (l *LinkedList) DeleteAtPosition(pos int) {
// 	if l.len < pos {
// 		fmt.Println("Invalid linked list")
// 	}
// 	ptr := l.head
// 	for i := 1; i < pos-1; i++ {
// 		ptr = ptr.next
// 	}
// 	ptr.next = ptr.next.next
// 	l.len--
// }

func (l *LinkedList) Reverse() *Node {
	var prev *Node
	curr := l.head
	var forward *Node
	for curr != nil {
		forward = curr.next
		curr.next = prev
		prev = curr
		curr = forward
	}
	fmt.Println(prev.data)
	fmt.Println(prev.next.data)
	return prev
}

func main1() {
	ll1 := LinkedList{}
	ll1.Insert(2)
	ll1.Insert(3)
	ll1.Insert(4)
	ll1.Insert(5)
	ll1.Insert(6)
	ll1.Insert(7)
	fmt.Println("============Before deletion.============")
	ll1.Display()

	ll1.DeleteAtStart()
	fmt.Println("============After deletion.============")
	ll1.Display()
	//
	//ll1.DeleteAtPosition(2)
	//fmt.Println("============After deletion.============")
	//ll1.Display()

	// ll := ll1.Reverse()
	// ll1.head = ll
	// fmt.Println("============After reverse.============")
	// ll1.Display()
}
