package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := []int{4, 7, 9, 6, 3, 2, 5, 8, 1}
	inorde := []int{7, 6, 9, 3, 4, 5, 8, 2, 1}
	node := buildTree(preorder, inorde)
	Inorder(node)
	fmt.Println()
	test := Invert(node)
	fmt.Println("---------------------")
	Inorder(test)
	fmt.Println()

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var i int
	for j := 0; j < len(inorder); j++ {
		if inorder[j] == preorder[0] {
			i = j
		}
	}
	lpre, rpre := preorder[1:i+1], preorder[i+1:]
	lin, rin := inorder[:i], inorder[i+1:]

	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(lpre, lin),
		Right: buildTree(rpre, rin),
	}
}

func Inorder(node *TreeNode) {
	if node == nil {
		return
	}
	Inorder(node.Left)
	fmt.Print(node.Val, " ")
	Inorder(node.Right)
}

func Invert(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}
	var queue1 []*TreeNode
	queue1 = append(queue1, node)

	for len(queue1) != 0 {
		node := queue1[0]
		tempo := node.Left
		node.Left = node.Right
		node.Right = tempo
		queue1 = queue1[1:]
		if node.Left != nil {
			queue1 = append(queue1, node.Left)
		}
		if node.Right != nil {
			queue1 = append(queue1, node.Right)
		}
	}
	return node
}

// func LevelOrderTraversal(root *TreeNode) []int {
// 	var queue1 []*TreeNode
// 	var data []int
// 	data = append(data, 0)
// 	temp := root
// 	queue1 = append(queue1, temp)

// 	for len(queue1) != 0 {

// 		temp = queue1[0]
// 		queue1 = queue1[1:]
// 		//for displaying
// 		//fmt.Print(temp.data, " ")
// 		data = append(data, temp.data)
// 		if temp.left != nil || temp.right != nil {
// 			queue1 = append(queue1, temp.left)
// 			queue1 = append(queue1, temp.right)
// 		}

// 	}
// 	return data
// }
