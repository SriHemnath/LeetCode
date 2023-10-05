package LinkedList

import (
	"errors"
	"fmt"
)

var (
	errIndexOutOfBound = errors.New("index out of bound")
)

type Node[T any] struct {
	value    T
	nextNode *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	length int
}

func NewList[T any](values ...T) (*LinkedList[T], error) {
	ll := &LinkedList[T]{}
	err := ll.Add(values...)
	return ll, err
}

func (n *LinkedList[T]) IsEmpty() bool {
	return n.length == 0
}

func (n *LinkedList[T]) Add(values ...T) error {
	for idx := range values {
		if err := n.InsertAtNode(n.length, values[idx]); err != nil {
			return err
		}
	}
	return nil
}

func (n *LinkedList[T]) InsertAtNode(index int, value T) error {
	if index < 0 || index > n.length {
		return errIndexOutOfBound
	}

	node := &Node[T]{
		value: value,
	}

	if n.IsEmpty() {
		n.head = node
		n.length++
		return nil
	}

	if index == 0 {
		n.InsertAtHead(value)
		return nil
	}
	if index == n.length {
		n.InsertAtTail(value)
		return nil
	}

	n.AddAtNode(index, value)
	return nil
}

func (l *LinkedList[T]) InsertAtHead(value T) {
	newNode := &Node[T]{
		value:    value,
		nextNode: l.head,
	}
	l.head = newNode
	l.length++
}

func (l *LinkedList[T]) InsertAtTail(value T) {
	node := &Node[T]{
		value:    value,
		nextNode: nil,
	}
	currentNode := l.head
	for currentNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}
	currentNode.nextNode = node
	l.length++
}

func (n *LinkedList[T]) AddValue(value T) {
	if n.length == 0 && n.head == nil {
		n.InsertAtHead(value)
		return
	}
	n.InsertAtTail(value)
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
	n.length++
	return nil
}

func (n *LinkedList[T]) GetValueAt(index int) (T, error) {
	var t T
	if index == 0 || index > n.length || n.length == 0 {
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
		return nil, errors.New("empty list")
	}
	list := make([]T, 0)
	curNode := n.head
	for curNode != nil {
		list = append(list, curNode.value)
		curNode = curNode.nextNode
	}
	return list, nil
}
