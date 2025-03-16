package parser

import (
	"github.com/victortsokonov/TheMonkeyLanguage/interpreter/ast"
	"github.com/victortsokonov/TheMonkeyLanguage/interpreter/lexer"
	"github.com/victortsokonov/TheMonkeyLanguage/interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

// reading parsers
