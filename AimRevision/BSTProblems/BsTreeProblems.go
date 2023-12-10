package main

import "fmt"

type node1 struct {
	left  *node1
	data  int
	right *node1
}

type BsTree struct {
	root *node1
}

func (t *BsTree) insertIntoTree(data int) {
	if t.root == nil {
		t.root = &node1{data: data}
	} else {
		t.root.insertIntoTree(data)
	}
}
func (node *node1) insertIntoTree(data int) {
	if data <= node.data {
		if node.left == nil {
			node.left = &node1{data: data}
		} else {
			node.left.insertIntoTree(data)
		}
	} else {
		if node.right == nil {
			node.right = &node1{data: data}
		} else {
			node.right.insertIntoTree(data)
		}
	}
}

func printInOrder(n *node1) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%d \t", n.data)
		printInOrder(n.right)
	}
}

func printPreOrder(n *node1) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d \t", n.data)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func main() {
	var t BsTree
	t.insertIntoTree(4)
	t.insertIntoTree(2)
	t.insertIntoTree(5)
	t.insertIntoTree(1)
	t.insertIntoTree(3)
	t.insertIntoTree(7)
	t.insertIntoTree(6)

	printInOrder(t.root)
	fmt.Println()
	printPreOrder(t.root)
}