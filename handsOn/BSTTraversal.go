package main

import "fmt"

type bstnode struct {
	data      int
	leftNode  *bstnode
	rightNode *bstnode
}

type bst struct {
	rootNode *bstnode
}

func (b *bst) insert(value int) {
	b.insertRec(b.rootNode, value)
}

func (b *bst) insertRec(new1 *bstnode, num int) *bstnode {

	if b.rootNode == nil {
		b.rootNode = &bstnode{
			data: num,
		}
		return b.rootNode
	}

	if new1 == nil {
		return &bstnode{data: num}
	}

	if num <= new1.data {
		new1.leftNode = b.insertRec(new1.leftNode, num)
	}

	if num > new1.data {
		new1.rightNode = b.insertRec(new1.rightNode, num)
	}
	return new1
}

func (b *bst) inorder() {
	b.inorderRec(b.rootNode)
	fmt.Println()
}

func (b *bst) inorderRec(node *bstnode) {
	if node != nil {
		b.inorderRec(node.leftNode)
		fmt.Print(node.data, " ")
		b.inorderRec(node.rightNode)
	}
}

func main3() {
	treeImp := &bst{}
	treeImp.insert(5)
	treeImp.insert(3)
	treeImp.insert(7)
	treeImp.insert(6)
	treeImp.insert(8)
	treeImp.insert(2)
	treeImp.inorder()

}
