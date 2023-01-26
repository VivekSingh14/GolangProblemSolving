package main

import "fmt"

type TreeNode1 struct {
	left  *TreeNode1
	data  int
	right *TreeNode1
	next  *TreeNode1
}

func Insert1(root *TreeNode1, key int) *TreeNode1 {
	newNode := TreeNode1{nil, key, nil, nil}
	if root == nil {
		root = &newNode
		return root
	}
	if key < root.data {
		root.left = Insert1(root.left, key)

	} else {
		root.right = Insert1(root.right, key)
	}
	return root
}

func main() {
	var root *TreeNode1

	root = Insert1(root, 4)
	root = Insert1(root, 2)
	root = Insert1(root, 6)
	root = Insert1(root, 3)
	root = Insert1(root, 1)
	root = Insert1(root, 5)
	root = Insert1(root, 7)

	fmt.Println(root)
}

func LevelOrderTraversalNew(root *TreeNode1) *TreeNode1 {
	return nil
}
