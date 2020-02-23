package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var buferr *bufio.Writer = bufio.NewWriter(os.Stderr)

func main() {
	var args []string = os.Args

	defer func() {
		if r := recover(); r != nil {
			buferr.WriteString(args[0])
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

	if len(args) == 4 {
		evaluation(args)
	} else {
		panic("missing operand")
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
