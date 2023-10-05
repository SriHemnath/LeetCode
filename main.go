package main

import (
	"fmt"
	"github.com/SriHemnath/LeetCode/LinkedList"
	"log"
)

func main() {
	ll, err := LinkedList.NewList[int]()
	if err != nil {
		log.Fatal(err)
	}
	ll.PrintAll()
	ll.AddValue(1)
	ll.AddValue(1)
	ll.AddValue(2)
	ll.PrintAll()

	ll1, err := LinkedList.NewList[string]("Karnan", "Arjunan", "Bheeman", "Duriyodanan")
	if err != nil {
		log.Fatal(err)
	}
	v, err := ll1.GetValueAt(2)
	if err == nil {
		fmt.Printf("Value at index 2 %v\n", v)
	} else {
		fmt.Println(err)
	}
	ll1.InsertAtHead("Draupathai")
	ll1.InsertAtTail("Nakula")
	ll1.InsertAtNode(5, "Yudhistira")
	ll1.PrintAll()
}
