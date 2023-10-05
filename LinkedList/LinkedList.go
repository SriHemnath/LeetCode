package LinkedList

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	value    T
	nextNode *Node[T]
}

type LinkedList[T any] struct {
	Head   *Node[T]
	Length int
}

func NewList[T any](values ...T) *LinkedList[T] {
	ll := &LinkedList[T]{
		Head:   nil,
		Length: 0,
	}
	for _, v := range values {
		ll.AddValue(v)
	}
	return ll
}

func (n *LinkedList[T]) AddValue(value T) {
	newNode := &Node[T]{
		value:    value,
		nextNode: nil,
	}
	if n.Length == 0 && n.Head == nil {
		n.Length++
		n.Head = newNode
		return
	}
	currentNode := n.Head
	for currentNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}
	currentNode.nextNode = newNode
	n.Length++
}

func (n *LinkedList[T]) PrintAll() {
	if n.Length == 0 {
		fmt.Println("Empty List")
	}
	curNode := n.Head
	for curNode != nil {
		fmt.Println(curNode.value)
		curNode = curNode.nextNode
	}
}

func (n *LinkedList[T]) DeleteAtNode(index int) error {
	if n.Length < index || n.Length == 0 || n.Head == nil || index == 0 {
		return fmt.Errorf("Incorrect position. Total nodes in element is %d ", n.Length)
	}
	currentNode := n.Head
	for i := 1; i < index-1; i++ {
		currentNode = currentNode.nextNode
	}
	if currentNode == nil {
		return nil
	}

	nextNode := currentNode.nextNode
	currentNode.nextNode = nextNode.nextNode
	return nil
}

func (n *LinkedList[T]) AddAtNode(index int, value T) error {
	if n.Length < index || n.Length == 0 || n.Head == nil || index == 0 {
		return fmt.Errorf("Incorrect position. Total nodes in element is %d ", n.Length)
	}
	currentNode := n.Head
	for i := 1; i < index; i++ {
		currentNode = currentNode.nextNode
	}
	newNode := &Node[T]{
		value:    value,
		nextNode: nil,
	}
	newNode.nextNode = currentNode.nextNode
	currentNode.nextNode = newNode
	return nil
}

func (n *LinkedList[T]) GetValueAt(index int) (T, error) {
	var t T
	if index == 0 || index < n.Length {
		return t, fmt.Errorf("Incorrect position. Total nodes in element is %d ", n.Length)
	}

	currentNode := n.Head
	for i := 1; i < index; i++ {
		currentNode = currentNode.nextNode
	}
	return currentNode.value, nil
}

func (n *LinkedList[T]) GetAllValues() ([]T, error) {
	if n.Length == 0 {
		return nil, errors.New("Empty list")
	}
	list := make([]T, 0)
	curNode := n.Head
	for curNode != nil {
		list = append(list, curNode.value)
		curNode = curNode.nextNode
	}
	return list, nil
}
