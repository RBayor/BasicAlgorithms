package main

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

type SinglyLinkedList struct {
	head   *LinkedListNode
	length int
}

func (l *SinglyLinkedList) push(value int) {
	n := LinkedListNode{}
	n.value = value

	if l.length == 0 {
		l.head = &n
		l.length++
		return
	}

	ptr := l.head

	for i := 0; i < l.length; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.length++
			return
		}
		ptr = ptr.next
	}
}

func main() {

}
