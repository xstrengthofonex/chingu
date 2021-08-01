package scanner

import "github.com/xstrengthofonex/chingu/token"

type Scanner struct {
	src string
}

func New(src string) *Scanner {
	return &Scanner{src: src}
}


func (s *Scanner) Scan() token.Token {
	return token.EndOfFile()
}
