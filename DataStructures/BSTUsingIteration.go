package main

import "fmt"

type Node2 struct {
	left  *Node2
	data  int
	right *Node2
}

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

// func inOrderTraversal(temp *BstItr) {
// 	var stack []*Node2
// 	//var res []int
// 	cur := temp.root

// 	for cur != nil {
// 		for cur != nil {
// 			stack = append(stack, cur)
// 			cur = cur.left
// 		}
// 		te := stack[len(stack)-1]
// 		fmt.Print(te.data, "\t")
// 		cur = cur.right
// 	}
// }

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
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.data)
		curr = curr.right
	}

	return result
}

// func preorderTraversal(root *Node2) []int {

// }

func main() {

	var t BstItr
	t.InsertIntoTree1(4)
	t.InsertIntoTree1(2)
	t.InsertIntoTree1(5)
	t.InsertIntoTree1(1)
	t.InsertIntoTree1(3)

	arr := inorderTraversal(t.root)
	fmt.Println(arr)

}
