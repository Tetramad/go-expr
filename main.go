package main

import (
	"bufio"
	"go-expr/expression"
	"os"
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

	rpn := expression.InfixToPostfix(os.Args[1:])
	for _, str := range rpn {
		bufout.WriteString(str)
	}
	bufout.WriteString("\n")
}
