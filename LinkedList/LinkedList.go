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
	head   *Node[T]
	length int
}

func NewList[T any](values ...T) *LinkedList[T] {
	ll := &LinkedList[T]{
		head:   nil,
		length: 0,
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
	if n.length == 0 && n.head == nil {
		n.length++
		n.head = newNode
		return
	}
	currentNode := n.head
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
	curNode := n.head
	for curNode != nil {
		fmt.Println(curNode.value)
		curNode = curNode.nextNode
	}
}

func (n *LinkedList[T]) DeleteAtNode(index int) error {
	if n.length < index || n.length == 0 || n.head == nil || index == 0 {
		return fmt.Errorf("Incorrect position. Total nodes in element is %d ", n.length)
	}
	currentNode := n.head
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
	if n.length < index || n.length == 0 || n.head == nil || index == 0 {
		return fmt.Errorf("Incorrect position. Total nodes in element is %d ", n.length)
	}
	currentNode := n.head
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
	if index == 0 || index < n.length || n.length == 0 {
		return t, fmt.Errorf("Incorrect position. Total nodes in element is %d ", n.length)
	}

	currentNode := n.head
	for i := 1; i < index; i++ {
		currentNode = currentNode.nextNode
	}
	return currentNode.value, nil
}

func (n *LinkedList[T]) GetAllValues() ([]T, error) {
	if n.length == 0 {
		return nil, errors.New("Empty list")
	}
	list := make([]T, 0)
	curNode := n.head
	for curNode != nil {
		list = append(list, curNode.value)
		curNode = curNode.nextNode
	}
	return list, nil
}
