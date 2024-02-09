package parser

import (
	"github.com/yama-coder/monkey/ast"
	"github.com/yama-coder/monkey/lexer"
	"github.com/yama-coder/monkey/token"
)

type Parser struct {
	L *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{L: l}
	// read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.L.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
