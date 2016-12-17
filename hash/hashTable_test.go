package hash

import (
	"testing"
)

func TestPut(t *testing.T) {
	h := HashTable{}
	h.Put("Mindy", "99")
	h.Put("Paul", "100")
	v := h.Get("Mindy")
	if v != "99" {
		t.Errorf("key abc get %s expect %s \n", v, "99")
	}
	v = h.Get("Paul")
	if v != "100" {
		t.Errorf("key abc get %s expect %s \n", v, "99")
	}
}

func TestRemove(t *testing.T) {
	h := HashTable{}
	h.Put("Mindy", "99")
	h.Put("Paul", "100")
	h.Remove("Paul")
	v := h.Get("Mindy")
	if v != "99" {
		t.Errorf("key abc get %s expect %s \n", v, "")
	}
}
