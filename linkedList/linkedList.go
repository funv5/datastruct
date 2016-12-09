package linkedList

import (
	"fmt"
)

// LinkedList ...
type LinkedList struct {
	Length int
	Head   *node
}

// LinkedList node
type node struct {
	Element interface{}
	Next    *node
}

// Append add element to LinkedList
func (l *LinkedList) Append(element interface{}) {
	n := node{}
	n.Element = element

	if l.Head == nil {
		l.Head = &n
	} else {
		current := l.Head
		for {
			if current.Next == nil {
				current.Next = &n
				break
			}
			current = current.Next
		}
	}
	l.Length++
}

// Insert element at position
func (l *LinkedList) Insert(position int, element interface{}) error {
	if position >= 0 && position <= l.Length {
		n := node{}
		n.Element = element
		current := l.Head
		index := 0

		if position == index {
			n.Next = current
			l.Head = &n
		} else {
			for {
				if current.Next == nil {
					break
				}
				current = current.Next
			}
		}

		l.Length++
	}
}

//Remove remove element
func (l *LinkedList) Remove(element interface{}) {
	for {

	}
}

// RemoveAt remove at position
func (l *LinkedList) RemoveAt(position int) error {

}

// Print print LinkedList
func (l *LinkedList) Print() {
	fmt.Printf("%#v", l)
}
