package expression

// OperatorToken type represents operators
type OperatorToken struct {
	symbol     string
	precedence uint8
}

// OperatorError type is error type
type OperatorError struct {
	msg string
}

func (err OperatorError) Error() string {
	return err.msg
}

var operators [5]OperatorToken = [5]OperatorToken{
	OperatorToken{"*", 5}, OperatorToken{"/", 5}, OperatorToken{"%", 5},
	OperatorToken{"+", 4}, OperatorToken{"-", 4},
}

func operatorFromString(str string) (OperatorToken, error) {
	var found OperatorToken
	for _, op := range operators {
		if op.symbol == str {
			found.symbol = op.symbol
			found.precedence = op.precedence
			return found, nil
		}
	}
	return found, OperatorError{"Operater not found"}
}
