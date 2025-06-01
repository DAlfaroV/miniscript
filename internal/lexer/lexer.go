// internal/lexer/lexer.go

package lexer

import (
	"strconv"
	"strings"
	"unicode"
)

// Lexer realiza el análisis léxico sobre el texto fuente de MiniScript.
type Lexer struct {
	source  string  // Texto completo a escanear
	tokens  []Token // Lista de tokens generados
	start   int     // Índice de inicio del lexema actual
	current int     // Índice del carácter actual
	line    int     // Línea actual en el texto (comienza en 1)
	column  int     // Columna actual en la línea (comienza en 1)
}

// NewLexer crea una nueva instancia de Lexer inicializada.
func NewLexer(source string) *Lexer {
	return &Lexer{
		source:  source,
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
		column:  1,
	}
}

// ScanTokens recorre todo el texto y devuelve la lista de tokens o un error léxico.
func (l *Lexer) ScanTokens() ([]Token, error) {
	for !l.isAtEnd() {
		l.start = l.current
		if err := l.scanToken(); err != nil {
			return nil, err
		}
	}
	// Al completar, agregar token EOF
	l.tokens = append(l.tokens, Token{
		Type:    TOKEN_EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    l.line,
		Column:  l.column,
	})
	return l.tokens, nil
}

// scanToken procesa un solo token a partir del carácter en l.current.
func (l *Lexer) scanToken() error {
	ch := l.advance()
	switch ch {
	case ' ', '\r', '\t':
		// Ignorar espacios y tabulaciones
		return nil
	case '\n':
		l.line++
		l.column = 1
		return nil
	case '/':
		if l.match('/') {
			l.skipComment()
			return nil
		}
		l.addToken(TOKEN_SLASH)
		return nil
	case '"':
		return l.string()
	case '+':
		l.addToken(TOKEN_PLUS)
		return nil
	case '-':
		l.addToken(TOKEN_MINUS)
		return nil
	case '*':
		l.addToken(TOKEN_ASTERISK)
		return nil
	case '%':
		l.addToken(TOKEN_PERCENT)
		return nil
	case '^':
		l.addToken(TOKEN_CARET)
		return nil
	case '(':
		l.addToken(TOKEN_LPAREN)
		return nil
	case ')':
		l.addToken(TOKEN_RPAREN)
		return nil
	case '[':
		l.addToken(TOKEN_LBRACKET)
		return nil
	case ']':
		l.addToken(TOKEN_RBRACKET)
		return nil
	case '{':
		l.addToken(TOKEN_LBRACE)
		return nil
	case '}':
		l.addToken(TOKEN_RBRACE)
		return nil
	case ',':
		l.addToken(TOKEN_COMMA)
		return nil
	case ':':
		l.addToken(TOKEN_COLON)
		return nil
	case '.':
		l.addToken(TOKEN_DOT)
		return nil
	case ';':
		l.addToken(TOKEN_SEMICOLON)
		return nil
	case '!':
		if l.match('=') {
			l.addToken(TOKEN_NEQ)
		} else {
			return &LexError{
				Message: "Unexpected character '!' (did you mean '!=')",
				Line:    l.line,
				Column:  l.column,
			}
		}
		return nil
	case '=':
		if l.match('=') {
			l.addToken(TOKEN_EQ)
		} else {
			l.addToken(TOKEN_ASSIGN)
		}
		return nil
	case '<':
		if l.match('=') {
			l.addToken(TOKEN_LTE)
		} else {
			l.addToken(TOKEN_LT)
		}
		return nil
	case '>':
		if l.match('=') {
			l.addToken(TOKEN_GTE)
		} else {
			l.addToken(TOKEN_GT)
		}
		return nil
	default:
		if isDigit(ch) {
			l.number()
			return nil
		} else if isAlpha(ch) {
			l.identifier()
			return nil
		} else {
			return &LexError{
				Message: "Unexpected character '" + string(ch) + "'",
				Line:    l.line,
				Column:  l.column,
			}
		}
	}
}

// advance consume el carácter actual y avanza los índices de current y column.
func (l *Lexer) advance() rune {
	ch := rune(l.source[l.current])
	l.current++
	l.column++
	return ch
}

// match verifica si el siguiente carácter coincide con 'expected'. Si coincide, lo consume.
func (l *Lexer) match(expected rune) bool {
	if l.isAtEnd() {
		return false
	}
	if rune(l.source[l.current]) != expected {
		return false
	}
	l.current++
	l.column++
	return true
}

// peek devuelve el carácter actual sin consumirlo.
func (l *Lexer) peek() rune {
	if l.isAtEnd() {
		return '\000'
	}
	return rune(l.source[l.current])
}

// peekNext devuelve el carácter después del actual, sin consumirlo.
func (l *Lexer) peekNext() rune {
	if l.current+1 >= len(l.source) {
		return '\000'
	}
	return rune(l.source[l.current+1])
}

// isAtEnd indica si se alcanzó el final de la fuente.
func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

// addToken crea un token sin valor literal y lo agrega a la lista.
func (l *Lexer) addToken(tType TokenType) {
	text := l.source[l.start:l.current]
	l.tokens = append(l.tokens, Token{
		Type:    tType,
		Lexeme:  text,
		Literal: nil,
		Line:    l.line,
		Column:  l.column,
	})
}

// addTokenLiteral crea un token con valor literal y lo agrega a la lista.
func (l *Lexer) addTokenLiteral(tType TokenType, literal interface{}) {
	text := l.source[l.start:l.current]
	l.tokens = append(l.tokens, Token{
		Type:    tType,
		Lexeme:  text,
		Literal: literal,
		Line:    l.line,
		Column:  l.column,
	})
}

// skipComment avanza hasta el final de la línea, ignorando el comentario.
func (l *Lexer) skipComment() {
	for l.peek() != '\n' && !l.isAtEnd() {
		l.advance()
	}
}

// string maneja literales de cadena entre comillas dobles.
func (l *Lexer) string() error {
	for l.peek() != '"' && !l.isAtEnd() {
		if l.peek() == '\n' {
			l.line++
			l.column = 1
		}
		l.advance()
	}

	if l.isAtEnd() {
		return &LexError{
			Message: "Unterminated string literal",
			Line:    l.line,
			Column:  l.column,
		}
	}

	// Consumir la comilla de cierre
	l.advance()

	// Extraer el contenido sin las comillas
	raw := l.source[l.start+1 : l.current-1]
	// Reemplazar "" por " dentro del contenido
	value := strings.ReplaceAll(raw, `""`, `"`)

	l.addTokenLiteral(TOKEN_STRING, value)
	return nil
}

// number maneja literales numéricas (enteros y flotantes).
func (l *Lexer) number() {
	for isDigit(l.peek()) {
		l.advance()
	}
	// Verificar punto decimal
	if l.peek() == '.' && isDigit(l.peekNext()) {
		l.advance()
		for isDigit(l.peek()) {
			l.advance()
		}
	}
	raw := l.source[l.start:l.current]
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		// Si no se pudo convertir, marcamos como token ilegal
		l.addTokenLiteral(TOKEN_ILLEGAL, raw)
		return
	}
	l.addTokenLiteral(TOKEN_NUMBER, val)
}

// identifier maneja reconocimientos de identificadores y palabras clave.
func (l *Lexer) identifier() {
	for isAlphaNumeric(l.peek()) {
		l.advance()
	}
	text := l.source[l.start:l.current]
	// Chequear si es palabra clave
	if tokType, ok := keywords[text]; ok {
		l.addToken(tokType)
	} else {
		l.addTokenLiteral(TOKEN_IDENTIFIER, text)
	}
}

// isDigit retorna true si ch es un dígito [0-9].
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

// isAlpha retorna true si ch es letra [A-Za-z] o guión bajo _.
func isAlpha(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

// isAlphaNumeric retorna true si ch es letra, dígito o _.
func isAlphaNumeric(ch rune) bool {
	return isAlpha(ch) || isDigit(ch)
}
