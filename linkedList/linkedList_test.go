package linkedList

import "testing"

func TestAppend(t *testing.T) {
	l := LinkedList{}

	l.Append("99")
	l.Append("88")
	current := l.Head()
	index := 0
	for {
		if current == nil {
			break
		}
		if index == 0 && current.Element != "99" {
			t.Errorf("index %d Element is %s \n", index, current.Element)
		}
		if index == 1 && current.Element != "88" {
			t.Errorf("index %d Element is %s \n", index, current.Element)
		}
		index++
		current = current.Next
	}
}

func TestInsert(t *testing.T) {
	l := LinkedList{}

	l.Insert(0, "99")
	l.Insert(1, 88)
	l.Insert(0, 100)

	current := l.Head()
	index := 0
	for {
		if current == nil {
			break
		}
		if index == 0 && current.Element != 100 {
			t.Errorf("index %d Element is %s \n", index, current.Element)
		}
		if index == 1 && current.Element != "99" {
			t.Errorf("index %d Element is %s \n", index, current.Element)
		}
		if index == 2 && current.Element != 88 {
			t.Errorf("index %d Element is %s \n", index, current.Element)
		}
		index++
		current = current.Next
	}
}

func TestRemove(t *testing.T) {
	l := LinkedList{}

	l.Insert(0, "99")
	l.Insert(1, 88)
	l.Insert(0, 100)

	if l.Remove(88) != 2 {
		t.Errorf("remove Element %d 's index  is not %d \n", 88, 2)
	}
	l.Remove(100)

	if l.Head().Element != "99" {
		t.Errorf("index %d Element is %s \n", 0, l.Head().Element)
	}
}

func TestRemoveAt(t *testing.T) {
	l := LinkedList{}

	l.Insert(0, "99")
	l.Insert(1, 88)
	l.Insert(0, 100)

	element0 := l.RemoveAt(0)
	if element0 != 100 {
		t.Errorf("remove  index %d is not  Element %d 's\n", 0, element0)
	}
	element1 := l.RemoveAt(1)
	if element1 != 88 {
		t.Errorf("remove  index %d is not  Element %d 's\n", 1, element1)
	}
	element2 := l.RemoveAt(2)
	if element2 != nil {
		t.Errorf("remove not exists index is not Element %d 's\n", element2)
	}

}

func TestIndexOf(t *testing.T) {
	l := LinkedList{}

	l.Insert(0, "99")
	l.Insert(1, 88)
	l.Insert(0, 100)

	index0 := l.IndexOf(100)
	if index0 != 0 {
		t.Errorf("index %d is not  Element %d \n", index0, 100)
	}
	index1 := l.IndexOf("99")
	if index1 != 0 {
		t.Errorf("index %d is not  Element %s \n", index1, "99")
	}
}
