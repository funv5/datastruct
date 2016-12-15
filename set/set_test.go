package set

import "testing"

func TestAdd(t *testing.T) {
	s := Set{}
	s.Add("99")
	s.Add("88")
	s.Add(100)

	if s.Size() != 3 {
		t.Errorf("Size is 3 return %d \n", s.Size())
	}
	if !s.Has("99") {
		t.Error("Add 99 but not found ")
	}
	if !s.Has("88") {
		t.Error("Add 99 but not found ")
	}
	if !s.Has(100) {
		t.Error("Add 99 but not found ")
	}
}

func TestRemove(t *testing.T) {
	s := Set{}
	s.Add("99")
	s.Add("88")
	if s.Remove("88") == false && s.Has("88") {
		t.Errorf("Has 99 return true expect false")
	}
}

func TestHas(t *testing.T) {
	s := Set{}
	s.Add("99")
	if s.Has("99") == false {
		t.Errorf("Has 99 return true expect false")
	}
}

func TestClear(t *testing.T) {
	s := Set{}
	s.Add("99")
	s.Add("88")
	s.Add("77")
	if s.Size() == 0 {
		t.Errorf("Size is 3 return %d \n", s.Size())
	}
	s.Clear()
	if s.Size() != 0 {
		t.Error("Clear but size is not 0 ")
	}
}

func TestSize(t *testing.T) {
	s := Set{}
	if s.Size() != 0 {
		t.Errorf("Size is 0 return %d \n", s.Size())
	}
	s.Add("99")
	s.Add("88")
	s.Add("77")

	if s.Size() != 3 {
		t.Errorf("Size is 3 return %d \n", s.Size())
	}
}

func TestValues(t *testing.T) {
	s := Set{}
	m := map[interface{}]bool{
		"99": false,
		"88": false,
		"77": false,
	}
	for k := range m {
		s.Add(k)
	}

	for i, v := 0, s.Values(); i < len(v); i++ {
		if _, ok := m[v[i]]; !ok {
			t.Errorf("Values get %s not in map m\n", v[i])
		} else {
			m[v[i]] = true
		}
	}
	for k, v := range m {
		if v != true {
			t.Errorf("map m's key %s not true \n", k)
		}
	}

}
