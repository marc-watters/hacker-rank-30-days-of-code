package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Solution struct{}

func (s Solution) insert(head *Node, data int) *Node {
	p := &Node{data: data}
	if head == nil {
		head = p
	} else if head.next == nil {
		head.next = p
	} else {
		start := head
		for start.next != nil {
			start = start.next
		}
		start.next = p
	}
	return head
}

func (s Solution) display(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (s Solution) removeDuplicates(head *Node) *Node {
	current := head
	for current != nil && current.next != nil {
		if current.data == current.next.data {
			current.next = current.next.next
		} else {
			current = current.next
		}
	}
	return head
}

func main() {
	s := new(Solution)

	var inputs int
	if _, err := fmt.Scan(&inputs); err != nil {
		panic(err)
	}

	var head *Node = nil
	for range inputs {
		var data int
		if _, err := fmt.Scan(&data); err != nil {
			panic(err)
		}
		head = s.insert(head, data)
	}

	head = s.removeDuplicates(head)
	s.display(head)
}
