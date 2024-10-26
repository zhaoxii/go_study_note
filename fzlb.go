package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func ReverseNode(head *Node) *Node {
	curNode := head
	var preNode *Node
	var afterNode *Node

	for curNode != nil {
		afterNode = curNode.next
		curNode.next = preNode
		preNode = curNode
		curNode = afterNode
	}

	return preNode
}

func main() {
	n1 := &Node{val: 1, next: nil}
	n2 := &Node{val: 2, next: n1}
	n3 := &Node{val: 3, next: n2}
	n4 := &Node{val: 4, next: n3}
	temp := n4
	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
	fmt.Println("...........")
	r := ReverseNode(n4)
	for r != nil {
		fmt.Println(r.val)
		r = r.next
	}
}
