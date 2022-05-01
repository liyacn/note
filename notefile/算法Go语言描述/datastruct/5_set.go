package datastruct

// 可以利用map的key不重复的特性来实现set

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable](items []T) *Set[T] {
	m := make(map[T]struct{})
	for _, item := range items {
		m[item] = struct{}{}
	}
	return &Set[T]{m: m}
}

func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.m, v)
}

func (s *Set[T]) Clear() {
	clear(s.m)
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Items() []T {
	items := make([]T, 0, len(s.m))
	for item := range s.m {
		items = append(items, item)
	}
	return items
}
