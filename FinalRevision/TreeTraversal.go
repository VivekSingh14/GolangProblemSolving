package main

import "fmt"

type Node struct {
	left  *Node
	data  int
	right *Node
}

type BstItr struct {
	root *Node
}

func (t *BstItr) InsertIntoTree1(data int) {
	if t.root == nil {
		t.root = &Node{data: data}
	} else {
		t.root.InsertIntoTree1(data)
	}
}

func (node *Node) InsertIntoTree1(data int) {
	if data <= node.data {
		if node.left == nil {
			node.left = &Node{data: data}
		} else {
			node.left.InsertIntoTree1(data)
		}
	} else {
		if node.right == nil {
			node.right = &Node{data: data}
		} else {
			node.right.InsertIntoTree1(data)
		}
	}
}

func preorderTraversal(root *Node) []int {

	var arr []int
	var stack []*Node
	temp := root
	stack = append(stack, temp)
	for len(stack) > 0 {
		for temp != nil {
			arr = append(arr, temp.data)
			if temp.right != nil {
				stack = append(stack, temp.right)
			}
			temp = temp.left
		}
		temp = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

	}
	return arr
}

func main() {

	var t BstItr
	t.InsertIntoTree1(4)
	t.InsertIntoTree1(2)
	t.InsertIntoTree1(6)
	t.InsertIntoTree1(3)
	t.InsertIntoTree1(1)
	t.InsertIntoTree1(5)
	t.InsertIntoTree1(7)

	// arr := inorderTraversal(t.root)
	// fmt.Println(arr)

	arr1 := preorderTraversal(t.root)
	fmt.Println(arr1)

	// arr2 := postorderTraversal(t.root)
	// fmt.Println(arr2)

}
