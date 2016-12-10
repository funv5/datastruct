package linkedList

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := LinkedList{}
	l.Append("99")

	if l.Length() != 1 || l.Head().Element != "99" {
		l.Print()
		t.Error("error at Append")
	}

	l.Insert(0, "88")

	if l.Head().Element != "88" || l.Head().Next.Element != "99" {
		l.Print()
		t.Error("error at Insert")
	}

	if l.Remove("88") != 0 {
		l.Print()
		t.Error("error at Remove 88")
	}
	if l.Head().Element != "99" {
		l.Print()
		t.Error("error at Remove")
	}

	l.Remove("99")
	if l.Length() != 0 {
		l.Print()
		t.Error("error at Remove All")
	}

	l.Append("99")
	l.Append("88")

	l.RemoveAt(0)

	if l.Head().Element != "88" || l.Length() != 1 {
		l.Print()
		t.Error("error at RemoveAt")
	}

	if l.IndexOf("88") == -1 || l.IndexOf("99") != -1 {
		l.Print()
		t.Error("error at IndexOf")
	}

	l.Remove("88")

	if l.IsEmpty() != true || l.Size() != 0 {
		l.Print()
		t.Error("error at IsEmpty && Size")
	}
}
