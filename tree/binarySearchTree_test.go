package tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	b := BinarySearchTree{}
	b.Insert(11)
	b.Insert(9)
	b.Insert(12)
	b.Insert(8)
	b.Insert(13)
	b.Insert(100)
	b.Insert(3)
	b.Insert(9)
	b.InOrderTraverse(console)
	b.PreOrderTraverse(console)
	b.PostOrderTraverse(console)
	if b.Min() != 3 {
		t.Errorf("min is 3 but get %d \n", b.Min())
	}

	if b.Max() != 100 {
		t.Errorf("Max is 100 but get %d \n", b.Max())
	}

	if b.Search(3) != true {
		t.Errorf("Search 3 but not found \n")
	}

	if b.Search(2) == true {
		t.Errorf("Search 2 but found \n")
	}
}

func console(key interface{}) {
	fmt.Println(key)
}
