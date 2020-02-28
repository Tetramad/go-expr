package expression

// Operator type is represents operator
type Operator interface {
	Symbol() string
	Precedence() uint8
	OperandCount() int
	Evaluate(...interface{}) (interface{}, error)
}

type sum struct{}

func (op sum) Symbol() string {
	return "+"
}

func (op sum) Precedence() uint8 {
	return 4
}

func (op sum) OperandCount() int {
	return 2
}

func (op sum) Evaluate(operands ...interface{}) (interface{}, error) {
	if len(operands) != op.OperandCount() {
		return nil, OperatorError{"the number of operands does not match"}
	}
	lhs, lhsok := operands[0].(int32)
	rhs, rhsok := operands[1].(int32)
	if !lhsok || !rhsok {
		return nil, OperatorError{"non-integer arugment"}
	}
	return lhs + rhs, nil
}

type subtraction struct{}

func (op subtraction) Symbol() string {
	return "-"
}

func (op subtraction) Precedence() uint8 {
	return 4
}

func (op subtraction) OperandCount() int {
	return 2
}

func (op subtraction) Evaluate(operands ...interface{}) (interface{}, error) {
	if len(operands) != op.OperandCount() {
		return nil, OperatorError{"the number of operands does not match"}
	}
	lhs, lhsok := operands[0].(int32)
	rhs, rhsok := operands[1].(int32)
	if !lhsok || !rhsok {
		return nil, OperatorError{"non-interger argument"}
	}
	return lhs - rhs, nil
}

type multiplication struct{}

func (op multiplication) Symbol() string {
	return "*"
}

func (op multiplication) Precedence() uint8 {
	return 5
}

func (op multiplication) OperandCount() int {
	return 2
}

func (op multiplication) Evaluate(operands ...interface{}) (interface{}, error) {
	if len(operands) != op.OperandCount() {
		return nil, OperatorError{"the number of operands does not match"}
	}
	lhs, lhsok := operands[0].(int32)
	rhs, rhsok := operands[1].(int32)
	if !lhsok || !rhsok {
		return nil, OperatorError{"non-integer argument"}
	}
	return lhs * rhs, nil
}

type division struct{}

func (op division) Symbol() string {
	return "/"
}

func (op division) Precedence() uint8 {
	return 5
}

func (op division) OperandCount() int {
	return 2
}

func (op division) Evaluate(operands ...interface{}) (interface{}, error) {
	if len(operands) != op.OperandCount() {
		return nil, OperatorError{"the number of operands does not match"}
	}
	lhs, lhsok := operands[0].(int32)
	rhs, rhsok := operands[1].(int32)
	if !lhsok || !rhsok {
		return nil, OperatorError{"non-integer argument"}
	}
	if rhs == 0 {
		return nil, OperatorError{"division by zero"}
	}
	return lhs / rhs, nil
}

type modulo struct{}

func (op modulo) Symbol() string {
	return "%"
}

func (op modulo) Precedence() uint8 {
	return 5
}

func (op modulo) OperandCount() int {
	return 2
}

func (op modulo) Evaluate(operands ...interface{}) (interface{}, error) {
	if len(operands) != op.OperandCount() {
		return nil, OperatorError{"the number of operands does not match"}
	}
	lhs, lhsok := operands[0].(int32)
	rhs, rhsok := operands[1].(int32)
	if !lhsok || !rhsok {
		return nil, OperatorError{"non-integer argument"}
	}
	if rhs == 0 {
		return nil, OperatorError{"division by zero"}
	}
	return lhs % rhs, nil
}

// OperatorError type is error type
type OperatorError struct {
	msg string
}

func (err OperatorError) Error() string {
	return err.msg
}

var operators [5]Operator = [5]Operator{
	sum{}, subtraction{},
	multiplication{}, division{}, modulo{},
}

func operatorFromString(str string) (Operator, error) {
	for _, op := range operators {
		if op.Symbol() == str {
			return op, nil
		}
	}
	return nil, OperatorError{"Operater not found"}
}
