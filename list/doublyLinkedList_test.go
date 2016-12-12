package list

import (
	"testing"
)

func TestDubltyLinkedListInsert(t *testing.T) {
	d := DoublyLinkedList{}
	d.Insert(0, "99")
	d.Insert(1, "88")
	d.Insert(1, 100)

	current := d.head
	index := 0
	for {
		if current == nil {
			break
		}
		if index == 0 && current.Element != "99" {
			t.Errorf("index %d get %s expect string 99 ", index, current.Element)
		}
		if index == 1 && current.Element != 100 {
			t.Errorf("index %d get %s expect int 100 ", index, current.Element)
		}
		if index == 2 && current.Element != "88" {
			t.Errorf("index %d get %s expect string 88 ", index, current.Element)
		}
		current = current.Next
	}
}
