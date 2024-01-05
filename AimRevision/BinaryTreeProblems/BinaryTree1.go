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
	return nil
}
