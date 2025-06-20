package ast

type Program struct {
	Statements []*Statement `{ @@ }`
}

type Statement struct {
	Print      *PrintStmt  `  @@`
	Assignment *Assignment `| @@`
}

type PrintStmt struct {
	Expr *Expr `"print" @@`
}

type Assignment struct {
	Name  string `@Ident "="`
	Value *Expr  `@@`
}

type Expr struct {
	Left *Term     `@@`
	Rest []*OpTerm `@@*`
}

type OpTerm struct {
	Op   string `@("+" | "-" | "*" | "/" | "==" | "!=" | ">" | "<" | ">=" | "<=")`
	Term *Term  `@@`
}

type Term struct {
	Int   *int     `  @Int`
	Float *float64 `| @Float`
	Ident *string  `| @Ident`
	Expr  *Expr    `| "(" @@ ")"`
}
