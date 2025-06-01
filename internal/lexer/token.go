package lexer

// TokenType representa los distintos tipos de tokens en MiniScript.
type TokenType int

const (
	// Tokens especiales
	TOKEN_ILLEGAL TokenType = iota
	TOKEN_EOF

	// Identificadores y literales
	TOKEN_IDENTIFIER // nombre de variable o función
	TOKEN_NUMBER     // literal numérico
	TOKEN_STRING     // literal de cadena
	TOKEN_TRUE       // literal booleano true
	TOKEN_FALSE      // literal booleano false
	TOKEN_NIL        // literal nil

	// Palabras clave
	TOKEN_IF
	TOKEN_ELSE
	TOKEN_ELSEIF
	TOKEN_END
	TOKEN_WHILE
	TOKEN_FOR
	TOKEN_FUNCTION
	TOKEN_RETURN
	TOKEN_BREAK
	TOKEN_CONTINUE
	TOKEN_PRINT
	TOKEN_RANGE

	// Operadores aritméticos y de comparación
	TOKEN_PLUS     // +
	TOKEN_MINUS    // -
	TOKEN_ASTERISK // *
	TOKEN_SLASH    // /
	TOKEN_PERCENT  // %
	TOKEN_CARET    // ^
	TOKEN_EQ       // ==
	TOKEN_NEQ      // !=
	TOKEN_GT       // >
	TOKEN_GTE      // >=
	TOKEN_LT       // <
	TOKEN_LTE      // <=
	TOKEN_ASSIGN   // =

	// Delimitadores y puntuación
	TOKEN_LPAREN    // (
	TOKEN_RPAREN    // )
	TOKEN_LBRACKET  // [
	TOKEN_RBRACKET  // ]
	TOKEN_LBRACE    // {
	TOKEN_RBRACE    // }
	TOKEN_COMMA     // ,
	TOKEN_COLON     // :
	TOKEN_DOT       // .
	TOKEN_SEMICOLON // ;

	// Operadores lógicos basados en palabras
	TOKEN_AND
	TOKEN_OR
	TOKEN_NOT
)

// Token es la estructura que representa un token individual.
type Token struct {
	Type    TokenType   // El tipo de token (uno de los valores de TokenType)
	Lexeme  string      // El texto exacto extraído de la fuente
	Literal interface{} // Valor “parseado” (float64 para números, string sin comillas, bool para true/false)
	Line    int         // Número de línea donde apareció el token
	Column  int         // Número de columna aproximada donde empieza el token
}

// keywords mapea cadenas literales a su tipo de token correspondiente.
var keywords = map[string]TokenType{
	"if":       TOKEN_IF,
	"else":     TOKEN_ELSE,
	"elseif":   TOKEN_ELSEIF,
	"end":      TOKEN_END,
	"while":    TOKEN_WHILE,
	"for":      TOKEN_FOR,
	"function": TOKEN_FUNCTION,
	"return":   TOKEN_RETURN,
	"break":    TOKEN_BREAK,
	"continue": TOKEN_CONTINUE,
	"print":    TOKEN_PRINT,
	"range":    TOKEN_RANGE,
	"true":     TOKEN_TRUE,
	"false":    TOKEN_FALSE,
	"nil":      TOKEN_NIL,
	"and":      TOKEN_AND,
	"or":       TOKEN_OR,
	"not":      TOKEN_NOT,
}
