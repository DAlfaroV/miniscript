package parser

type Node interface {
	NodeType() string
}

type Statement interface {
	Node
	isStatement()
}

type Expression interface {
	Node
	isExpression()
}

type Program struct {
	Statements []Statement
}

func (p *Program) NodeType() string { return "Program" }

type ExpressionStmt struct {
	Expr Expression
}

func (s *ExpressionStmt) NodeType() string { return "ExpressionStmt" }
func (s *ExpressionStmt) isStatement()     {}

type PrintStmt struct {
	Value Expression
}

func (s *PrintStmt) NodeType() string { return "PrintStmt" }
func (s *PrintStmt) isStatement()     {}

type AssignmentStmt struct {
	Name  string
	Value Expression
}

func (s *AssignmentStmt) NodeType() string { return "AssignmentStmt" }
func (s *AssignmentStmt) isStatement()     {}

type IfStmt struct {
	Condition   Expression
	ThenBlock   []Statement
	ElseIfConds []Expression
	ElseIfBods  [][]Statement
	ElseBlock   []Statement
}

func (s *IfStmt) NodeType() string { return "IfStmt" }
func (s *IfStmt) isStatement()     {}

type WhileStmt struct {
	Condition Expression
	Body      []Statement
}

func (s *WhileStmt) NodeType() string { return "WhileStmt" }
func (s *WhileStmt) isStatement()     {}

type ForStmt struct {
	VarName   string
	StartExpr Expression
	EndExpr   Expression
	Body      []Statement
}

func (s *ForStmt) NodeType() string { return "ForStmt" }
func (s *ForStmt) isStatement()     {}

type FunctionStmt struct {
	Name       string
	Parameters []string
	Body       []Statement
}

func (s *FunctionStmt) NodeType() string { return "FunctionStmt" }
func (s *FunctionStmt) isStatement()     {}

type ReturnStmt struct {
	Value Expression
}

func (s *ReturnStmt) NodeType() string { return "ReturnStmt" }
func (s *ReturnStmt) isStatement()     {}

type BreakStmt struct{}

func (s *BreakStmt) NodeType() string { return "BreakStmt" }
func (s *BreakStmt) isStatement()     {}

type ContinueStmt struct{}

func (s *ContinueStmt) NodeType() string { return "ContinueStmt" }
func (s *ContinueStmt) isStatement()     {}

type BinaryExpr struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (e *BinaryExpr) NodeType() string { return "BinaryExpr" }
func (e *BinaryExpr) isExpression()    {}

type UnaryExpr struct {
	Operator string
	Right    Expression
}

func (e *UnaryExpr) NodeType() string { return "UnaryExpr" }
func (e *UnaryExpr) isExpression()    {}

type LiteralExpr struct {
	Value interface{}
}

func (e *LiteralExpr) NodeType() string { return "LiteralExpr" }
func (e *LiteralExpr) isExpression()    {}

type VariableExpr struct {
	Name string
}

func (e *VariableExpr) NodeType() string { return "VariableExpr" }
func (e *VariableExpr) isExpression()    {}

type GroupingExpr struct {
	Expression Expression
}

func (e *GroupingExpr) NodeType() string { return "GroupingExpr" }
func (e *GroupingExpr) isExpression()    {}

type CallExpr struct {
	Callee    Expression
	Arguments []Expression
}

func (e *CallExpr) NodeType() string { return "CallExpr" }
func (e *CallExpr) isExpression()    {}
