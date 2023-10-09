package collection

import "sort"

type StringSet map[string]struct{}

func NewStringSet(vs ...string) StringSet {
	return make(StringSet).AddAll(vs...)
}

func (s StringSet) AddAll(vs ...string) StringSet {
	for _, v := range vs {
		s[v] = struct{}{}
	}
	return s
}

func (s StringSet) Union(other StringSet) StringSet {
	for k := range other {
		s[k] = struct{}{}
	}
	return s
}

func (s StringSet) Add(v string) {
	s[v] = struct{}{}
}

func (s StringSet) Contains(v string) bool {
	_, exists := s[v]
	return exists
}

func (s StringSet) Sorted() []string {
	ss := make([]string, 0, len(s))
	for k := range s {
		ss = append(ss, k)
	}
	sort.Strings(ss)
	return ss
}
