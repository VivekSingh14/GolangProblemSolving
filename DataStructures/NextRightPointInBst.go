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

	LevelOrderTraversaltest(root)

	fmt.Println("******************")

	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))

	arr := []byte{'a', 'b', 'c'}
	reverseString(arr)

}

func LevelOrderTraversalNew(root *TreeNode1) *TreeNode1 {

	if root == nil {
		return nil
	}

	var que []*TreeNode1
	temp := root
	que = append(que, temp)
	for len(que) > 0 {
		var row []*TreeNode1
		cnt := len(que)
		for cnt > 0 {
			cnt = cnt - 1
			tempdata := que[0]
			que = que[1:]
			row = append(row, tempdata)
			if tempdata.left != nil {
				que = append(que, tempdata.left)
			}
			if tempdata.right != nil {
				que = append(que, tempdata.right)
			}
			for i := 0; i < len(row)-1; i++ {
				row[i].next = row[i+1]
			}
		}
	}
	//fmt.Println(data)

	return root
}

func LevelOrderTraversaltest(root *TreeNode1) *TreeNode1 {

	var que []*TreeNode1
	//var data []int
	temp := root
	que = append(que, temp)
	for len(que) > 0 {
		tempdata := que[0]
		que = que[1:]
		//data = append(data, tempdata.data)
		fmt.Println(tempdata.next)
		if tempdata.left != nil {
			que = append(que, tempdata.left)
		}
		if tempdata.right != nil {
			que = append(que, tempdata.right)
		}
	}
	//fmt.Println(data)

	return nil
}

func containsDuplicate(nums []int) bool {
	map1 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		val, ok := map1[nums[i]]
		if ok {
			val = val + 1
			map1[nums[i]] = val
		} else {
			map1[nums[i]] = 1
		}
	}
	for _, i := range map1 {
		if i > 1 {
			return true
		}
	}
	return false

}

func reverseString(s []byte) {
	max := len(s) - 1
	y := max
	for i := 0; i <= max/2; i++ {
		temp := s[i]
		s[i] = s[y]
		s[y] = temp
		y = y - 1
	}
	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}

}

func maxSlidingWindow(nums []int, k int) []int {

	for i := 0; i < k; i++ {

	}
	return nil
}
