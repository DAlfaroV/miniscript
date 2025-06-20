package lexer

import "fmt"

type LexError struct {
	Message string
	Line    int
	Column  int
}

func (e *LexError) Error() string {
	return fmt.Sprintf("[LexError Line:%d Col:%d] %s", e.Line, e.Column, e.Message)
}
