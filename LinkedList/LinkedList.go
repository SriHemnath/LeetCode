package LinkedList

import "fmt"

type Node[T any] struct {
	value    T
	nextNode *Node[T]
}

type LinkedList[T any] struct {
	Head   *Node[T]
	length int
}

func NewList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		Head:   nil,
		length: 0,
	}
}

func (n *LinkedList[T]) AddValue(value T) {
	newNode := &Node[T]{
		value:    value,
		nextNode: nil,
	}
	if n.length == 0 && n.Head == nil {
		n.length++
		n.Head = newNode
		return
	}
	currentNode := n.Head
	for currentNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}
	currentNode.nextNode = newNode
	n.length++
}

func (n *LinkedList[T]) PrintAll() {
	if n.length == 0 {
		fmt.Println("Empty List")
	}
	curNode := n.Head
	for curNode != nil {
		fmt.Println(curNode.value)
		curNode = curNode.nextNode
	}
}
