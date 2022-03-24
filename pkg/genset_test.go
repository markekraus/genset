package genset

import "testing"

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
