package main

import "fmt"

type TreeNode struct {
	left  *TreeNode
	data  int
	right *TreeNode
}

func Insert(root *TreeNode, key int) *TreeNode {
	newNode := TreeNode{nil, key, nil}
	if root == nil {
		root = &newNode
		return root
	}
	if key < root.data {
		root.left = Insert(root.left, key)

	} else {
		root.right = Insert(root.right, key)
	}
	return root
}

func InorderTree(root *TreeNode) {
	temp := root

	if temp == nil {
		return
	}

	InorderTree(temp.left)
	fmt.Print(temp.data, " \t")
	InorderTree(temp.right)
}

func PreOrder(root *TreeNode) {
	temp := root

	if temp == nil {
		return
	}

	fmt.Print(temp.data, " \t")
	PreOrder(temp.left)
	PreOrder(temp.right)
}

func PostOrder(root *TreeNode) {
	temp := root

	if temp == nil {
		return
	}

	PostOrder(temp.left)
	PostOrder(temp.right)
	fmt.Print(temp.data, " \t")
}

func LeafNode(root *TreeNode) {
	temp := root

	if temp == nil {
		return
	}
	LeafNode(temp.left)
	if temp.left == nil && temp.right == nil {
		fmt.Println(temp.data)
	}
	LeafNode(temp.right)

}

func main() {
	var root *TreeNode

	root = Insert(root, 4)
	root = Insert(root, 2)
	root = Insert(root, 5)
	root = Insert(root, 1)
	root = Insert(root, 3)
	root = Insert(root, 10)
	root = Insert(root, 8)
	root = Insert(root, 6)

	PreOrder(root)
	fmt.Println()
	InorderTree(root)
	fmt.Println()
	PostOrder(root)
	fmt.Println()
	fmt.Println("---------------")
	LeafNode(root)

}
