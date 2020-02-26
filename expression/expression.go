package expression

import "strconv"

// InfixToPostfix transform infix-notated expression string to postfix-notated expression string using token that consists of pair of symbol string and precedence.
func InfixToPostfix(strs []string) []string {
	var rpn []string
	stack := newTokenStack()

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
			if stack.size == 0 {
				stack.push(t)
			} else {
				for stack.size != 0 && stack.top().precedence > t.precedence {
					rpn = append(rpn, stack.top().symbol)
					stack.pop()
				}
				stack.push(t)
			}
		}
	}
	for stack.size != 0 {
		rpn = append(rpn, stack.top().symbol)
		stack.pop()
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
	stack := newInt32Stack()

	for _, str := range strs {
		if isDigitString(str) {
			conv, err := strconv.ParseInt(str, 10, 32)
			if err != nil {
				panic("non-ingeter argument")
			}
			stack.push(int32(conv))
		} else {
			if stack.size < 2 {
				panic("syntax error: missing argument")
			}
			lhs := stack.top()
			stack.pop()
			rhs := stack.top()
			stack.pop()
			switch str {
			case "+":
				stack.push(lhs + rhs)
			case "-":
				stack.push(lhs - rhs)
			case "*":
				stack.push(lhs * rhs)
			case "/":
				if rhs == 0 {
					panic("division by zero")
				}
				stack.push(lhs / rhs)
			default:
				panic("syntax error: unexpected argument")
			}
		}
	}
	return stack.top()
}
