package expression

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
