package genset

import "testing"

type myType struct {
	a, b int
}

func TestNew(t *testing.T) {
	s := New[int]()
	if len(s.list) != 0 {
		t.Errorf("len(s.list) = %d, want %d", len(s.list), 0)
	}
}

func TestHas(t *testing.T) {
	//s := New[int]()
}

func TestAdd(t *testing.T) {
	//s := New[int]()
}

func TestRemove(t *testing.T) {
	//s := New[int]()
}

func TestLen(t *testing.T) {
	//s := New[int]()
}

func TestClear(t *testing.T) {
	//s := New[int]()
}

func TestAddMulti(t *testing.T) {
	//s := New[int]()
}

func TestFilter(t *testing.T) {
	//s := New[int]()
}

func TestUnion(t *testing.T) {
	//s := New[int]()
}

func TestInterect(t *testing.T) {
	//s := New[int]()
}

func TestDifference(t *testing.T) {
	//s := New[int]()
}
