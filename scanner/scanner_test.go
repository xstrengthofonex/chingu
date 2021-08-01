package scanner

import (
	"testing"

	"github.com/xstrengthofonex/chingu/token"
)

var tokens = []token.Token{
	token.EndOfFile(),
}

func TestScan(t *testing.T) {
	src := ""
	for _, tok := range tokens {
		src += tok.Literal 
	}

	scanner := New(src)
	for _, want := range tokens {
		got := scanner.Scan()
		if got != want {
			t.Errorf("got token %s, want %s", got, want)
		}
	}
}
