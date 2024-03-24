package main

import "fmt"

// type Contract interface {
// 	InOrder()
// 	PostOrder()
// 	PreOrder()
// }

type Node struct {
	left  *Node
	data  int
	right *Node
}

type BSTtree struct {
	root *Node
}

func (bst *BSTtree) insertIntoTree(data int) {

	if bst.root == nil {
		bst.root = &Node{data: data}
	} else {
		bst.root.insertIntoTree(data)
	}
}

func (node *Node) insertIntoTree(data int) {
	if data <= node.data {
		if node.left == nil {
			node.left = &Node{data: data}
		} else {
			node.left.insertIntoTree(data)
		}
	} else {
		if node.right == nil {
			node.right = &Node{data: data}
		} else {
			node.right.insertIntoTree(data)
		}
	}
}

func inOrder(node *Node) {
	if node == nil {
		return
	} else {
		inOrder(node.left)
		fmt.Print(node.data, "\t")
		inOrder(node.right)
	}
}

func preOrder(node *Node) {
	if node == nil {
		return
	} else {
		fmt.Print(node.data, "\t")
		preOrder(node.left)
		preOrder(node.right)
	}

}

func preOrderItr(node *Node) {

}

func postOrder(node *Node) {
	if node == nil {
		return
	} else {
		postOrder(node.left)
		postOrder(node.right)
		fmt.Print(node.data, "\t")
	}

}

func main() {
	var bst BSTtree

	bst.insertIntoTree(4)
	bst.insertIntoTree(2)
	bst.insertIntoTree(5)
	bst.insertIntoTree(1)
	bst.insertIntoTree(3)
	//bst.insertIntoTree(10)
	bst.insertIntoTree(8)
	bst.insertIntoTree(6)
	inOrder(bst.root)
	fmt.Println()
	preOrder(bst.root)
	fmt.Println()
	postOrder(bst.root)
	fmt.Println()
}
