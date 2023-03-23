package main

import (
	"errors"
	"fmt"
)

var ErrInvalidNode = errors.New("node is not valid")

type Node interface {
	SetValue(v int) error
	GetValue() int
}

//single linked list
type SLLNode struct {
	next         *SLLNode
	value        int
	SNodeMessage string
}

func (sll *SLLNode) SetValue(v int) error {
	if sll == nil {
		return ErrInvalidNode
	}
	sll.value = v
	return nil
}

func (sll *SLLNode) GetValue() int {
	return sll.value
}

func NewSLLNode() *SLLNode {
	return &SLLNode{SNodeMessage: "This is a message from normal node"}
}

//Power linked list
type PowerNode struct {
	next         *PowerNode
	value        int
	PNodeMessage string
}

func (pn *PowerNode) SetValue(v int) error {
	if pn == nil {
		return ErrInvalidNode
	}
	pn.value = v * 10
	return nil
}

func (pn *PowerNode) GetValue() int {
	return pn.value
}

func NewPowerNode() *PowerNode {
	return &PowerNode{PNodeMessage: "This is a message from power node"}
}

func createNode(v int) Node {
	//sn := NewSLLNode()
	sn := NewPowerNode()
	sn.SetValue(v)
	return sn
}

func main() {

	n := createNode(5)

	switch conreten := n.(type) {
	case *SLLNode:
		fmt.Println("Type is SLLNode, message: ", conreten.SNodeMessage)
	case *PowerNode:
		fmt.Println("Type is PowerNode, message: ", conreten.PNodeMessage)
	}
}
