package main

import "fmt"

//Implementation of the Doubly Linked List.
type LinkedListNode struct {
	value    int
	next     *LinkedListNode
	previous *LinkedListNode
}

type DoublyLinkedList struct {
	head   *LinkedListNode
	length int
}

func (l *DoublyLinkedList) pushFirst(value int) {
	previous := l.head
	l.head = &LinkedListNode{value: value}
	l.head.next = previous
	l.head.previous = nil
	l.length++
}

func (l DoublyLinkedList) getAt(index int) int {
	if index < 0 || index >= l.length {
		return -1
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value
}

func (l DoublyLinkedList) log() {

	current := l.head
	for l.length != 0 {
		fmt.Printf("%d <--> ", current.value)
		current = current.next
		l.length--
	}

	fmt.Printf("null")
	fmt.Println()
}

func main() {
	list := DoublyLinkedList{}

	list.pushFirst(5)
	list.pushFirst(12)
	list.pushFirst(23)
	list.pushFirst(43)
	list.pushFirst(87)

	// list.log()
	fmt.Println(list.getAt(10))
}
