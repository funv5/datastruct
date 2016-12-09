package linkedList

import "fmt"

// LinkedList ...
type LinkedList struct {
	length int
	head   *Node
}

// Node ...
type Node struct {
	Element interface{}
	Next    *Node
}

// Append add element to LinkedList
func (l *LinkedList) Append(element interface{}) {
	n := Node{}
	n.Element = element

	if l.head == nil {
		l.head = &n
	} else {
		current := l.head
		for {
			if current.Next == nil {
				current.Next = &n
				break
			}
			current = current.Next
		}
	}
	l.length++
}

// Insert element at position
func (l *LinkedList) Insert(position int, element interface{}) {
	if position >= 0 && position <= l.length {
		n := Node{}
		n.Element = element
		current := l.head
		index := 0

		if position == index {
			n.Next = current
			l.head = &n
		} else {
			for {
				if current.Next == nil {
					break
				}
				current = current.Next
			}
		}
		l.length++
	}
}

//Remove remove element
func (l *LinkedList) Remove(element interface{}) {
	current := l.head
	if l.head.Element == element {
		l.head = l.head.Next
		l.length--
	}
	for {
		if current.Next == nil {
			break
		}
		if current.Next.Element == element {
			current.Next = current.Next.Next
			l.length--
		}
		current = current.Next
	}
}

// RemoveAt remove at position
func (l *LinkedList) RemoveAt(position int) {
	if position >= 0 && position < l.length {
		current := l.head
		index := 0
		for {
			if position == index {
				current = current.Next
				l.length--
				if index == 0 {
					l.head = current
				}
			}
			if current.Next == nil {
				break
			}
			index++
			current = current.Next
		}
	}
}

//IndexOf ...
func (l *LinkedList) IndexOf(element interface{}) int {
	current := l.head
	index := 0
	for {
		if current.Element == element {
			return index
		}
		if current.Next == nil {
			return -1
		}
		current = current.Next
	}
}

// IsEmpty ...
func (l *LinkedList) IsEmpty() bool {
	return l.length == 0
}

// Size ...
func (l *LinkedList) Size() int {
	return l.length
}

// Length ...
func (l *LinkedList) Length() int {
	return l.length
}

// Head ...
func (l *LinkedList) Head() Node {
	return *l.head
}

// Print print LinkedList
func (l *LinkedList) Print() {
	fmt.Printf("%#v", l)
}
