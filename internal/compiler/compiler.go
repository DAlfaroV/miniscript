package compiler

import (
	"fmt"
	"strings"

	"github.com/DAlfaroV/miniscript/internal/ast"
)

// Compiler genera código en C a partir del AST de MiniScript.
type Compiler struct {
	sb        strings.Builder
	variables map[string]bool // para saber qué variables ya están declaradas
}

// NewCompiler crea un nuevo compilador.
func NewCompiler() *Compiler {
	return &Compiler{
		variables: make(map[string]bool),
	}
}

// Compile toma el AST y devuelve código en C como string.
func (c *Compiler) Compile(program *ast.Program) (string, error) {
	c.sb.WriteString("#include \"runtime.h\"\n\n")
	c.sb.WriteString("int main() {\n")

	// recorrer las sentencias
	for _, stmt := range program.Statements {
		if err := c.compileStatement(stmt); err != nil {
			return "", err
		}
	}

	c.sb.WriteString("return 0;\n")
	c.sb.WriteString("}\n")

	return c.sb.String(), nil
}

// compileStatement transforma una Statement a C
func (c *Compiler) compileStatement(stmt *ast.Statement) error {
	switch {
	case stmt.Print != nil:
		return c.compilePrint(stmt.Print)

	case stmt.Assignment != nil:
		return c.compileAssignment(stmt.Assignment)

	default:
		return fmt.Errorf("sentencia no soportada aún en el compilador")
	}
}

// compilePrint transforma un print a value_print
func (c *Compiler) compilePrint(p *ast.PrintStmt) error {
	exprStr, err := c.compileExpr(p.Expr)
	if err != nil {
		return err
	}
	c.sb.WriteString(fmt.Sprintf("  value_print(%s);\n", exprStr))
	return nil
}

// compileAssignment transforma una asignación a C
func (c *Compiler) compileAssignment(a *ast.Assignment) error {
	exprStr, err := c.compileExpr(a.Value)
	if err != nil {
		return err
	}

	if _, exists := c.variables[a.Name]; !exists {
		c.sb.WriteString(fmt.Sprintf("  Value %s = %s;\n", a.Name, exprStr))
		c.variables[a.Name] = true
	} else {
		c.sb.WriteString(fmt.Sprintf("  %s = %s;\n", a.Name, exprStr))
	}

	return nil
}

// compileExpr transforma una expresión simple a string en C
// (de momento solo soporta un solo término)
func (c *Compiler) compileExpr(e *ast.Expr) (string, error) {
	left, err := c.compileTerm(e.Left)
	if err != nil {
		return "", err
	}

	// En esta versión simplificada ignoramos e.Rest
	// Se podría mejorar para soportar operadores binarios
	return left, nil
}

// compileTerm transforma un término básico a C
func (c *Compiler) compileTerm(t *ast.Term) (string, error) {
	switch {
	case t.Int != nil:
		return fmt.Sprintf("(Value){ .type=VAL_INT, .i=%d }", *t.Int), nil

	case t.Float != nil:
		return fmt.Sprintf("(Value){ .type=VAL_FLOAT, .f=%f }", *t.Float), nil

	case t.String != nil:
		return fmt.Sprintf("(Value){ .type=VAL_STRING, .s=\"%s\" }", *t.String), nil

	case t.Ident != nil:
		return *t.Ident, nil

	case t.Expr != nil:
		return c.compileExpr(t.Expr)

	default:
		// bool y nil
		if t.Ident != nil {
			switch *t.Ident {
			case "true":
				return "(Value){ .type=VAL_BOOL, .b=1 }", nil
			case "false":
				return "(Value){ .type=VAL_BOOL, .b=0 }", nil
			case "nil":
				return "(Value){ .type=VAL_NIL }", nil
			}
		}
		return "", fmt.Errorf("tipo de término no soportado aún en el compilador")
	}
}
