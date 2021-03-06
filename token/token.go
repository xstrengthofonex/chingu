package token

import "fmt"

type Kind int

const (
	EOF Kind = iota
	LPAREN
	RPAREN
	DOT
	PLUS
	MINUS
	STAR
	SLASH
	TICK
	IDENT
	FN 
	BAD
)

var kindNames = map[Kind]string{
	EOF: "EOF",
	LPAREN: "LPAREN",
	RPAREN: "RPAREN",
	DOT: "DOT",
	PLUS: "PLUS",
	MINUS: "MINUS",
	STAR: "STAR",
	SLASH: "SLASH",
	TICK: "TICK",
	IDENT: "IDENT",
	FN: "FN",
	BAD: "BAD",
}

func (k Kind) String() string {
	if s, ok := kindNames[k]; ok {
		return s
	}
	return "INVALID"
}

var Keywords = map[string]Kind {
	"fn": FN,
}

type Token struct {
	Kind    Kind
	Literal string
}

func New(kind Kind, literal string) Token {
	return Token{kind, literal}
}

func EndOfFile() Token {
	return Token{EOF, ""}
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Kind: %s, Literal: %q}", t.Kind, t.Literal)
}
