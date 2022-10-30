package main

func main() {

}

func isValid(s string) bool {
	stack := NewStack()

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack.Push(string(v))
		}
		if (v == ')' && stack.Top() == "(") || (v == '}' && stack.Top() == "{") || (v == ']' && stack.Top() == "[") {
			stack.Pop()
		}

	}

	return stack.count == 0
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []string
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n string) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() string {
	if s.count == 0 {
		return ""
	}
	s.count--
	return s.nodes[s.count]
}

func (s *Stack) Top() string {
	if s.count == 0 {
		return ""
	}
	return s.nodes[s.count-1]
}
