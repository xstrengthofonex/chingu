package token

import "fmt"

type Kind int

const (
	EOF Kind = iota
)

var kindNames = map[Kind]string{
	EOF: "EOF",
}

func (k Kind) String() string {
	if s, ok := kindNames[k]; ok {
		return s
	}
	return "INVALID"
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
	return fmt.Sprintf("Tokens{Kind: %s, Literal: %q}", t.Kind, t.Literal)
}
