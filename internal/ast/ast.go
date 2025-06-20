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
	Left  *Term  `@@`
	Op    string `[ @("+" | "-" | "*" | "/" | "==" | "!=" | ">" | "<" | ">=" | "<=") `
	Right *Term  `  @@ ]`
}

type Term struct {
	Number *float64 `  @Float`
	Ident  *string  `| @Ident`
}
