package hash

// HashTable ...
type HashTable struct {
	table map[int]interface{}
}

func loseloseHashCode(key string) int {
	hash := 0
	for _, char := range []rune(key) {
		hash += int(char)
	}
	return hash % 37
}

// Put ...
func (h *HashTable) Put(key string, value interface{}) {
	if h.table == nil {
		h.table = make(map[int]interface{})
	}
	h.table[loseloseHashCode(key)] = value
}

// Remove ...
func (h *HashTable) Remove(key string) {
	delete(h.table, loseloseHashCode(key))
}

// Get ...
func (h *HashTable) Get(key string) interface{} {
	return h.table[loseloseHashCode(key)]
}
