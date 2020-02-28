package expression

import (
	"go-expr/stack"
	"strconv"
)

// InfixToPostfix transform infix-notated expression string to postfix-notated expression string using token that consists of pair of symbol string and precedence.
func InfixToPostfix(strs []string) []interface{} {
	var rpn []interface{}
	stack := stack.NewStack()

	for _, str := range strs {
		switch {
		case isOperator(str):
			conv, err := operatorFromString(str)
			if err != nil {
				panic(err.Error())
			}
			for stack.Size() != 0 && stack.Top().(Operator).Precedence() > conv.Precedence() {
				rpn = append(rpn, stack.Top())
				stack.Pop()
			}
			stack.Push(conv)
		case isDigitString(str):
			conv, err := strconv.ParseInt(str, 10, 32)
			if err != nil {
				panic(err.Error())
			}
			rpn = append(rpn, int32(conv))
		default:
			panic("unexpected argument")
		}
	}
	for stack.Size() != 0 {
		rpn = append(rpn, stack.Top())
		stack.Pop()
	}
	return rpn
}

func isOperator(str string) bool {
	for _, op := range operators {
		if op.Symbol() == str {
			return true
		}
	}
	return false
}

func isDigitString(str string) bool {
	for _, chr := range str {
		if chr < '0' || chr > '9' {
			return false
		}
	}
	return true
}

// EvaluatePostfixStrings evaludate postfix-notated expression string to int32
func EvaluatePostfixStrings(rpn []interface{}) int32 {
	stack := stack.NewStack()

	for _, expr := range rpn {
		if conv, ok := expr.(Operator); ok {
			var operands []interface{}
			for i := 0; i < conv.OperandCount(); i++ {
				operands = append(operands, stack.Top())
				stack.Pop()
			}
			for i := 0; i < len(operands)/2; i++ {
				operands[i], operands[len(operands)-i-1] = operands[len(operands)-i-1], operands[i]
			}
			value, err := conv.Evaluate(operands...)
			if err != nil {
				panic(err.Error())
			}
			stack.Push(value)
		} else if conv, ok := expr.(int32); ok {
			stack.Push(conv)
		} else {
			panic("unkown error")
		}
	}
	return stack.Top().(int32)
}
