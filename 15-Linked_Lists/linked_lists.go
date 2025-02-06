package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func display(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}

func insert(head *Node, data int) *Node {
	if head == nil {
		return &Node{data: data}
	}

	current := head
	for current.next != nil {
		current = current.next
	}

	current.next = &Node{data: data}
	return head
}

func main() {
	var inputs int
	_, err := fmt.Scan(&inputs)
	if err != nil {
		panic(err)
	}

	var head *Node
	for range inputs {
		var data int
		_, err := fmt.Scan(&data)
		if err != nil {
			panic(err)
		}
		head = insert(head, data)
	}

	display(head)
}
