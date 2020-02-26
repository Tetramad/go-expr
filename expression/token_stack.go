package expression

type tokenStackNode struct {
	next  *tokenStackNode
	value OperatorToken
}

type tokenStack struct {
	head *tokenStackNode
	size int
}

func newTokenStack() *tokenStack {
	stack := new(tokenStack)
	stack.head = nil
	stack.size = 0
	return stack
}

func (stack *tokenStack) top() OperatorToken {
	return stack.head.value
}

func (stack *tokenStack) pop() {
	stack.head = stack.head.next
	stack.size--
}

func (stack *tokenStack) push(token OperatorToken) {
	node := new(tokenStackNode)
	node.next = stack.head
	node.value = token
	stack.head = node
	stack.size++
}
