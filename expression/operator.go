package expression

// OperatorToken type represents operators
type OperatorToken struct {
	symbol     string
	precedence uint8
}

var operators [5]OperatorToken = [5]OperatorToken{
	OperatorToken{"*", 5}, OperatorToken{"/", 5}, OperatorToken{"%", 5},
	OperatorToken{"+", 4}, OperatorToken{"-", 4},
}
