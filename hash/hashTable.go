package hash

import (
	"datastruct/list"
)

// HashTable ...
type HashTable struct {
	table map[int]*list.LinkedList
}

func djb2HashCode(key string) int {
	hash := 5381
	for _, r := range []rune(key) {
		hash += hash*33 + int(r)
	}
	return hash % 1013
}

func loseloseHashCode(key string) int {
	hash := 0
	for _, r := range []rune(key) {
		hash += int(r)
	}
	return hash % 37
}

type valuePair struct {
	Key   string
	Value interface{}
}

// Put ...
func (h *HashTable) Put(key string, value interface{}) {
	if h.table == nil {
		h.table = make(map[int]*list.LinkedList)
	}
	position := djb2HashCode(key)
	vp := valuePair{}
	vp.Key = key
	vp.Value = value
	if _, ok := h.table[position]; !ok {
		h.table[position] = &list.LinkedList{}
	}
	h.table[position].Append(vp)
}

// Remove ...
func (h *HashTable) Remove(key string) {
	position := djb2HashCode(key)
	if _, ok := h.table[position]; ok {
		vp := valuePair{}
		vp.Key = key
		current := h.table[position].Head()
		for {
			if current == nil {
				break
			}
			v, r := current.Element.(valuePair)
			if r && v.Key == key {
				h.table[position].Remove(v)
			}
			current = current.Next
		}
	}
}

// Get ...
func (h *HashTable) Get(key string) interface{} {
	position := djb2HashCode(key)
	if _, ok := h.table[position]; ok {
		current := h.table[position].Head()
		for {
			if current == nil {
				break
			}
			v, r := current.Element.(valuePair)
			if r && v.Key == key {
				return v.Value
			}
			current = current.Next
		}
	}
	return nil
}
