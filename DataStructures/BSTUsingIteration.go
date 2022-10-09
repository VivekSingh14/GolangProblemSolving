package main

import "fmt"

type Node2 struct {
	left  *Node2
	data  int
	right *Node2
}

//added comment
type BstItr struct {
	root *Node2
}

func (t *BstItr) InsertIntoTree1(data int) {
	if t.root == nil {
		t.root = &Node2{data: data}
	} else {
		t.root.InsertIntoTree1(data)
	}
}

func (node *Node2) InsertIntoTree1(data int) {
	if data <= node.data {
		if node.left == nil {
			node.left = &Node2{data: data}
		} else {
			node.left.InsertIntoTree1(data)
		}
	} else {
		if node.right == nil {
			node.right = &Node2{data: data}
		} else {
			node.right.InsertIntoTree1(data)
		}
	}
}

func inorderTraversal(root *Node2) []int {
	var stack []*Node2
	var result []int
	curr := root
	for curr != nil || len(stack) != 0 {
		// go left as far as possible
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}
		// add to result and go right
		curr = stack[len(stack)-1]   //fetching value
		stack = stack[:len(stack)-1] //stack pop
		result = append(result, curr.data)
		curr = curr.right
	}

	return result
}

func preorderTraversal(root *Node2) []int {

	var stack []*Node2
	var result []int

	curr := root
	stack = append(stack, curr)

	for len(stack) != 0 {

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.data)

		if curr.right != nil {
			stack = append(stack, curr.right)
		}
		if curr.left != nil {
			stack = append(stack, curr.left)
		}
	}
	return result
}

func postorderTraversal(root *Node2) []int {

	var stack1 []*Node2
	var stack2 []*Node2
	var result []int

	//curr := root
	stack1 = append(stack1, root)
	for len(stack1) != 0 {

		temp := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, temp)
		if temp.left != nil {
			stack1 = append(stack1, temp.left)
		}
		if temp.right != nil {
			stack1 = append(stack1, temp.right)
		}

	}
	j := len(stack2) - 1
	for j >= 0 {
		curr := stack2[j]
		result = append(result, curr.data)
		j--
	}
	return result
}

func mai3() {

	var t BstItr
	t.InsertIntoTree1(4)
	t.InsertIntoTree1(2)
	t.InsertIntoTree1(6)
	t.InsertIntoTree1(3)
	t.InsertIntoTree1(1)
	t.InsertIntoTree1(5)
	t.InsertIntoTree1(7)

	arr := inorderTraversal(t.root)
	fmt.Println(arr)

	arr1 := preorderTraversal(t.root)
	fmt.Println(arr1)

	arr2 := postorderTraversal(t.root)
	fmt.Println(arr2)

}
