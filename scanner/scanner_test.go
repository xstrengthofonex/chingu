package scanner

import (
	"fmt"
	"testing"

	"github.com/xstrengthofonex/chingu/token"
)

var tokens = []token.Token{
	token.New(token.LPAREN, "("),
	token.New(token.RPAREN, ")"),
	token.New(token.DOT, "."),
	token.New(token.PLUS, "+"),
	token.New(token.MINUS, "-"),
	token.New(token.STAR, "*"),
	token.New(token.SLASH, "/"),
	token.New(token.TICK, "`"),
	token.New(token.IDENT, "a"),
	token.New(token.IDENT, "foo"),
	token.New(token.IDENT, "bar"),
	token.New(token.IDENT, "foo_bar2"),
	token.New(token.IDENT, "_foo_bar2"),
	token.New(token.FN, "fn"),
	token.EndOfFile(),
}

func TestScan(t *testing.T) {
	src := ""
	for _, tok := range tokens {
		src += fmt.Sprintf(" \t%s\n", tok.Literal) 
	}

	scanner := New(src)
	for _, want := range tokens {
		got := scanner.Scan()
		if got != want {
			t.Errorf("got token %s, want %s", got, want)
		}
	}
}
