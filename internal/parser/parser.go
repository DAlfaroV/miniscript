package parser

import (
	"github.com/DAlfaroV/miniscript/internal/ast"

	"github.com/alecthomas/participle/v2"
)

var Parser = participle.MustBuild[ast.Program]()
