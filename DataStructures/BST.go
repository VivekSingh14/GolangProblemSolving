package main

type Node1 struct {
	left  *Node
	data  int
	right *Node
}

type BsTree struct {
	root *Node1
}

func (b *BsTree) insert(data int) *BsTree {
	node := Node1{}
	if b.root == nil {
		node.data = data
		node.left = nil
		node.right = nil
		b.root = &node
	} else {
		b.root.insert(data)
	}
	return b
}

func (n *Node1) insert(data int) {
	node := Node1{}
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			node.data = data
			node.left = nil
			node.right = nil
		}
	}

}
