package expression

import (
	"go-expr/stack"
	"strconv"
)

// InfixToPostfix transform infix-notated expression string to postfix-notated expression string using token that consists of pair of symbol string and precedence.
func InfixToPostfix(strs []string) []string {
	var rpn []string
	stack := stack.NewStack()

	for i, str := range strs {
		if i%2 == 0 {
			rpn = append(rpn, str)
		} else {
			t := OperatorToken{str, 0}
			for _, token := range operators {
				if token.symbol == t.symbol {
					t.precedence = token.precedence
					break
				}
			}
			if t.precedence == 0 {
				panic("systax error: unexprected argument ' '")
			}
			if stack.Size() == 0 {
				stack.Push(t)
			} else {
				for stack.Size() != 0 && stack.Top().(OperatorToken).precedence > t.precedence {
					rpn = append(rpn, stack.Top().(OperatorToken).symbol)
					stack.Pop()
				}
				stack.Push(t)
			}
		}
	}
	for stack.Size() != 0 {
		rpn = append(rpn, stack.Top().(OperatorToken).symbol)
		stack.Pop()
	}
	return rpn
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
func EvaluatePostfixStrings(strs []string) int32 {
	stack := stack.NewStack()

	for _, str := range strs {
		if isDigitString(str) {
			conv, err := strconv.ParseInt(str, 10, 32)
			if err != nil {
				panic("non-ingeter argument")
			}
			stack.Push(int32(conv))
		} else {
			if stack.Size() < 2 {
				panic("syntax error: missing argument")
			}
			lhs := stack.Top().(int32)
			stack.Pop()
			rhs := stack.Top().(int32)
			stack.Pop()
			switch str {
			case "+":
				stack.Push(lhs + rhs)
			case "-":
				stack.Push(lhs - rhs)
			case "*":
				stack.Push(lhs * rhs)
			case "/":
				if rhs == 0 {
					panic("division by zero")
				}
				stack.Push(lhs / rhs)
			default:
				panic("syntax error: unexpected argument")
			}
		}
	}
	return stack.Top().(int32)
}
