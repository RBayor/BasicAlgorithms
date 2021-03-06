package main

import (
	"fmt"
)

var count int

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}

	return t

}

func (n *BinaryNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (n *BinaryNode) search(num int) bool {
	count++

	if n == nil {
		return false
	}

	if num == n.data {
		return true
	}

	if num > n.data {
		return n.right.search(num)
	} else {
		return n.left.search(num)
	}

}

func main() {
	tree := &BinaryTree{}
	tree.insert(10).insert(20).insert(50).insert(3).insert(18)

	fmt.Println(tree.root.search(18))
	fmt.Println(count)

	// start := time.Now()
	// sign := false

	// for i := 0; i < 100000; i++ {
	// 	if sign {
	// 		tree.insert(-i)
	// 		sign = false
	// 	} else {
	// 		tree.insert(i)
	// 		sign = true
	// 	}

	// }
	// duration := time.Since(start)

	// fmt.Println("Time taken:", duration)

}
