package list

import "errors"

// DoublyLinkedList ...
type DoublyLinkedList struct {
	length int
	head   *DNode
	tail   *DNode
}

// DNode ...
type DNode struct {
	Element interface{}
	Next    *DNode
	Prev    *DNode
}

// Head ...
func (d *DoublyLinkedList) Head() *DNode {
	return d.head
}

// Tail ...
func (d *DoublyLinkedList) Tail() *DNode {
	return d.tail
}

// Length ...
func (d *DoublyLinkedList) Length() int {
	return d.length
}

// Insert ...
func (d *DoublyLinkedList) Insert(position int, element interface{}) error {
	if position < 0 || position > d.length {
		return errors.New("position out of range")
	}
	node := DNode{}
	node.Element = element

	if position == 0 {
		node.Next = d.head
		if d.head != nil {
			d.head.Prev = &node
		}
		d.head = &node
		return nil
	}
	if position == d.length {
		node.Prev = d.tail
		d.tail.Next = &node
		d.tail = &node
		return nil
	}
	current := d.head.Next
	index := 1
	for {
		if current == nil {
			break
		}
		if position == index {
			node.Prev = current.Prev
			node.Next = current
			current.Prev.Next = &node
			current.Prev = &node
			return nil
		}
		current = current.Next
	}
	d.length++
	return nil
}
