package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	data  int
}

type Solution struct{}

func (s Solution) insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{data: data}
	}

	if data <= root.data {
		cur := s.insert(root.left, data)
		root.left = cur
	} else {
		cur := s.insert(root.right, data)
		root.right = cur
	}

	return root
}

func (s Solution) levelOrder(root *Node) {
	queue := []*Node{root}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", node.data)

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	fmt.Println()
}

func main() {
	var inputs int
	if _, err := fmt.Scanf("%d", &inputs); err != nil {
		panic(err)
	}

	myTree := new(Solution)
	var root *Node = nil

	for range inputs {
		var data int
		if _, err := fmt.Scanf("%d", &data); err != nil {
			panic(err)
		}
		root = myTree.insert(root, data)
	}

	myTree.levelOrder(root)
}
