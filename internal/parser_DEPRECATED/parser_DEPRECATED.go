package parser

import (
	"fmt"

	"github.com/DAlfaroV/miniscript/internal/lexer"
	"github.com/DAlfaroV/miniscript/internal/parser/ast.go"
)

// Parser convierte tokens en un AST de MiniScript.
type Parser struct {
	tokens  []lexer.Token
	current int
}

// New crea un nuevo parser con la lista de tokens.
func New(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens}
}

// ParseProgram construye el nodo raíz con todas las sentencias.
func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	for !p.isAtEnd() {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
	}
	return prog
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.peek().Type {
	case lexer.TOKEN_PRINT:
		return p.parsePrint()
	case lexer.TOKEN_IF:
		return p.parseIf()
	case lexer.TOKEN_WHILE:
		return p.parseWhile()
	case lexer.TOKEN_FOR:
		return p.parseFor()
	case lexer.TOKEN_FUNCTION:
		return p.parseFunction()
	case lexer.TOKEN_RETURN:
		return p.parseReturn()
	case lexer.TOKEN_BREAK:
		p.advance()
		return &ast.BreakStmt{}
	case lexer.TOKEN_CONTINUE:
		p.advance()
		return &ast.ContinueStmt{}
	default:
		if p.match(lexer.TOKEN_IDENTIFIER) && p.match(lexer.TOKEN_ASSIGN) {
			name := p.previous(1).Lexeme
			value := p.parseExpression()
			return &ast.AssignmentStmt{Name: name, Value: value}
		}
		expr := p.parseExpression()
		return &ast.ExpressionStmt{Expr: expr}
	}
}

// parseExpression inicia el análisis de expresiones.
func (p *Parser) parseExpression() ast.Expression {
	return p.parseEquality()
}

// Precedencia: equality -> comparison -> term -> factor -> unary -> primary
func (p *Parser) parseEquality() ast.Expression {
	expr := p.parseComparison()
	for p.match(lexer.TOKEN_EQ, lexer.TOKEN_NEQ) {
		op := p.previous(0).Lexeme
		right := p.parseComparison()
		expr = &ast.BinaryExpr{Left: expr, Operator: op, Right: right}
	}
	return expr
}

func (p *Parser) parseComparison() ast.Expression {
	expr := p.parseTerm()
	for p.match(lexer.TOKEN_GT, lexer.TOKEN_GTE, lexer.TOKEN_LT, lexer.TOKEN_LTE) {
		op := p.previous(0).Lexeme
		right := p.parseTerm()
		expr = &ast.BinaryExpr{Left: expr, Operator: op, Right: right}
	}
	return expr
}

func (p *Parser) parseTerm() ast.Expression {
	expr := p.parseFactor()
	for p.match(lexer.TOKEN_PLUS, lexer.TOKEN_MINUS) {
		op := p.previous(0).Lexeme
		right := p.parseFactor()
		expr = &ast.BinaryExpr{Left: expr, Operator: op, Right: right}
	}
	return expr
}

func (p *Parser) parseFactor() ast.Expression {
	expr := p.parseUnary()
	for p.match(lexer.TOKEN_SLASH, lexer.TOKEN_ASTERISK, lexer.TOKEN_PERCENT, lexer.TOKEN_CARET) {
		op := p.previous(0).Lexeme
		right := p.parseUnary()
		expr = &ast.BinaryExpr{Left: expr, Operator: op, Right: right}
	}
	return expr
}

func (p *Parser) parseUnary() ast.Expression {
	if p.match(lexer.TOKEN_NOT, lexer.TOKEN_MINUS) {
		op := p.previous(0).Lexeme
		right := p.parseUnary()
		return &ast.UnaryExpr{Operator: op, Right: right}
	}
	return p.parsePrimary()
}

func (p *Parser) parsePrimary() ast.Expression {
	tok := p.peek()
	switch tok.Type {
	case lexer.TOKEN_FALSE:
		p.advance()
		return &ast.LiteralExpr{Value: false}
	case lexer.TOKEN_TRUE:
		p.advance()
		return &ast.LiteralExpr{Value: true}
	case lexer.TOKEN_NIL:
		p.advance()
		return &ast.LiteralExpr{Value: nil}
	case lexer.TOKEN_NUMBER, lexer.TOKEN_STRING:
		p.advance()
		return &ast.LiteralExpr{Value: tok.Literal}
	case lexer.TOKEN_IDENTIFIER:
		p.advance()
		return &ast.VariableExpr{Name: tok.Lexeme}
	case lexer.TOKEN_LPAREN:
		p.advance()
		expr := p.parseExpression()
		p.consume(lexer.TOKEN_RPAREN, "Se esperaba ')' después de la expresión")
		return &ast.GroupingExpr{Expression: expr}
	default:
		panic(fmt.Sprintf("Token inesperado en expresión: %v", tok))
	}
}

func (p *Parser) parseIf() ast.Statement {
	p.advance() // consumir 'if'
	cond := p.parseExpression()
	thenBlock := p.parseBlock()

	var elseifConds []ast.Expression
	var elseifBodies [][]ast.Statement
	for p.match(lexer.TOKEN_ELSEIF) {
		elifCond := p.parseExpression()
		elifBody := p.parseBlock()
		elseifConds = append(elseifConds, elifCond)
		elseifBodies = append(elseifBodies, elifBody)
	}

	var elseBlock []ast.Statement
	if p.match(lexer.TOKEN_ELSE) {
		elseBlock = p.parseBlock()
	}

	return &ast.IfStmt{
		Condition:   cond,
		ThenBlock:   thenBlock,
		ElseIfConds: elseifConds,
		ElseIfBods:  elseifBodies,
		ElseBlock:   elseBlock,
	}
}

func (p *Parser) parseBlock() []ast.Statement {
	var stmts []ast.Statement
	for !p.check(lexer.TOKEN_END) && !p.isAtEnd() {
		st := p.parseStatement()
		if st != nil {
			stmts = append(stmts, st)
		}
	}
	p.consume(lexer.TOKEN_END, "Se esperaba 'end' al cerrar bloque")
	return stmts
}

func (p *Parser) parsePrint() ast.Statement {
	p.advance()
	value := p.parseExpression()
	return &ast.PrintStmt{Value: value}
}

func (p *Parser) parseWhile() ast.Statement {
	p.advance()
	cond := p.parseExpression()
	body := p.parseBlock()
	return &ast.WhileStmt{Condition: cond, Body: body}
}

func (p *Parser) parseFor() ast.Statement {
	p.advance()
	name := p.consume(lexer.TOKEN_IDENTIFIER, "Se esperaba identificador en for").Lexeme
	p.consume(lexer.TOKEN_ASSIGN, "Se esperaba '=' en for")
	start := p.parseExpression()
	p.consume(lexer.TOKEN_RANGE, "Se esperaba 'range/to' en for")
	end := p.parseExpression()
	body := p.parseBlock()
	return &ast.ForStmt{VarName: name, StartExpr: start, EndExpr: end, Body: body}
}

func (p *Parser) parseFunction() ast.Statement {
	p.advance()
	name := p.consume(lexer.TOKEN_IDENTIFIER, "Se esperaba nombre de función").Lexeme
	p.consume(lexer.TOKEN_LPAREN, "Se esperaba '('")
	var params []string
	if !p.check(lexer.TOKEN_RPAREN) {
		params = append(params, p.consume(lexer.TOKEN_IDENTIFIER, "Se esperaba parámetro").Lexeme)
		for p.match(lexer.TOKEN_COMMA) {
			params = append(params, p.consume(lexer.TOKEN_IDENTIFIER, "Se esperaba parámetro").Lexeme)
		}
	}
	p.consume(lexer.TOKEN_RPAREN, "Se esperaba ')'")
	body := p.parseBlock()
	return &ast.FunctionStmt{Name: name, Parameters: params, Body: body}
}

func (p *Parser) parseReturn() ast.Statement {
	p.advance()
	val := p.parseExpression()
	return &ast.ReturnStmt{Value: val}
}

// Métodos auxiliares
func (p *Parser) match(types ...lexer.TokenType) bool {
	for _, t := range types {
		if p.peek().Type == t {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) consume(t lexer.TokenType, msg string) lexer.Token {
	if p.peek().Type == t {
		return p.advance()
	}
	panic(msg)
}

func (p *Parser) check(t lexer.TokenType) bool {
	return p.peek().Type == t
}

func (p *Parser) advance() lexer.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.tokens[p.current-1]
}

func (p *Parser) previous(offset int) lexer.Token {
	return p.tokens[p.current-offset]
}

func (p *Parser) peek() lexer.Token {
	return p.tokens[p.current]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == lexer.TOKEN_EOF
}
