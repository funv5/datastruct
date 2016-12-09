package linkedList

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := LinkedList{}
	l.Length++
	if l.Head == nil {
		l.Print()
		t.Error("未初始化的地址为空")
	}
}
