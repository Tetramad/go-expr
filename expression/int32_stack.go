package expression

type int32StackNode struct {
	next  *int32StackNode
	value int32
}

type int32Stack struct {
	head *int32StackNode
	size int
}

func newInt32Stack() *int32Stack {
	stack := new(int32Stack)
	stack.head = nil
	stack.size = 0
	return stack
}

func (stack *int32Stack) top() int32 {
	return stack.head.value
}

func (stack *int32Stack) pop() {
	stack.head = stack.head.next
	stack.size--
}

func (stack *int32Stack) push(value int32) {
	node := new(int32StackNode)
	node.next = stack.head
	node.value = value
	stack.head = node
	stack.size++
}
