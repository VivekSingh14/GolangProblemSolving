package main

import "fmt"

//queue operations....
type Contract interface {
	Enqueue(int)
	Dequeue()
	DisplayQue()
}

type Node1 struct {
	data int
	next *Node
}

type CustomQueue struct {
	start *Node
}

func (que *CustomQueue) Enqueue(data int) {
	newNode := Node{data, nil}
	if que.start == nil {
		que.start = &newNode
		return
	}

	temp := que.start

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &newNode
}

func (que *CustomQueue) Dequeue() {
	if que.start == nil {
		return
	}
	que.start = que.start.next
}

func (que *CustomQueue) DisplayQue() {

	temp := que.start

	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()

}

func main1() {
	newQue := CustomQueue{}
	newQue.Enqueue(2)
	newQue.Enqueue(3)
	newQue.Enqueue(4)
	newQue.Enqueue(5)

	newQue.DisplayQue()

	newQue.Dequeue()

	newQue.DisplayQue()

}
