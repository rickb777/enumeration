package collection

import (
	"cmp"
	"slices"
)

type Set[T cmp.Ordered] map[T]struct{}

func NewSet[T cmp.Ordered](vs ...T) Set[T] {
	return make(Set[T]).AddAll(vs...)
}

func (s Set[T]) AddAll(vs ...T) Set[T] {
	for _, v := range vs {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	for k := range other {
		s[k] = struct{}{}
	}
	return s
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Contains(v T) bool {
	_, exists := s[v]
	return exists
}

// Sorted returns the set as a sorted slice.
func (s Set[T]) Sorted() []T {
	ss := make([]T, 0, len(s))
	for k := range s {
		ss = append(ss, k)
	}
	slices.Sort(ss)
	return ss
}
