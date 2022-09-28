package main

import (
	"fmt"
	"sort"

	"golang.org/x/exp/slices"
)

type Node1 struct {
	left  *Node1
	data  int
	right *Node1
}

type BsTree struct {
	root *Node1
}

func (t *BsTree) InsertIntoTree(data int) {
	if t.root == nil {
		t.root = &Node1{data: data}
	} else {
		t.root.InsertIntoTree(data)
	}
}

func (node *Node1) InsertIntoTree(data int) {
	if data <= node.data {
		if node.left == nil {
			node.left = &Node1{data: data}
		} else {
			node.left.InsertIntoTree(data)
		}
	} else {
		if node.right == nil {
			node.right = &Node1{data: data}
		} else {
			node.right.InsertIntoTree(data)
		}
	}
}

func printPreOrder(n *Node1) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d \t", n.data)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printInOrder(n *Node1) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%d \t", n.data)
		printInOrder(n.right)
	}
}

func printPostOrder(n *Node1) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%d \t", n.data)
	}
}

func searchNode(n *Node1, data int) {
	if n == nil {
		fmt.Println("Node not found.")
		return
	}
	if n.data == data {
		fmt.Println("Found it.")
	} else if data > n.data {
		searchNode(n.right, data)
	} else {
		searchNode(n.left, data)
	}

}

func deleteNode(root *Node1, data int) *Node1 {

	var parent *Node1
	curr := root

	for curr != nil && curr.data != data {
		parent = curr

		if data < curr.data {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	// if to be deleted node is root node
	if curr == nil {
		return root
	}
	//if to be deleted node is leaf node
	if curr.left == nil && curr.right == nil {
		if curr != root {
			if parent.left == curr {
				parent.left = nil
			} else {
				parent.right = nil
			}
		} else {
			root = nil
		}
		// if the node to be deleted, having both child
	} else if curr.left != nil && curr.right != nil {
		minNode := getMinKey(curr.right)

		val := minNode.data

		deleteNode(root, val)

		curr.data = val
	} else {
		var child *Node1
		if curr.left != nil {
			child = curr.left
		} else {
			child = curr.right
		}

		if curr != root {
			if curr == parent.left {
				parent.left = child
			} else {
				parent.right = child
			}
		} else {
			root = child
		}
	}
	return root
}

func getMinKey(tempNode *Node1) *Node1 {
	for tempNode.left != nil {
		tempNode = tempNode.left
	}
	return tempNode
}

func main() {
	var t BsTree
	t.InsertIntoTree(4)
	t.InsertIntoTree(2)
	t.InsertIntoTree(5)
	t.InsertIntoTree(1)
	t.InsertIntoTree(3)
	t.InsertIntoTree(10)
	t.InsertIntoTree(8)
	t.InsertIntoTree(6)
	//t.InsertIntoTree(5)
	//t.InsertIntoTree(7)
	//t.InsertIntoTree(2)

	fmt.Println("PreOrder.....")
	printPreOrder(t.root)

	fmt.Println("\nInOrder.....")
	printInOrder(t.root)

	fmt.Println("\nPostOrder.....")
	printPostOrder(t.root)

	fmt.Println("\nSearch in tree........")
	searchNode(t.root, 10)

	fmt.Println("\ndelete in tree........")
	temp := deleteNode(t.root, 10)

	fmt.Println("\nInOrder.....")
	printInOrder(temp)

	// arr := make([]int, 8)
	// arr = append(arr, 4)
	// arr = append(arr, 2)
	// arr = append(arr, 5)
	// arr = append(arr, 1)
	// arr = append(arr, 3)
	// arr = append(arr, 10)
	// arr = append(arr, 8)
	// arr = append(arr, 6)

	arr := []int{8, 4, 2, 5, 1, 3, 10}

	newtree := convert(arr)
	fmt.Println("\nInOrder..... after convert.........")
	printInOrder(newtree)

	fmt.Println("\nIs BST.....")
	fmt.Println(isBst(newtree))

	fmt.Println("\nIn Order Predecessor.....")
	fmt.Println(inOrderPredecessor(newtree, 3))

	fmt.Println("\nMin and Max in .....")
	fmt.Println(findMaxMin(newtree))

	fmt.Println("\nK th Max .....")
	fmt.Println(findKthMax(newtree, 2))

	fmt.Println("\nPair with given sum .....")
	PairWithGivenSum(newtree, 15)

	fmt.Println("\nHeight or depth of tree .....")
	fmt.Println(heightOrDepthOfTree(newtree))

}

func convert(arr []int) *Node1 {
	sort.Ints(arr)
	temp := convert1(arr, 0, len(arr)-1)
	return temp
}

func convert1(arr []int, low int, high int) *Node1 {
	if low > high {
		return nil
	}

	mid := (low + high) / 2

	root := Node1{nil, arr[mid], nil}

	root.left = convert1(arr, low, mid-1)
	root.right = convert1(arr, mid+1, high)

	return &root
}

//not a perfect solution
func isBst(node *Node1) bool {
	if node == nil {
		return true
	}
	if node.left != nil && node.left.data > node.data {
		return false
	}

	if node.right != nil && node.right.data < node.data {
		return false
	}

	if !isBst(node.left) || !isBst(node.right) {
		return false
	}

	return true
}

func inOrderPredecessor(root *Node1, key int) int {
	var predec int
	if root != nil {
		if root.data == key {
			if root.left != nil {
				t := root.left

				for t.right != nil {
					t = t.right
				}
				predec = t.data
			}
			return predec
		} else if root.data < key {
			node := root.right
			inOrderPredecessor(node, key)
		}
	}
	return 0
}

func findMaxMin(root *Node1) (int, int) {

	node := root

	for node.left != nil {
		node = node.left
	}
	min := node.data

	node = root
	for node.right != nil {
		node = node.right
	}
	max := node.data
	return min, max
}

func findKthMax(root *Node1, k int) int {
	store := storeInOrder(root)
	storeCap := len(store)
	return store[storeCap-k]
}

func storeInOrder(n *Node1) []int {
	var stack []*Node1
	var result []int
	curr := n
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

func PairWithGivenSum(root *Node1, sum int) {

	var stack []*Node1
	slice1 := make([]int, 4)
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		tempSum := sum - curr.data
		//if tempSum !=
		if slices.Contains(slice1, tempSum) {
			fmt.Println(curr.data, " : ", tempSum)
			return
		} else {
			slice1 = append(slice1, curr.data)
		}
		curr = curr.right
	}
}

func heightOrDepthOfTree(root *Node1) int {
	var x, y int
	temp := root
	if temp != nil {
		x = heightOrDepthOfTree(temp.left)
		y = heightOrDepthOfTree(temp.right)

		if x > y {
			return x + 1
		} else {
			return y + 1
		}
	}
	return 0
}
