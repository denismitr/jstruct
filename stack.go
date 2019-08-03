package jstruct

type stack struct {
	items []*jsNode
}

func (s *stack) len() int {
	return len(s.items)
}

func (s *stack) contains(n *jsNode) bool {
	if len(s.items) == 0 {
		return false
	}

	for i := range s.items {
		if s.items[i].uniqueName() == n.uniqueName() {
			return true
		}
	}

	return false
}

func (s *stack) push(n *jsNode) {
	if s.contains(n) {
		return
	}

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
