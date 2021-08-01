package scanner

import (
	"github.com/xstrengthofonex/chingu/token"
)

type Scanner struct {
	src   string
	pos   int
	start int
}

func New(src string) *Scanner {
	return &Scanner{src: src}
}

func (s *Scanner) eof() bool {
	return s.pos >= len(s.src)
}

func (s *Scanner) advance() rune {
	if s.eof() {
		return '\x00'
	}
	s.pos++
	return rune(s.src[s.pos-1])
}

func (s *Scanner) peek() rune {
	if s.eof() {
		return '\x00'
	}
	return rune(s.src[s.pos])
}

func (s *Scanner) Scan() token.Token {
	if s.eof() {
		return token.EndOfFile()
	}
	s.start = s.pos
	c := s.advance()
	switch c {
	case ' ', '\n', '\t':
		return s.Scan()
	case '(':
		return token.New(token.LPAREN, "(")
	case ')':
		return token.New(token.RPAREN, ")")
	case '.':
		return token.New(token.DOT, ".")
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
	default:
		if isInitialIdent(c) {
			for isSubsequentIdent(s.peek()) {
				s.advance()
			}
			return token.New(token.IDENT, s.src[s.start:s.pos])
		} else {
			return token.New(token.BAD, s.src[s.start:s.pos])
		}
	}
}

func isAlpha(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isInitialIdent(c rune) bool {
	return isAlpha(c) || c == '_'
}

func isSubsequentIdent(c rune) bool {
	return isInitialIdent(c) || isDigit(c)
}
