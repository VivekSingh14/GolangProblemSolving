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

func PreOrderTree(root *TreeNode) {
	temp := root

	if temp == nil {
		return
	}

	fmt.Print(temp.data, " \t")
	PreOrderTree(temp.left)
	PreOrderTree(temp.right)
}

func PostOrderTree(root *TreeNode) {
	temp := root

	if temp == nil {
		return
	}

	PostOrderTree(temp.left)
	PostOrderTree(temp.right)
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
	root = Insert(root, 6)
	root = Insert(root, 3)
	root = Insert(root, 1)
	root = Insert(root, 5)
	root = Insert(root, 7)
	//root = Insert(root, 6)

	PreOrderTree(root)
	fmt.Println("PreOrder")
	InorderTree(root)
	fmt.Println("InOrder")
	PostOrderTree(root)
	fmt.Println("PostOrder")
	fmt.Println("---------------")
	LeafNode(root)
	fmt.Println("---------------")
	InOrderUsingIteration(root)
	fmt.Println("\n---------------")
	PreOrderUsingIteration(root)

}

func InOrderUsingIteration(root *TreeNode) {
	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}
		curr = stack[len(stack)-1]   //fetching value
		stack = stack[:len(stack)-1] //stack pop
		fmt.Print(curr.data, "  ")
		curr = curr.right
	}

}

func PreOrderUsingIteration(root *TreeNode) {
	var stack []*TreeNode
	curr := root

	stack = append(stack, curr)

	for len(stack) != 0 {
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(curr.data, " ")
		if curr.right != nil {
			stack = append(stack, curr.right)
		}

		if curr.left != nil {
			stack = append(stack, curr.left)
		}
	}

}
