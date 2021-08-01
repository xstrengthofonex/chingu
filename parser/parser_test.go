package parser

import (
	"fmt"
	"testing"

	"github.com/xstrengthofonex/chingu/ast"
	"github.com/xstrengthofonex/chingu/scanner"
	"github.com/xstrengthofonex/chingu/token"
)

type Parser struct {
	scanner *scanner.Scanner
	eof bool
	current token.Token
	previous token.Token
}

func New(src string) *Parser {
	p := &Parser{scanner: scanner.New(src)}
	p.next()
	return p
}

func (p *Parser) next() {
	if !p.eof {
		tok := p.scanner.Scan()
		if tok.Kind == token.EOF {
			p.eof = true
		}
		p.previous = p.current
		p.current = tok
	}
}

func (p *Parser) match(kind token.Kind) bool {
	if p.current.Kind == kind {
		p.next()
		return true
	}
	return false
}

func (p *Parser) expect(kind token.Kind) {
	if p.current.Kind == kind {
		p.next()
		return
	}
	panic(fmt.Errorf("expected %s, got %s", kind, p.current.Kind))
}

func (p *Parser) parseExpr() ast.Expr {
	if p.match(token.IDENT) {
		return &ast.Symbol{Value: p.previous.Literal}
	} else if p.match(token.LPAREN) {
		operator := p.parseExpr()
		operands := []ast.Expr{}
		for p.current.Kind != token.RPAREN && !p.eof {
			operands = append(operands, p.parseExpr())
		}
		p.expect(token.RPAREN)
		return &ast.Call{Operator: operator, Operands: operands}
	} else if p.match(token.FN) {
		params := []*ast.Symbol{}
		for p.match(token.IDENT) {
			params = append(params, &ast.Symbol{Value: p.previous.Literal})
		}
		p.expect(token.DOT)
		body := p.parseExpr()
		return &ast.Fn{Params: params, Body: body}
	}
	panic(fmt.Errorf("unexpected token %s", p.current))
}


func (p *Parser) Parse() (expr ast.Expr, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				panic(r)
			}
		} 
	}()

	expr = p.parseExpr()

	return expr, err
}

var parseTests = []struct{
	src string
	want ast.Expr
}{
	{"foo", &ast.Symbol{Value: "foo"}},
	{"bar", &ast.Symbol{Value: "bar"}},
	{"(f x y)", &ast.Call{
		Operator: &ast.Symbol{Value: "f"},
		Operands: []ast.Expr{
			&ast.Symbol{Value: "x"},
			&ast.Symbol{Value: "y"}}}},
	{"fn x . x", &ast.Fn{
		Params: []*ast.Symbol{{Value: "x"}},
		Body: &ast.Symbol{Value: "x"},
	}},
}

func TestParse(t *testing.T) {
	for _, tc := range parseTests {
		parser := New(tc.src)
		got, err := parser.Parse()
		
		if err != nil {
			t.Fatal(err)
		}

		if !tc.want.Equals(got) {
			t.Errorf("%q: got %v, want %v", tc.src, got, tc.want)
		}
	}
}
