package tree

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	b := BinarySearchTree{}
	b.Insert(11)
	b.Insert(9)
	b.Insert(12)
	b.Insert(8)
	b.Insert(13)
	b.InOrderTraverse(console)
}

func console(key interface{}) {
	fmt.Println(key)
}
