package parser

import "fmt"

type ParseError struct {
	Message string
	Line    int
	Column  int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("[ParseError Line:%d Col:%d] %s", e.Line, e.Column, e.Message)
}
