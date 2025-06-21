package parser

import (
	"github.com/DAlfaroV/miniscript/internal/ast"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	miniLexer = lexer.MustSimple([]lexer.SimpleRule{
		{Name: "Comment", Pattern: `//[^\n]*`},
		// {Name: "Keyword", Pattern: `\b(if|while|print|end|true|false|nil)\b`},
		{Name: "Ident", Pattern: `[a-zA-Z_]\w*`},
		{Name: "Float", Pattern: `\d+\.\d+`}, // Primero Float
		{Name: "Int", Pattern: `\d+`},        // Luego Int
		{Name: "String", Pattern: `"([^"\\]|\\.)*"`},
		{Name: "Operator", Pattern: `==|!=|<=|>=|[+\-*/=<>]`},
		{Name: "Punct", Pattern: `[{}\(\)]`},
		{Name: "Whitespace", Pattern: `[ \t\r\n]+`},
	})
)

var Parser = participle.MustBuild[ast.Program](
	participle.Lexer(miniLexer),
	participle.Elide("Whitespace", "Comment"),
	participle.Unquote("String"),
)
