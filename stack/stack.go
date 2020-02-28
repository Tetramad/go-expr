package stack

type node struct {
	next  *node
	value interface{}
}

// Stack is a generic stack use interface{}
type Stack struct {
	head *node
	size int
}

// NewStack make empty stack.
func NewStack() *Stack {
	stack := new(Stack)
	stack.head = nil
	stack.size = 0
	return stack
}

// Top returns value of top of stack.
func (stack *Stack) Top() interface{} {
	if stack.size == 0 {
		panic("stack hasn't top value.")
	}
	return stack.head.value
}

// Pop pop out top of stack. It returns anything.
func (stack *Stack) Pop() {
	if stack.size == 0 {
		panic("stack hasn't any value to pop")
	}
	stack.head = stack.head.next
	stack.size--
}

// Push put value to top of stack.
func (stack *Stack) Push(value interface{}) {
	nd := new(node)
	nd.next = stack.head
	nd.value = value
	stack.head = nd
	stack.size++
}

// Size returns size of stack.
func (stack *Stack) Size() int {
	return stack.size
}
