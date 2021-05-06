package stackDemo

// 栈
type Stack struct {
	Items []string
}

func NewStack() (s *Stack) {

	return &Stack{
		Items: []string{},
	}
}

func (s *Stack) Push(x string) {
	s.Items = append(s.Items, x)
}

// Pop 取出后删除
func (s *Stack) Pop() (top string) {
	if !s.IsEmpty() {
		top := s.Items[len(s.Items)-1]

		items := s.Items
		s.Items = items[0 : len(items)-1]

		return top
	}
	return top
}

// Peek 仅查看，不删除
func (s *Stack) Peek() string {
	if !s.IsEmpty() {
		return s.Items[len(s.Items)-1]
	}
	return ""
}

func (s *Stack) IsEmpty() bool {
	if len(s.Items) > 0 {
		return false
	} else {
		return true
	}
}

func (s *Stack) Size() int {
	return len(s.Items)
}
