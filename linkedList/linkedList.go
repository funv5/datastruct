package linkedList

import "fmt"

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
func (l *LinkedList) Insert(position int, element interface{}) {
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
	current := l.Head
	if l.Head.Element == element {
		l.Head = l.Head.Next
		l.Length--
	}
	for {
		if current.Next == nil {
			break
		}
		if current.Next.Element == element {
			current.Next = current.Next.Next
			l.Length--
		}
		current = current.Next
	}
}

// RemoveAt remove at position
func (l *LinkedList) RemoveAt(position int) {
	if position >= 0 && position < l.Length {
		current := l.Head
		index := 0
		for {
			if position == index {
				current = current.Next
				l.Length--
				if index == 0 {
					l.Head = current
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
	current := l.Head
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
	return l.Head == nil
}

// Size ...
func (l *LinkedList) Size() int {
	return l.Length
}

// Print print LinkedList
func (l *LinkedList) Print() {
	fmt.Printf("%#v", l)
}
