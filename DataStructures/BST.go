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
		printInOrder(n.left)
		fmt.Printf("%d \t", n.data)
		printInOrder(n.right)
	}
}

func printPostOrder(n *Node1) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%d \t", n.data)
	}
}

func searchNode(n *Node1, data int) {
	if n == nil {
		fmt.Println("Node not found.")
		return
	}
	if n.data == data {
		fmt.Println("Found it.")
	}
	if data > n.data {
		searchNode(n.right, data)
	} else {
		searchNode(n.left, data)
	}

}

func main() {
	var t BsTree
	t.InsertIntoTree(4)
	t.InsertIntoTree(2)
	t.InsertIntoTree(5)
	t.InsertIntoTree(1)
	t.InsertIntoTree(3)
	t.InsertIntoTree(10)
	t.InsertIntoTree(8)
	t.InsertIntoTree(6)
	//t.InsertIntoTree(5)
	//t.InsertIntoTree(7)
	//t.InsertIntoTree(2)

	fmt.Println("PreOrder.....")
	printPreOrder(t.root)

	fmt.Println("\nInOrder.....")
	printInOrder(t.root)

	fmt.Println("\nPostOrder.....")
	printPostOrder(t.root)

	fmt.Println("\nSearch in tree........")
	searchNode(t.root, 14)
}
