package hash

import (
	"testing"
)

func TestPut(t *testing.T) {
	h := HashTable{}
	h.Put("abc", "99")
	v := h.Get("abc")
	if v != "99" {
		t.Errorf("key abc get %s expect %s \n", v, "99")
	}
}

func TestRemove(t *testing.T) {
	h := HashTable{}
	h.Put("abc", "99")
	h.Remove("abc")
	v := h.Get("abc")
	if v == "99" {
		t.Errorf("key abc get %s expect %s \n", v, "")
	}
}
