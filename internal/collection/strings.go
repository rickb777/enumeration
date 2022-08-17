package collection

type StringSet map[string]struct{}

func NewStringSet() StringSet {
	return make(StringSet)
}

func (s StringSet) Add(v string) {
	s[v] = struct{}{}
}

func (s StringSet) Contains(v string) bool {
	_, exists := s[v]
	return exists
}
