package doubly

import (
	"reflect"
)

type Node[T any] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type DoublyLinkedList[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (d *DoublyLinkedList[T]) prepend(item T) {
	node := &Node[T]{
		value: item,
	}

	d.length++
	if d.length == 1 {
		d.head = node
		d.tail = node
		return
	}

	node.next = d.head
	d.head.prev = node
	d.head = node
}

func (d *DoublyLinkedList[T]) insertAt(item T, idx int) {
	if idx > d.length {
		return
	}

	if idx == d.length {
		d.append(item)
		return
	} else if idx == 0 {
		d.prepend(item)
		return
	}

	node := &Node[T]{
		value: item,
	}

	d.length++
	curr := d.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}

	node.next = curr
	node.prev = curr.prev
	curr.prev = node
	node.prev.next = curr
}

func (d *DoublyLinkedList[T]) append(item T) {
	node := &Node[T]{
		value: item,
	}

	d.length++
	if d.length == 1 {
		d.head = node
		d.tail = node
		return
	}

	node.prev = d.tail
	d.tail.next = node
	d.tail = node
}

func (d *DoublyLinkedList[T]) remove(item T) T {
	curr := d.head
	for i := 0; i < d.length; i++ {
		if reflect.DeepEqual(curr.value, item) {
			break
		}
		curr = curr.next
	}
	if curr == nil {
		return *new(T)
	}
	return d.removeNode(curr)
}

func (d *DoublyLinkedList[T]) removeNode(node *Node[T]) T {
	d.length--
	if d.length == 0 {
		return d.head.value
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == d.head {
		d.head = node.next
	}
	if node == d.tail {
		d.tail = node.prev
	}
	return node.value
}

func (d *DoublyLinkedList[T]) removeAt(idx int) T {
	curr := d.getAt(idx)
	if curr != nil {
		return *new(T)
	}
	return d.removeNode(curr)
}

func (d *DoublyLinkedList[T]) get(idx int) T {
	return d.getAt(idx).value
}

func (d *DoublyLinkedList[T]) getAt(idx int) *Node[T] {
	curr := d.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	if curr != nil {
		return curr
	}
	return new(Node[T])
}
