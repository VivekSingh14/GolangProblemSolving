package main

import "fmt"

type Node1 struct {
	left  *Node1
	data  int
	right *Node1
}

type BsTree struct {
	root *Node1
}

func (t *BsTree) InsertIntoTree(data int) {
	if t.root == nil {
		t.root = &Node1{data: data}
	} else {
		t.root.InsertIntoTree(data)
	}
}

func (node *Node1) InsertIntoTree(data int) {
	if data <= node.data {
		if node.left == nil {
			node.left = &Node1{data: data}
		} else {
			node.left.InsertIntoTree(data)
		}
	} else {
		if node.right == nil {
			node.right = &Node1{data: data}
		} else {
			node.right.InsertIntoTree(data)
		}
	}
}

func printPreOrder(n *Node1) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d \t", n.data)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printInOrder(n *Node1) {
	if n == nil {
		return
	} else {
		printPreOrder(n.left)
		fmt.Printf("%d \t", n.data)
		printPreOrder(n.right)
	}
}

func main() {
	var t BsTree
	t.InsertIntoTree(6)
	t.InsertIntoTree(2)
	t.InsertIntoTree(1)
	t.InsertIntoTree(4)
	t.InsertIntoTree(3)
	t.InsertIntoTree(5)
	t.InsertIntoTree(7)
	//t.InsertIntoTree(2)

	fmt.Println("PreOrder.....")
	printPreOrder(t.root)

	fmt.Println("\nInOrder.....")
	printInOrder(t.root)
}
