package dictionary

import "testing"

func TestSet(t *testing.T) {
	d := Dictionary{}
	d.Set("99", 100)
	if !d.Has("99") {
		t.Errorf("set 99 is not right")
	}
	d.Set("99", 100)
	if d.Has("88") {
		t.Errorf("has not set 88 but get")
	}
	d.Set("88", 101)
	if !d.Has("88") {
		t.Errorf("set 88 but can not get")
	}
}

func TestRemove(t *testing.T) {
	d := Dictionary{}
	d.Set("99", 100)
	d.Set("88", 100)
	if d.Remove("99") == false {
		t.Errorf("remove exist key get false")
	}
	if d.Remove("100") == true {
		t.Errorf("remove not exist key get true")
	}
	if d.Has("99") {
		t.Errorf("remove 99 but get 99")
	}

	if !d.Has("88") {
		t.Errorf("not remove 88 but 88 gone")
	}
}

func TestGet(t *testing.T) {
	d := Dictionary{}
	d.Set("99", 100)
	d.Set("88", 100)

	if d.Get("99") != d.Get("88") {
		t.Errorf("key 99 's value %d \n", d.Get("99"))
		t.Errorf("key 88 's value %d \n", d.Get("88"))
	}
}

func TestClear(t *testing.T) {
	d := Dictionary{}
	d.Set("99", 100)
	d.Set("88", 100)
	d.Clear()
	if d.Has("99") {
		t.Errorf("clear all but key 99 's value %d \n", d.Get("99"))
	}
}

func TestSize(t *testing.T) {
	d := Dictionary{}
	if d.Size() != 0 {
		t.Errorf("size is 0 get %d \n", d.Size())
	}
	d.Set("99", 100)
	d.Set("88", 100)
	if d.Size() != 2 {
		t.Errorf("size is 2 get %d \n", d.Size())
	}
}

func TestKeys(t *testing.T) {
	d := Dictionary{}
	d.Set("99", 100)
	d.Set("88", 100)

	m := map[interface{}]struct{}{
		"99": struct{}{},
		"88": struct{}{},
	}
	keys := d.Keys()
	for i := 0; i < len(keys); i++ {
		_, ok := m[keys[i]]
		if !ok {
			t.Errorf("%s not exist\n", keys[i])
		}
	}
}

func TestValues(t *testing.T) {
	d := Dictionary{}
	d.Set("99", 100)
	d.Set("88", 101)

	m := map[interface{}]struct{}{
		100: struct{}{},
		101: struct{}{},
	}
	values := d.Values()
	for i := 0; i < len(values); i++ {
		_, ok := m[values[i]]
		if !ok {
			t.Errorf("%s not exist\n", values[i])
		}
	}
}
