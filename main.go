package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var buferr *bufio.Writer = bufio.NewWriter(os.Stderr)

func main() {
	defer bufout.Flush()
	defer buferr.Flush()
	var args []string = os.Args
	var argc int = len(args)

	if argc == 1 {
		buferr.WriteString(args[0])
		buferr.WriteString(": ")
		buferr.WriteString("missing operand\n")
	} else {
		evaluation(args)
	}
}

func evaluation(args []string) {
	switch len(args) {
	case 2:
		bufout.WriteString(args[1])
	case 4:
		if isNumeric(args[1]) && isOperator(args[2]) && isNumeric(args[3]) {
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
					buferr.WriteString(args[0])
					buferr.WriteString(": ")
					buferr.WriteString("division by zero\n")
				} else {
					bufout.WriteString(toString(lhs / rhs))
					bufout.WriteRune('\n')
				}
			default:
				buferr.WriteString(args[0])
				buferr.WriteString(": ")
				buferr.WriteString("unknown error\n")
			}
		} else {
			buferr.WriteString(args[0])
			buferr.WriteString(": ")
			buferr.WriteString("unexpected format of argments")
		}
	default:
		buferr.WriteString(args[0])
		buferr.WriteString(": ")
		buferr.WriteString("unexpected number of argments")
	}
}

func isNumeric(str string) bool {
	switch len(str) {
	case 0:
		return false
	case 1:
		if str[0] < '0' || str[0] > '9' {
			return false
		}
	default:
		if str[0] != '-' && (str[0] < '0' || str[0] > '9') {
			return false
		}
		for _, r := range str[1:] {
			if r < '0' || r > '9' {
				return false
			}
		}
	}
	return true
}

func isOperator(str string) bool {
	if len(str) != 1 {
		return false
	}
	switch str[0] {
	case '+', '-', '*', '/':
		return true
	}
	return false
}

func toInt32(str string) int32 {
	var converted int32 = 0
	var isNegative = false
	if str[0] == '-' {
		isNegative = true
		str = str[1:]
	}
	for _, r := range str {
		var value int32 = r - '0'
		if value > 9 || value < 0 {
			if isNegative {
				return -converted
			}
			return converted
		}
		converted = converted*10 + value
	}
	if isNegative {
		return -converted
	}
	return converted
}

func toString(value int32) string {
	return fmt.Sprintf("%d", value)
}
