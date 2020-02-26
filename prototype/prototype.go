package prototype

// Operator ...
type Operator struct {
	symbol      string
	precedenece uint8
}

// Symbol ...
func (op Operator) Symbol() string {
	return op.symbol
}

// Precedence ...
func (op Operator) Precedence() uint8 {
	return op.precedenece
}

var operators [13]Operator = [13]Operator{
	Operator{"*", 5}, Operator{"/", 5}, Operator{"%", 5}, Operator{"&", 5},
	Operator{"+", 4}, Operator{"-", 4}, Operator{"|", 4},
	Operator{"=", 3}, Operator{"!=", 3}, Operator{"<", 3}, Operator{"<=", 3}, Operator{">", 3}, Operator{">=", 3},
}

// NewOperator ...
func NewOperator(symbol string) Operator {
	for _, op := range operators {
		if op.symbol == symbol {
			return op
		}
	}
	panic("syntax error: unexpected argument ' '")
}

type operatorStackNode struct {
	next  *operatorStackNode
	value Operator
}

// OperatorStack ...
type OperatorStack struct {
	top  *operatorStackNode
	size int
}

// Stack is exported
type Stack interface {
	Size() int
	Top() interface{}
	Push(value interface{})
	Pop()
}

// NewOperatorStack ...
func NewOperatorStack() *OperatorStack {
	stack := new(OperatorStack)
	stack.top = nil
	stack.size = 0
	return stack
}

// Size ...
func (stack *OperatorStack) Size() int {
	return stack.size
}

// Top ...
func (stack *OperatorStack) Top() interface{} {
	return stack.top.value
}

// Push ...
func (stack *OperatorStack) Push(value interface{}) {
	tmp := new(operatorStackNode)
	tmp.next = stack.top
	tmp.value = value.(Operator)
	stack.top = tmp
	stack.size++
}

// Pop ...
func (stack *OperatorStack) Pop() {
	tmp := stack.top.next
	stack.top = tmp
	stack.size--
}
