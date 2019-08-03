package jstruct

type stack struct {
	items []*jsNode
}

func (s *stack) len() int {
	return len(s.items)
}

func (s *stack) push(n *jsNode) {
	s.items = append(s.items, n)
}

func (s *stack) pop() (*jsNode, bool) {
	if s.len() == 0 {
		return nil, false
	}

	res := s.items[s.len()-1]
	s.items = s.items[:s.len()-1]

	return res, true
}
