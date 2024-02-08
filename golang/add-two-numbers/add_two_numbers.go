package main

import "fmt"

type Node struct {
	next *Node
	data int
}

func (node *Node) add(data int) {
	if node != nil {
		n := &Node{
			next: nil,
			data: data,
		}

		for node.next != nil {
			node = node.next
		}

		node.next = n
	}
	node.data = data
}

func (node *Node) print() {
	if node == nil {
		fmt.Println("list is emtpy")
	}

	head_node := node

	for node.next != nil {
		fmt.Printf("%d\n", node.data)
		node = node.next
	}

	node = head_node
}

func add_two_number(list1 *Node, list2 *Node) *Node {
	result := &Node{}

	return result
}

func main() {

}
