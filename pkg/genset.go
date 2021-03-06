package genset

// Set is a Generic Set/Hash Set containing values of type T.
type Set[T comparable] struct {
	list map[T]struct{}
}

// Has returns true if the value of type T is in Set[T] s and false if it is not in s.
func (s *Set[T]) Has(value T) bool {
	_, ok := s.list[value]
	return ok
}

// Add idempotently adds a value of type T to the Set[T] s and returns true if it was added to s and false if it was already in s.
func (s *Set[T]) Add(value T) bool {
	r := s.Has(value)
	if !r {
		s.list[value] = struct{}{}
	}
	return !r
}

// Remove removes the value of type T from the Set[T] s and returns true if it was removed from s and false if it was not in s.
func (s *Set[T]) Remove(value T) bool {
	r := s.Has(value)
	if r {
		delete(s.list, value)
	}
	return r
}

// Len returns the size of Set[T] s.
func (s *Set[T]) Len() int {
	return len(s.list)
}

// Clear removes all items from Set[T] s.
func (s *Set[T]) Clear() {
	s.list = nil
	s.list = make(map[T]struct{})
}

// AddMulti adds multiple values of type T to Set[T] s.
func (s *Set[T]) AddMulti(list ...T) {
	for _, v := range list {
		s.Add(v)
	}
}

// FilterFunc is a predicate definition for the Filter method that accepts a value of type T and returns a bool.
type FilterFunc[T comparable] func(value T) bool

// Filter returns a subset of Set[T] s, that contains only the values that satisfies the given predicate P.
func (s *Set[T]) Filter(P FilterFunc[T]) *Set[T] {
	res := New[T]()
	for v := range s.list {
		if P(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

// Union returns a Set[T] that is a union of Set[T] s and Set[T] other.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	res := New[T]()
	for v := range s.list {
		res.Add(v)
	}

	for v := range other.list {
		res.Add(v)
	}
	return res
}

// Intersect returns a Set[T] that is an intersaction of Set[T] s and Set[T] other.
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	res := New[T]()
	for v := range s.list {
		if other.Has(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

// Difference returns the subset from Set[T] s, that doesn't exists in Set[T] other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	res := New[T]()
	for v := range s.list {
		if other.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
}

// Values returns a []T containing all values in Set[T] s.
func (s *Set[T]) Values() []T {
	res := make([]T, s.Len())
	i := 0
	for k := range s.list {
		res[i] = k
		i++
	}
	return res
}

// Range is a Generator that will return all values in Set[T] s and will abort when chan abort is closed.
func (s *Set[T]) Range(abort <-chan struct{}) chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for k := range s.list {
			select {
			case ch <- k:
			case <-abort:
				return
			}
		}
	}()
	return ch
}

// New creates and new Set[T] s and returns *s.
func New[T comparable]() *Set[T] {
	s := &Set[T]{}
	s.list = make(map[T]struct{})
	return s
}
