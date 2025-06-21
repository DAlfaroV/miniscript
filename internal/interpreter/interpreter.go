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

// execStatement ejecuta una única sentencia del programa.
// Actualmente soporta sentencias de impresión y asignación.
func (e *Env) execStatement(stmt *ast.Statement) error {
	switch {
	case stmt.Print != nil:
		// Si es una sentencia de impresión, evalúa la expresión y la muestra.
		value, err := e.evalExpr(stmt.Print.Expr)
		if err != nil {
			return err
		}
		fmt.Println(value)

	case stmt.Assignment != nil:
		// Si es una asignación, evalúa la expresión y guarda el resultado con el nombre de la variable.
		value, err := e.evalExpr(stmt.Assignment.Value)
		if err != nil {
			return err
		}
		e.Variables[stmt.Assignment.Name] = value

	default:
		// Si no es ninguna de las anteriores, es una sentencia no reconocida.
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
// un número entero, flotante, cadena, identificador (variable), o una subexpresión.
func (e *Env) evalTerm(term *ast.Term) (interface{}, error) {
	switch {
	case term.Int != nil:
		// Literal entero
		return *term.Int, nil

	case term.Float != nil:
		// Literal flotante
		return *term.Float, nil

	case term.String != nil:
		// Literal de cadena
		return *term.String, nil

	case term.Ident != nil:
		// Variable: busca su valor en el entorno
		val, ok := e.Variables[*term.Ident]
		if !ok {
			return nil, fmt.Errorf("variable no definida: %s", *term.Ident)
		}
		return val, nil

	case term.Expr != nil:
		// Subexpresión entre paréntesis
		return e.evalExpr(term.Expr)

	default:
		// Término inválido o nulo
		return nil, errors.New("término inválido")
	}
}

// applyOperator aplica una operación binaria básica entre dos operandos.
// Soporta suma, resta, multiplicación y división para enteros y floats.
// También permite concatenación con "+" para strings.
func applyOperator(left interface{}, op string, right interface{}) (interface{}, error) {
	switch l := left.(type) {

	case int:
		r, ok := right.(int)
		if !ok {
			return nil, errors.New("tipos incompatibles: se esperaba int")
		}
		switch op {
		case "+":
			return l + r, nil
		case "-":
			return l - r, nil
		case "*":
			return l * r, nil
		case "/":
			return l / r, nil
		default:
			return nil, fmt.Errorf("operador no soportado para enteros: %s", op)
		}

	case float64:
		r, ok := right.(float64)
		if !ok {
			return nil, errors.New("tipos incompatibles: se esperaba float")
		}
		switch op {
		case "+":
			return l + r, nil
		case "-":
			return l - r, nil
		case "*":
			return l * r, nil
		case "/":
			return l / r, nil
		default:
			return nil, fmt.Errorf("operador no soportado para floats: %s", op)
		}

	case string:
		r, ok := right.(string)
		if op == "+" && ok {
			// Concatenación de cadenas con "+"
			return l + r, nil
		}
		return nil, errors.New("solo se permite '+' entre strings")

	default:
		return nil, errors.New("tipo de dato no soportado en operación")
	}
}
