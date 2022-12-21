package main

import "fmt"

//added hd variable because we want to find out horizontal distance between nodes
type TreeNode struct {
	left  *TreeNode
	data  int
	hd    int
	right *TreeNode
}

func Insert(root *TreeNode, key int) *TreeNode {
	newNode := TreeNode{nil, key, 0, nil}
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
	fmt.Println("-------Inorder--------")
	InOrderUsingIteration(root)
	fmt.Println("\n-------PreOrder--------")
	PreOrderUsingIteration(root)
	fmt.Println("\n-------PostOrder--------")
	PostOrderUsingIteration(root)
	fmt.Println("\n-------Height of tree: --------")
	fmt.Println(HeightOfTree(root))
	fmt.Println("\n-------Level Order--------")
	arr := LevelOrderTraversal(root)
	fmt.Println(arr)
	fmt.Println("\n-------Right View of tree--------")
	RightViewOfTree(root)
	fmt.Println("\n-------Sum of leaf Nodes--------")
	SumOfLeaf(root)
	fmt.Println("\n-------Cousins of given node--------")
	cousins := FindCousin(arr, 5)
	fmt.Println(cousins)
	fmt.Println("\n---------Bottom view of tree------------")
	BottomViewOfTree(root)
	fmt.Println("\n---------Top view of tree------------")
	TopViewOfTree(root)
	fmt.Println("\n---------Max depth or height using level order traversal and queue------------")
	fmt.Println(MaxDepthOrHeight(root))

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

func PostOrderUsingIteration(root *TreeNode) {
	var stack1 []*TreeNode
	var stack2 []*TreeNode

	curr := root
	stack1 = append(stack1, curr)
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
		fmt.Print(curr.data, " ")
		j--
	}

}

func HeightOfTree(root *TreeNode) int {
	var x, y int
	temp := root
	if temp != nil {
		x = HeightOfTree(temp.left)
		y = HeightOfTree(temp.right)

		if x > y {
			return x + 1
		} else {
			return y + 1
		}
	}
	return 0
}

func LevelOrderTraversal(root *TreeNode) []int {
	var queue1 []*TreeNode
	var data []int
	data = append(data, 0)
	temp := root
	queue1 = append(queue1, temp)

	for len(queue1) != 0 {

		temp = queue1[0]
		queue1 = queue1[1:]
		//for displaying
		//fmt.Print(temp.data, " ")
		data = append(data, temp.data)
		if temp.left != nil || temp.right != nil {
			queue1 = append(queue1, temp.left)
			queue1 = append(queue1, temp.right)
		}

	}
	return data
}

func SumOfLeaf(root *TreeNode) {
	var stack []*TreeNode
	curr := root
	var sum int
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}
		curr = stack[len(stack)-1]   //fetching value
		stack = stack[:len(stack)-1] //stack pop
		if curr.left == nil && curr.right == nil {
			sum = sum + curr.data
		}
		curr = curr.right
	}
	fmt.Println(sum)
}

func FindCousin(tree []int, data int) []int {
	var cousins []int
	for i := 1; i < len(tree)/2; i++ {
		//loop will go till len/2 because leftchild and rightchild can go before that only, for proof check below 2 statements
		leftchild := tree[2*i]
		rightchild := tree[2*i+1]
		if (data == leftchild || data == rightchild) && i > 1 {
			//now we will find grand parent of data
			grandparentPosition := i / 2
			var childOfGrandParent int
			if data > tree[grandparentPosition] {
				childOfGrandParent = (2 * grandparentPosition)
			} else {
				childOfGrandParent = (2 * grandparentPosition) + 1
			}
			cousins = append(cousins, tree[childOfGrandParent*2])
			cousins = append(cousins, tree[(childOfGrandParent*2)+1])

		}
	}
	//finding the parent of given node
	return cousins
}

func RightViewOfTree(root *TreeNode) {
	var que []*TreeNode
	temp := root

	que = append(que, temp)

	for len(que) != 0 {
		size := len(que)
		for i := 0; i < size; {
			i++
			curr := que[0]
			que = que[1:]
			if i == size {
				fmt.Print(curr.data, " ")
			}
			if curr.left != nil {
				que = append(que, curr.left)
			}
			if curr.right != nil {
				que = append(que, curr.right)
			}

		}
	}
}

func BottomViewOfTree(root *TreeNode) {
	var queue2 []*TreeNode
	map1 := make(map[int]int)
	temp := root

	queue2 = append(queue2, temp)

	for len(queue2) != 0 {

		temp = queue2[0]
		queue2 = queue2[1:]
		tempHd := temp.hd

		map1[tempHd] = temp.data
		if temp.left != nil {
			temp.left.hd = tempHd - 1
			queue2 = append(queue2, temp.left)
		}
		if temp.right != nil {
			temp.right.hd = tempHd + 1
			queue2 = append(queue2, temp.right)
		}

	}
	for _, element := range map1 {
		fmt.Print(" ", element, " ")
	}
	fmt.Println()
}

func TopViewOfTree(root *TreeNode) {
	var queue2 []*TreeNode
	map1 := make(map[int]int)
	temp := root

	queue2 = append(queue2, temp)

	for len(queue2) != 0 {

		temp = queue2[0]
		queue2 = queue2[1:]
		tempHd := temp.hd
		//main condition to check if top view will be udpated or not
		if _, ok := map1[tempHd]; !ok {
			map1[tempHd] = temp.data
		}
		if temp.left != nil {
			temp.left.hd = tempHd - 1
			queue2 = append(queue2, temp.left)
		}
		if temp.right != nil {
			temp.right.hd = tempHd + 1
			queue2 = append(queue2, temp.right)
		}

	}
	for _, element := range map1 {
		fmt.Print(" ", element, " ")
	}
	fmt.Println()
}

func MaxDepthOrHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var que []*TreeNode
	temp := root
	depth := 0
	que = append(que, temp)
	que = append(que, nil)
	for len(que) != 0 {

		temp = que[0]
		que = que[1:]
		if temp == nil {
			depth = depth + 1
		}
		if temp != nil {

			if temp.left != nil {
				que = append(que, temp.left)

			}
			if temp.right != nil {
				que = append(que, temp.right)
			}
		} else if len(que) != 0 {
			que = append(que, nil)
		}

	}
	return depth
}
