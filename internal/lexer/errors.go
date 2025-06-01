package lexer

import "fmt"

// LexError representa un error léxico ocurrido durante el escaneo.
type LexError struct {
	Message string // Descripción del error
	Line    int    // Línea donde ocurrió el error
	Column  int    // Columna aproximada donde ocurrió el error
}

// Error implementa la interfaz error.
func (e *LexError) Error() string {
	return fmt.Sprintf("[LexError Line:%d Col:%d] %s", e.Line, e.Column, e.Message)
}
