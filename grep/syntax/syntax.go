package syntax

import (
	"github.com/vela-public/lua/grep/syntax/ast"
	"github.com/vela-public/lua/grep/syntax/lexer"
)

func Parse(s string) (*ast.Node, error) {
	return ast.Parse(lexer.NewLexer(s))
}

func Special(b byte) bool {
	return lexer.Special(b)
}
