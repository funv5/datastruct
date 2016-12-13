package list

import "testing"

func TestDubltyLinkedListInsert(t *testing.T) {
	d := DoublyLinkedList{}
	d.Insert(0, "99")
	d.Insert(1, "88")
	d.Insert(1, 100)
	d.Insert(2, 101)

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
		if index == 2 && current.Element != 101 {
			t.Errorf("index %d get %s expect string 101 ", index, current.Element)
		}
		index++
		current = current.Next
	}
}

func TestDubltyLinkedListRemoveAt(t *testing.T) {
	d := DoublyLinkedList{}
	d.Insert(0, "99")
	d.Insert(1, "88")
	d.Insert(1, 100)
	d.Insert(2, 101)

	del, err := d.RemoveAt(1)
	if del != 100 {
		t.Errorf("remove index 1 expect 100 get %#v err: %#v \n", del, err)
	}

	del, err = d.RemoveAt(0)
	if del != "99" {
		t.Errorf("remove index 1 expect 99 get %d err: %#v \n", del, err)
	}
	current := d.head
	index := 0
	for {
		if current == nil {
			break
		}
		if index == 0 && current.Element != 101 {
			t.Errorf("index %d get %s expect int 101 ", index, current.Element)
		}
		if index == 1 && current.Element != "88" {
			t.Errorf("index %d get %s expect string 88 ", index, current.Element)
		}
		index++
		current = current.Next
	}
}
