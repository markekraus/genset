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
	s := New[int]()
	s.Add(3)
	s.Add(2)
	s.Add(90)
	if !s.Has(3) {
		t.Errorf("s.Has(3) = %v, wanted %v", s.Has(3), true)
	}
	if !s.Has(2) {
		t.Errorf("s.Has(2) = %v, wanted %v", s.Has(2), true)
	}
	if !s.Has(90) {
		t.Errorf("s.Has(90) = %v, wanted %v", s.Has(90), true)
	}
	if s.Has(180) {
		t.Errorf("s.Has(180) = %v, wanted %v", s.Has(3), false)
	}
}

func TestAdd(t *testing.T) {
	s := New[int]()
	var r bool
	r = s.Add(3)
	if !r {
		t.Errorf("s.Add(3) = %v, wanted %v", r, true)
	}
	r = s.Add(2)
	if !r {
		t.Errorf("s.Add(2) = %v, wanted %v", r, true)
	}
	r = s.Add(90)
	if !r {
		t.Errorf("s.Add(90) = %v, wanted %v", r, true)
	}
	r = s.Add(90)
	if r {
		t.Errorf("s.Add(90) = %v, wanted %v", r, false)
	}
}

func TestRemove(t *testing.T) {
	s := New[int]()
	s.Add(3)
	s.Add(2)
	s.Add(90)
	var r bool
	r = s.Remove(3)
	if !r {
		t.Errorf("s.Remove(3) = %v, wanted %v", r, true)
	}
	r = s.Remove(2)
	if !r {
		t.Errorf("s.Remove(2) = %v, wanted %v", r, true)
	}
	r = s.Remove(180)
	if r {
		t.Errorf("s.Remove(180) = %v, wanted %v", r, false)
	}
	if s.Len() != 1 {
		t.Errorf("s.Len() = %d, want %d", s.Len(), 1)
	}
}

func TestLen(t *testing.T) {
	s := New[int]()
	s.Add(3)
	if s.Len() != 1 {
		t.Errorf("s.Len() = %d, want %d", s.Len(), 1)
	}
	s.Add(2)
	if s.Len() != 2 {
		t.Errorf("s.Len() = %d, want %d", s.Len(), 2)
	}
	s.Add(90)
	if s.Len() != 3 {
		t.Errorf("s.Len() = %d, want %d", s.Len(), 3)
	}
	s.Remove(90)
	if s.Len() != 2 {
		t.Errorf("s.Len() = %d, want %d", s.Len(), 2)
	}
	s.Remove(90)
	if s.Len() != 2 {
		t.Errorf("s.Len() = %d, want %d", s.Len(), 2)
	}
	s.Remove(180)
	if s.Len() != 2 {
		t.Errorf("s.Len() = %d, want %d", s.Len(), 2)
	}
}

func TestClear(t *testing.T) {
	s := New[int]()
	s.Add(3)
	s.Add(2)
	s.Add(90)
	if s.Len() != 3 {
		t.Errorf("s.Len() = %d, want %d", s.Len(), 3)
	}
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("s.Len() = %d, want %d", s.Len(), 0)
	}
}

func TestAddMulti(t *testing.T) {
	s := New[int]()
	s.AddMulti(3, 2, 90, 90)
	if s.Len() != 3 {
		t.Errorf("s.Len() = %d, want %d", s.Len(), 3)
	}
	if !s.Has(3) {
		t.Errorf("s.Has(3) = %v, wanted %v", s.Has(3), true)
	}
	if !s.Has(2) {
		t.Errorf("s.Has(2) = %v, wanted %v", s.Has(2), true)
	}
	if !s.Has(90) {
		t.Errorf("s.Has(90) = %v, wanted %v", s.Has(90), true)
	}
}

func TestFilter(t *testing.T) {
	s := New[int]()
	s.Add(3)
	s.Add(2)
	s.Add(90)
	r := s.Filter(func(value int) bool {
		return value < 50
	})
	if r.Len() != 2 {
		t.Errorf("r.Len() = %d, want %d", r.Len(), 2)
	}
	if r.Has(90) {
		t.Errorf("r.Has(90) = %v, wanted %v", r.Has(90), false)
	}
	if !r.Has(3) {
		t.Errorf("r.Has(3) = %v, wanted %v", r.Has(3), true)
	}
	if !r.Has(2) {
		t.Errorf("r.Has(2) = %v, wanted %v", r.Has(2), true)
	}
}

func TestUnion(t *testing.T) {
	s1 := New[int]()
	s1.Add(3)
	s1.Add(2)
	s1.Add(90)
	s2 := New[int]()
	s2.Add(5)
	s2.Add(6)
	s2.Add(90)
	r := s1.Union(s2)
	if r.Len() != 5 {
		t.Errorf("r.Len() = %d, want %d", r.Len(), 5)
	}
	if !r.Has(90) {
		t.Errorf("r.Has(90) = %v, wanted %v", r.Has(90), true)
	}
	if !r.Has(3) {
		t.Errorf("r.Has(3) = %v, wanted %v", r.Has(3), true)
	}
	if !r.Has(2) {
		t.Errorf("r.Has(2) = %v, wanted %v", r.Has(2), true)
	}
	if !r.Has(5) {
		t.Errorf("r.Has(5) = %v, wanted %v", r.Has(5), true)
	}
	if !r.Has(6) {
		t.Errorf("r.Has(6) = %v, wanted %v", r.Has(6), true)
	}
}

func TestIntersect(t *testing.T) {
	s1 := New[int]()
	s1.Add(3)
	s1.Add(2)
	s1.Add(90)
	s2 := New[int]()
	s2.Add(5)
	s2.Add(6)
	s2.Add(90)
	r := s1.Intersect(s2)
	if r.Len() != 1 {
		t.Errorf("r.Len() = %d, want %d", r.Len(), 1)
	}
	if !r.Has(90) {
		t.Errorf("r.Has(90) = %v, wanted %v", r.Has(90), true)
	}
	if r.Has(3) {
		t.Errorf("r.Has(3) = %v, wanted %v", r.Has(3), false)
	}
	if r.Has(2) {
		t.Errorf("r.Has(2) = %v, wanted %v", r.Has(2), false)
	}
	if r.Has(5) {
		t.Errorf("r.Has(5) = %v, wanted %v", r.Has(5), false)
	}
	if r.Has(6) {
		t.Errorf("r.Has(6) = %v, wanted %v", r.Has(6), false)
	}
}

func TestDifference(t *testing.T) {
	s1 := New[int]()
	s1.Add(3)
	s1.Add(2)
	s1.Add(90)
	s2 := New[int]()
	s2.Add(5)
	s2.Add(6)
	s2.Add(90)
	r := s1.Difference(s2)
	if r.Len() != 2 {
		t.Errorf("r.Len() = %d, want %d", r.Len(), 2)
	}
	if r.Has(90) {
		t.Errorf("r.Has(90) = %v, wanted %v", r.Has(90), false)
	}
	if !r.Has(3) {
		t.Errorf("r.Has(3) = %v, wanted %v", r.Has(3), true)
	}
	if !r.Has(2) {
		t.Errorf("r.Has(2) = %v, wanted %v", r.Has(2), true)
	}
	if r.Has(5) {
		t.Errorf("r.Has(5) = %v, wanted %v", r.Has(5), false)
	}
	if r.Has(6) {
		t.Errorf("r.Has(6) = %v, wanted %v", r.Has(6), false)
	}
}

func TestValues(t *testing.T) {
	s := New[int]()
	s.Add(3)
	s.Add(2)
	s.Add(90)
	s.Add(90)
	r := s.Values()
	if len(r) != 3 {
		t.Errorf("len(r) = %d, want %d", len(r), 3)
	}
	for i := 0; i < len(r); i++ {
		if r[i] != 3 && r[i] != 2 && r[i] != 90 {
			t.Errorf("r[i] = %d, want %d || %d || %d", r[i], 3, 2, 90)
		}
	}
}

func TestRange(t *testing.T) {
	s := New[int]()
	s.Add(3)
	s.Add(2)
	s.Add(90)
	s.Add(90)
	abort := make(chan struct{})
	i := 0
	for v := range s.Range(abort) {
		i++
		if v != 3 && v != 2 && v != 90 {
			t.Errorf("v = %d, want %d || %d || %d", v, 3, 2, 90)
			close(abort)
			break
		}
	}
	if i != s.Len() {
		t.Errorf("i = %d, want %d", i, s.Len())
	}
}

func TestStructs(t *testing.T) {
	s1 := New[*myType]()
	o1 := &myType{1, 2}
	o2 := &myType{3, 4}
	o3 := &myType{5, 6}
	o4 := &myType{6, 7}
	o5 := &myType{8, 9}
	s1.Add(o1)
	s1.Add(o2)
	s1.Add(o3)
	s1.Add(o3)
	if s1.Len() != 3 {
		t.Errorf("s1.Len() = %d, want %d", s1.Len(), 3)
	}
	if !s1.Has(o1) {
		t.Errorf("s1.Has(o1) = %v, wanted %v", s1.Has(o1), true)
	}
	if !s1.Has(o2) {
		t.Errorf("s1.Has(o2) = %v, wanted %v", s1.Has(o2), true)
	}
	if !s1.Has(o3) {
		t.Errorf("s1.Has(o3) = %v, wanted %v", s1.Has(o3), true)
	}
	if s1.Has(o4) {
		t.Errorf("s1.Has(o4) = %v, wanted %v", s1.Has(o4), false)
	}
	s2 := New[*myType]()
	s2.Add(o1)
	s2.Add(o3)
	s2.Add(o4)
	s2.Add(o5)
	u := s1.Union(s2)
	if u.Len() != 5 {
		t.Errorf("u.Len() = %d, want %d", u.Len(), 5)
	}
	if !u.Has(o1) {
		t.Errorf("u.Has(o1) = %v, wanted %v", u.Has(o1), true)
	}
	if !u.Has(o2) {
		t.Errorf("u.Has(o2) = %v, wanted %v", u.Has(o2), true)
	}
	if !u.Has(o3) {
		t.Errorf("u.Has(o3) = %v, wanted %v", u.Has(o3), false)
	}
	if !u.Has(o4) {
		t.Errorf("u.Has(o4) = %v, wanted %v", u.Has(o4), false)
	}
	if !u.Has(o5) {
		t.Errorf("u.Has(o4) = %v, wanted %v", u.Has(o5), false)
	}
}
