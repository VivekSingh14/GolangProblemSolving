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

	if fast.next == nil {
		return slow.data
	}

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

func (ll *LinkedList) RemoveNthFromEnd(target int) {
	count := 0

	temp := ll.head

	for temp != nil {
		count++
		temp = temp.next
	}
	temp = ll.head
	if count <= 1 {
		return
	} else if count == target {
		ll.head = ll.head.next
	} else {
		//fmt.Println("count: ", count)

		for i := 0; i < count-(target+1); i++ {
			temp = temp.next
		}
		temp.next = temp.next.next
	}
}

func (ll *LinkedList) RemoveNthFromEndOpt(target int) {

	fast := ll.head
	slow := ll.head

	for i := 0; i < target; i++ {
		fast = fast.next
	}

	if fast == nil {
		fast = ll.head.next
		ll.head = fast
		return
	}

	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	slow.next = slow.next.next

}

func (ll *LinkedList) CyclicList() bool {

	fast := ll.head.next
	slow := ll.head

	for fast != nil {
		if fast.data == slow.data {
			return true
		}
		fast = fast.next.next
		slow = slow.next
	}
	return false

	//or other method with o(1) space complexity
	// we will try to update the data every time and try to find out whether this same data comes back again
	// for slow != nil {
	// 	if slow.data == math.MaxInt {
	// 		return true
	// 	}

	// 	slow.data = math.MaxInt
	// 	slow = slow.next
	// }
	// return false
}

func main() {

	ll := LinkedList{}

	ll.Insert(2)
	ll.Insert(4)
	ll.Insert(6)
	ll.Insert(7)
	ll.Insert(9)
	//ll.Insert(11)
	//ll.Insert(14)

	//ll.Iterate()
	//fmt.Println(ll.MiddleNode())

	//ll.Reverse()

	//ll.Iterate()

	//ll.RemoveNthFromEnd(2)

	//ll.Iterate()

	//ll.RemoveNthFromEndOpt(2)

	ll.Iterate()

}
