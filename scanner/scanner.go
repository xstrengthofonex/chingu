package scanner

import (
	"github.com/xstrengthofonex/chingu/token"
)

type Scanner struct {
	src string
	pos int
}

func New(src string) *Scanner {
	return &Scanner{src: src}
}

func (s *Scanner) eof() bool {
	return s.pos >= len(s.src)
}

func (s *Scanner) advance() {
	if !s.eof() {
		s.pos ++ 
	}
}

func (s *Scanner) peek() rune {
	if s.eof() {
		return '\x00'
	}
	return rune(s.src[s.pos])
}

func (s *Scanner) Scan() token.Token {
	s.advance()
	switch s.peek() {
	case ' ', '\n', '\t':
		return s.Scan()
	case '(':
		return token.New(token.LPAREN, "(")
	case ')':
		return token.New(token.RPAREN, ")")
	case '+':
		return token.New(token.PLUS, "+")
	case '-':
		return token.New(token.MINUS, "-")
	case '*':
		return token.New(token.STAR, "*")
	case '/':
		return token.New(token.SLASH, "/")
	case '`':
		return token.New(token.TICK, "`")
	}
	return token.EndOfFile()
}
