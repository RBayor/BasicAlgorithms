class LinkedListNode {
  constructor(value, next) {
    this.value = value;
    this.next = next;
  }
}

/**
 * Implementation of a singly linked list
 */
class SinglyLinkedList {
  constructor() {
    this.head = null;
    this.length = 0;
  }

  pushFirst(value) {
    this.head = new LinkedListNode(value, this.head);
    this.length++;

    return this;
  }

  pushLast(value) {
    let newNode = new LinkedListNode(value);
    let current;

    if (!this.head) {
      this.head = newNode;
    } else {
      current = this.head;

      while (current.next) {
        current = current.next;
      }

      current.next = newNode;
    }
    this.length++;
  }

  pushAt(value, index) {
    if (index < 0 || index > this.length) return;

    if (index === 0) return this.pushFirst(value);

    const newNode = new LinkedListNode(value);
    let prev, curr;

    curr = this.head;
    let count = 0;
    while (count < index) {
      prev = curr;
      count++;
      curr = curr.next;
    }
    newNode.next = curr;
    prev.next = newNode;

    this.length++;
  }
  getFirst() {
    return this.head.value;
  }

  getAt(index) {
    if (index < 0 || index > this.length) return null;
    let current = this.head;

    for (let i = 0; i < index; i++) current = current.next;
    return current.value;
  }

  getLast() {
    let curr = this.head;
    while (curr.next) {
      curr = curr.next;
    }
    return curr.value;
  }

  pop() {
    this.head = this.head.next;
    this.length--;
  }

  popAt(index) {
    if (index < 0 || index > this.length) return null;

    if (index === 0) return this.pop();

    let curr = this.head;
    let prev;
    let count = 0;

    while (count < index) {
      prev = curr;
      curr = curr.next;
      count++;
    }

    prev.next = curr.next;
    this.length--;
  }

  clear() {
    this.head = null;
    this.length = 0;
  }

  log() {
    let output = "";
    let current = this.head;

    while (current) {
      output = `${output}${current.value} -> `;
      current = current.next;
    }
    console.log(`${output}null`);
  }
}

let list = new SinglyLinkedList();

list.pushFirst(14);
list.pushFirst(5);
list.pushFirst(7);
list.pushLast(87);
list.pushAt(100, 3);
list.pushAt(45, 0);
list.log();
// console.log(list.getFirst());
// console.log(list.getLast());
console.log(list.getAt(8));
// list.popAt(1);
// list.log();
