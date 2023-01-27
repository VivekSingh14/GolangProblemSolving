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

	//fmt.Println(root)
	LevelOrderTraversalNew(root)
}

func LevelOrderTraversalNew(root *TreeNode1) *TreeNode1 {

	var que []*TreeNode1
	var data []int
	temp := root
	que = append(que, temp)
	for len(que) > 0 {
		tempdata := que[0]
		que = que[1:]
		data = append(data, tempdata.data)
		if tempdata.left != nil {
			que = append(que, tempdata.left)
		}
		if tempdata.right != nil {
			que = append(que, tempdata.right)
		}
	}
	fmt.Println(data)

	return nil
}
