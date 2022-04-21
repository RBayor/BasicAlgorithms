package main

import "fmt"

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

type SinglyLinkedList struct {
	head   *LinkedListNode
	length int
}

func (l *SinglyLinkedList) pushFirst(value int) {
	previous := l.head
	l.head = &LinkedListNode{value: value}
	l.head.next = previous
	l.length++
}

func (l *SinglyLinkedList) pushLast(value int) {
	node := &LinkedListNode{value: value}
	var current *LinkedListNode

	if l.length == 0 {
		l.head = node
		return
	}

	current = l.head

	for current.next != nil {
		current = current.next
	}

	current.next = node
	l.length++

}

func (l *SinglyLinkedList) pushAt(value, index int) {

	if l.length == 0 {
		l.pushFirst(value)
		return
	}

	if l.length < index {
		return
	}

	node := &LinkedListNode{value: value}
	current := l.head
	var previous *LinkedListNode

	for i := 0; i < index; i++ {
		previous = current
		current = current.next
	}

	node.next = current
	previous.next = node
	l.length++
}

func (l SinglyLinkedList) getFirst() string {
	return fmt.Sprintf("%d", l.head.value)
}

func (l SinglyLinkedList) getAt(index int) string {
	if index < 0 || index >= l.length {
		return ""
	}

	if index == 0 {
		return l.getFirst()

	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return fmt.Sprintf("%d", current.value)
}

func (l SinglyLinkedList) getLast() int {
	current := l.head

	for current.next != nil {
		current = current.next
	}

	return current.value
}

func (l *SinglyLinkedList) pop() {
	l.head = l.head.next
	l.length--
}

func (l *SinglyLinkedList) popAt(index int) {
	if index < 0 || index >= l.length {
		return
	}

	if index == 0 {
		l.pop()
		return
	}

	count := 0
	current := l.head
	var previous *LinkedListNode

	for count < index {
		previous = current
		current = current.next
		count++
	}

	previous.next = current.next
	l.length--

}

func (l *SinglyLinkedList) popValue(value int) {

	if l.length == 0 {
		return
	}

	if l.head.value == value {
		l.pop()
		return
	}

	previous := l.head

	for previous.next.value != value {
		if previous.next.next == nil {
			return
		}
		previous = previous.next
	}
	previous.next = previous.next.next
	l.length--
}

func (l *SinglyLinkedList) clear() {
	l.head = nil
	l.length = 0
}

func (l SinglyLinkedList) log() {

	current := l.head
	for l.length != 0 {
		fmt.Printf("%d -> ", current.value)
		current = current.next
		l.length--
	}
	fmt.Printf("null")
	fmt.Println()
}

func main() {
	list := SinglyLinkedList{}

	list.pushFirst(14)
	list.pushFirst(5)
	list.pushFirst(7)
	list.pushFirst(87)
	list.pushFirst(100)
	list.pushFirst(45)
	list.pushLast(90)
	list.pushAt(49, 3)

	// fmt.Println(list.getFirst())
	// fmt.Println(list.getAt(7))
	// fmt.Println(list.getLast())
	list.log()
	// list.popAt(20)
	// list.popValue(45)
	// list.clear()
	// list.pop()
	// list.log()

	// fmt.Println(list)

}
