package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{}

	q.Enqueue("77")
	if q.IsEmpty() {
		t.Error("Enqueue element but queue is Empty!")
	}

	q.Enqueue("88")
	ele := q.Dequeue()
	if ele != "77" {
		t.Error("Dequeue element but wrongß")
	}

	if q.Front() != "88" {
		t.Error("Front is not 88")
	}

	if q.Size() != 1 {
		q.Print()
		t.Error("Queue size is not right")
	}

	q.Clear()

	if q.Size() != 0 {
		q.Print()
		t.Error("size is wrong")
	}
}
