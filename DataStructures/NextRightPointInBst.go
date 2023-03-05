package main

import (
	"fmt"
	"math"
)

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

	fmt.Println("********reverse integer**********")
	fmt.Println(reverse(-356))

	fmt.Println("********digits to number**********")
	nums1 := []int{9, 9, 9, 9}
	fmt.Println(plusOne(nums1))

	fmt.Println("********strings are anagram **********")
	str1 := ""
	str2 := "l"
	fmt.Println(isAnagram(str1, str2))

	fmt.Println("******** first unique chars **********")
	str3 := "loveleetcode"

	fmt.Println(firstUniqChar(str3))

	fmt.Println("******** max in array **********")
	nums2 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums2, 3))

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

func reverse(x int) int {
	temp := x
	min, max := math.MinInt32, math.MaxInt32
	var newnum int
	if temp < 0 {
		temp = -temp
	}
	for temp > 0 {

		newnum = 10*newnum + temp%10
		temp = temp / 10
	}
	if newnum < min || newnum > max {
		return 0
	}
	if x < 0 {
		return -newnum
	}
	return newnum
}

func plusOne(digits []int) []int {

	maxlength := len(digits)
	digits[maxlength-1] = digits[maxlength-1] + 1
	for i := maxlength - 1; i > 0; i-- {
		if digits[i] == 10 {
			digits[i] = 0
			digits[i-1] = digits[i-1] + 1
		}
	}
	if digits[0] == 10 {
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var arr [26]int
	var arr1 [26]int

	for i := 0; i < len(s); i++ {
		arr[s[i]-97] = arr[s[i]-97] + 1
	}

	for i := 0; i < len(t); i++ {
		arr1[t[i]-97] = arr1[t[i]-97] + 1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != arr1[i] {
			return false
		}
	}

	return true
}

func firstUniqChar(s string) int {
	var arr [26]int

	if len(s) == 1 {
		return 0
	}

	for i := 0; i < len(s); i++ {
		arr[s[i]-97] = arr[s[i]-97] + 1
	}

	for i := 0; i < len(s); i++ {
		if arr[s[i]-97] == 1 {
			return i
		}
	}
	return -1
}

func maxSlidingWindow(nums []int, k int) []int {
	var res []int

	start := 0
	end := k

	res = append(res, findmax(nums, start, end))

	for j := k; j < len(nums); j++ {
		start = start + 1
		end = j + 1
		res = append(res, findmax(nums, start, end))
	}

	return res
}
func findmax(temp []int, start int, end int) int {
	max := math.MinInt
	for i := start; i < end; i++ {
		if temp[i] > max {
			max = temp[i]
		}
	}
	return max
}
