package main

import (
	"bufio"
	"fmt"
	"go-expr/prototype"
	"os"
	"strconv"
)

var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var buferr *bufio.Writer = bufio.NewWriter(os.Stderr)

func main() {
	defer func() {
		if r := recover(); r != nil {
			buferr.WriteString(os.Args[0])
			buferr.WriteString(": ")
			s, ok := r.(string)
			if !ok {
				buferr.WriteString("unknown error")
			} else {
				buferr.WriteString(s)
			}
			buferr.WriteString("\n")
		}
		bufout.Flush()
		buferr.Flush()
	}()

	//bufout.WriteString(fmt.Sprintf("%d\n", buildTree(os.Args[1:]).eval()))
	var opstack prototype.Stack = prototype.NewOperatorStack()
	var rpnstrs []string
	for i, str := range os.Args[1:] {
		if i%2 == 1 {
			op := prototype.NewOperator(str)
			if opstack.Size() == 0 {
				opstack.Push(op)
			} else {
				for opstack.Size() != 0 && opstack.Top().(prototype.Operator).Precedence() > op.Precedence() {
					rpnstrs = append(rpnstrs, opstack.Top().(prototype.Operator).Symbol())
					opstack.Pop()
				}
				opstack.Push(op)
			}
		} else {
			rpnstrs = append(rpnstrs, str)
		}
	}
	for opstack.Size() != 0 {
		rpnstrs = append(rpnstrs, opstack.Top().(prototype.Operator).Symbol())
		opstack.Pop()
	}
	for _, str := range rpnstrs {
		bufout.WriteString(str)
	}
}

func evaluation(args []string) {
	var lhs int32 = toInt32(args[1])
	var rhs int32 = toInt32(args[3])
	switch args[2] {
	case "+":
		bufout.WriteString(toString(lhs + rhs))
		bufout.WriteRune('\n')
	case "-":
		bufout.WriteString(toString(lhs - rhs))
		bufout.WriteRune('\n')
	case "*":
		bufout.WriteString(toString(lhs * rhs))
		bufout.WriteRune('\n')
	case "/":
		if rhs == 0 {
			panic("division by zero")
		} else {
			bufout.WriteString(toString(lhs / rhs))
			bufout.WriteRune('\n')
		}
	default:
		panic(fmt.Sprintf("unknown operator '%s'", args[2]))
	}
}

func toInt32(str string) int32 {
	converted, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		panic("non-integer argument")
	}
	return int32(converted)
}

func toString(value int32) string {
	return fmt.Sprintf("%d", value)
}

type expression interface {
	eval() int32
}
type number struct {
	value int32
}
type operator struct {
	value string
	lhs   expression
	rhs   expression
}

func (n number) eval() int32 {
	return n.value
}

func (op operator) eval() int32 {
	lhs := op.lhs.eval()
	rhs := op.rhs.eval()
	switch op.value {
	case "+":
		return lhs + rhs
	case "-":
		return lhs - rhs
	case "*":
		return lhs * rhs
	case "/":
		if rhs == 0 {
			panic("division by zero")
		}
		return lhs / rhs
	default:
		panic("syntax error: unexpected argument")
	}
}

func buildTree(strs []string) expression {
	var root operator
	switch len(strs) {
	case 0:
		panic("missing operand")
	case 1:
		return number{toInt32(strs[0])}
	default:
		root = operator{strs[1], nil, nil}
		root.lhs = number{toInt32(strs[0])}
		for i, str := range strs[2:] {
			if i%2 == 0 {
				root.rhs = number{toInt32(str)}
			} else {
				root = operator{str, root, nil}
			}
		}
	}
	return expression(root)
}
