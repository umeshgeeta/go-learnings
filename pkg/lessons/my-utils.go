package lessons

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{data: []interface{}{}}
}

func (s *Stack) Push(val interface{}) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *Stack) Size() int {
	return len(s.data)
}
