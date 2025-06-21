package ast

// Program representa el programa completo de MiniScript.
// Contiene una lista de sentencias (statements) que se ejecutarán en orden.
type Program struct {
	Statements []*Statement `{ @@ }` // Una o más sentencias
}

// Statement representa una sentencia individual del lenguaje.
// Puede ser una instrucción de impresión (print) o una asignación de variable.
type Statement struct {
	Print      *PrintStmt  `  @@` // Sentencia de impresión: print <expr>
	Assignment *Assignment `| @@` // Sentencia de asignación: <ident> = <expr>
}

// PrintStmt representa una sentencia de impresión.
// Contiene una única expresión que será evaluada y mostrada.
type PrintStmt struct {
	Expr *Expr `"print" @@` // La expresión a imprimir
}

// Assignment representa una asignación de valor a una variable.
// Asocia un nombre (identificador) a una expresión evaluada.
type Assignment struct {
	Name  string `@Ident "="` // Nombre de variable
	Value *Expr  `@@`         // Valor asignado a través de una expresión
}

// Expr representa una expresión compuesta por un término inicial (Left)
// seguido de cero o más pares operador-término (Rest), permitiendo expresiones como: 1 + 2 * 3.
type Expr struct {
	Left *Term     `@@`  // Término inicial
	Rest []*OpTerm `@@*` // Secuencia opcional de operaciones binarias
}

// OpTerm representa una operación binaria: operador seguido de un término.
// Se usa para construir expresiones con múltiples operaciones como: a + b - c.
type OpTerm struct {
	Op   string `@("+" | "-" | "*" | "/" | "==" | "!=" | ">" | "<" | ">=" | "<=")` // Operador binario
	Term *Term  `@@`                                                               // Operando derecho
}

// Term representa los elementos más básicos de una expresión.
// Puede ser un entero, un número flotante, un string, un identificador (variable), o una subexpresión entre paréntesis.
type Term struct {
	Int    *int     `  @Int`       // Literal entero, ej: 42
	Float  *float64 `| @Float`     // Literal flotante, ej: 3.14
	String *string  `| @String`    // Literal de cadena, ej: "hola"
	Ident  *string  `| @Ident`     // Identificador, ej: variable
	Expr   *Expr    `| "(" @@ ")"` // Subexpresión entre paréntesis, ej: (a + b)
}
