package main

import (
	"fmt"
	"math"
)

type Node struct {
	left  *Node
	right *Node
	data  int
}

type Solution struct{}

func (s Solution) insert(root *Node, data int) *Node {
	node := &Node{data: data}
	if root == nil {
		return node
	}

	if data <= root.data {
		current := s.insert(root.left, data)
		root.left = current
	} else {
		current := s.insert(root.right, data)
		root.right = current
	}

	return root
}

func (s Solution) getHeight(root *Node) float64 {
	if root == nil {
		return -1
	} else {
		left := s.getHeight(root.left)
		right := s.getHeight(root.right)

		return math.Max(left, right) + 1
	}
}

func main() {
	var inputs int
	if _, err := fmt.Scan(&inputs); err != nil {
		panic(err)
	}

	S := new(Solution)
	root := (*Node)(nil)
	for range inputs {
		var data int
		if _, err := fmt.Scan(&data); err != nil {
			panic(err)
		}
		root = S.insert(root, data)
	}

	height := S.getHeight(root)
	fmt.Println(height)
}
