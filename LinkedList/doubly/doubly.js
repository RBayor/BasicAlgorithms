class LinkedListNode {
  constructor(value, next, previous) {
    this.value = value;
    this.next = next;
    this.previous = previous;
  }
}

/**
 * Implementation of a doubly linked list
 */
class DoublyLinkedList {
  constructor() {
    this.head = null;
    this.length = 0;
  }

  pushFirst(value) {
    if (!this.head) return (this.head = new LinkedListNode(value));

    this.head = new LinkedListNode(value, this.head, null);
    this.head.next.previous = this.head;

    this.length++;

    return this;
  }

  getAt(index) {
    if (index < 0 || index > this.length) return null;

    let current = this.head;
    let count = 0;
    while (count < index) {
      current = current.next;
      count++;
    }
    return current.value;
  }

  log() {
    let output = "";
    let current = this.head;

    while (current) {
      output += `${current.value} <--> `;
      current = current.next;
    }
    console.log(`${output}null`);
  }
}

let list = new DoublyLinkedList();

list.pushFirst(5);
list.pushFirst(12);
list.pushFirst(23);
list.pushFirst(43);
list.pushFirst(87);

list.log();
// console.log(list.getAt(4));
// console.log(list.getAt(5));
// console.log(list.getAt(3));
// console.log(list.getAt(1));
