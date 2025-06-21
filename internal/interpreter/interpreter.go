package interpreter

import (
	"errors"
	"fmt"

	"github.com/DAlfaroV/miniscript/internal/ast"
)

// Env representa el entorno de ejecución del programa MiniScript.
// Guarda las variables y sus valores durante la ejecución.
type Env struct {
	Variables map[string]interface{}
}

// Run ejecuta el programa completo recorriendo todas las sentencias del AST.
// Si ocurre un error en alguna sentencia, se detiene la ejecución.
func (e *Env) Run(program *ast.Program) error {
	for _, stmt := range program.Statements {
		// Ejecuta cada sentencia en orden
		if err := e.execStatement(stmt); err != nil {
			return err
		}
	}
	return nil
}

// execStatement ejecuta una sola sentencia del AST
func (e *Env) execStatement(stmt *ast.Statement) error {
	switch {
	case stmt.Print != nil:
		// print expr
		value, err := e.evalExpr(stmt.Print.Expr)
		if err != nil {
			return err
		}
		fmt.Println(value)

	case stmt.Assignment != nil:
		// x = expr
		value, err := e.evalExpr(stmt.Assignment.Value)
		if err != nil {
			return err
		}
		e.Variables[stmt.Assignment.Name] = value

	case stmt.If != nil:
		// if condición ... end
		cond, err := e.evalExpr(stmt.If.Condition)
		if err != nil {
			return err
		}
		// Verifica que sea bool
		condBool, ok := cond.(bool)
		if !ok {
			return errors.New("la condición del 'if' no es booleana")
		}
		if condBool {
			// Ejecuta el bloque del if
			for _, s := range stmt.If.Body {
				if err := e.execStatement(s); err != nil {
					return err
				}
			}
		}

	case stmt.While != nil:
		// while condición ... end
		for {
			cond, err := e.evalExpr(stmt.While.Condition)
			if err != nil {
				return err
			}
			condBool, ok := cond.(bool)
			if !ok {
				return errors.New("la condición del 'while' no es booleana")
			}
			if !condBool {
				break
			}
			for _, s := range stmt.While.Body {
				if err := e.execStatement(s); err != nil {
					return err
				}
			}
		}

	default:
		return errors.New("sentencia desconocida")
	}

	return nil
}

// evalExpr evalúa una expresión compuesta: un término inicial seguido de operadores y términos adicionales.
// Ejemplo: 1 + 2 * 3 se representa como un Expr con un Left y un slice de OpTerm.
func (e *Env) evalExpr(expr *ast.Expr) (interface{}, error) {
	// Evalúa el término inicial (izquierdo)
	left, err := e.evalTerm(expr.Left)
	if err != nil {
		return nil, err
	}

	result := left

	// Aplica cada operación binaria de izquierda a derecha
	for _, opTerm := range expr.Rest {
		right, err := e.evalTerm(opTerm.Term)
		if err != nil {
			return nil, err
		}

		// Aplica el operador (como +, -, *, /, etc.)
		result, err = applyOperator(result, opTerm.Op, right)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// evalTerm evalúa un término individual que puede ser:
// un número entero, flotante, cadena, identificador (variable), literal especial o una subexpresión.
func (e *Env) evalTerm(term *ast.Term) (interface{}, error) {
	switch {
	case term.Int != nil:
		// Literal entero (por ejemplo: 42)
		return *term.Int, nil

	case term.Float != nil:
		// Literal flotante (por ejemplo: 3.14)
		return *term.Float, nil

	case term.String != nil:
		// Literal de cadena (por ejemplo: "Hola")
		return *term.String, nil

	case term.Ident != nil:
		ident := *term.Ident

		// Verifica si el identificador es un literal booleano o nil
		switch ident {
		case "true":
			return true, nil
		case "false":
			return false, nil
		case "nil":
			return nil, nil
		}

		// Si no es un literal especial, se busca como variable en el entorno
		val, ok := e.Variables[ident]
		if !ok {
			return nil, fmt.Errorf("variable no definida: %s", ident)
		}
		return val, nil

	case term.Expr != nil:
		// Subexpresión entre paréntesis (por ejemplo: (3 + 4))
		return e.evalExpr(term.Expr)

	default:
		// El término no contiene nada reconocible
		return nil, errors.New("término inválido")
	}
}

// applyOperator aplica una operación binaria básica entre dos operandos.
// Soporta suma, resta, multiplicación y división para enteros y floats.
// También permite concatenación con "+" para strings.
// applyOperator aplica una operación binaria entre dos operandos.
func applyOperator(left interface{}, op string, right interface{}) (interface{}, error) {
	switch l := left.(type) {
	case int:
		switch r := right.(type) {
		case int:
			switch op {
			case "+":
				return l + r, nil
			case "-":
				return l - r, nil
			case "*":
				return l * r, nil
			case "/":
				return l / r, nil
			case "==":
				return l == r, nil
			case "!=":
				return l != r, nil
			case "<":
				return l < r, nil
			case "<=":
				return l <= r, nil
			case ">":
				return l > r, nil
			case ">=":
				return l >= r, nil
			default:
				return nil, fmt.Errorf("operador no soportado para enteros: %s", op)
			}
		default:
			return nil, errors.New("tipos incompatibles con enteros")
		}

	case float64:
		switch r := right.(type) {
		case float64:
			switch op {
			case "+":
				return l + r, nil
			case "-":
				return l - r, nil
			case "*":
				return l * r, nil
			case "/":
				return l / r, nil
			case "==":
				return l == r, nil
			case "!=":
				return l != r, nil
			case "<":
				return l < r, nil
			case "<=":
				return l <= r, nil
			case ">":
				return l > r, nil
			case ">=":
				return l >= r, nil
			default:
				return nil, fmt.Errorf("operador no soportado para floats: %s", op)
			}
		default:
			return nil, errors.New("tipos incompatibles con floats")
		}

	case string:
		r, ok := right.(string)
		if op == "+" && ok {
			return l + r, nil
		} else if (op == "==" || op == "!=") && ok {
			return (l == r), nil
		}
		return nil, fmt.Errorf("operador no soportado para strings: %s", op)

	case bool:
		r, ok := right.(bool)
		if !ok {
			return nil, errors.New("tipos incompatibles con booleanos")
		}
		switch op {
		case "==":
			return l == r, nil
		case "!=":
			return l != r, nil
		default:
			return nil, fmt.Errorf("operador no soportado para booleanos: %s", op)
		}

	default:
		return nil, errors.New("operador no soportado para ese tipo")
	}
}
