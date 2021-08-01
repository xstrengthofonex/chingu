package ast

import (
	"fmt"
	"strings"
)

type Expr interface {
	Equals(Expr) bool
	String() string
}

type Symbol struct {
	Value string
}

func (s1 *Symbol) Equals(other Expr) bool {
	if other == nil {
		return false
	} else if s2, ok := other.(*Symbol); ok {
		return s1.Value == s2.Value
	}
	return false
}

func (s *Symbol) String() string {
	return s.Value
}

type Call struct {
	Operator Expr
	Operands []Expr
}

func (c1 *Call) Equals(other Expr) bool {
	if other == nil {
		return false
	} else if c2, ok := other.(*Call); ok {
		if !c1.Operator.Equals(c2.Operator) {
			return false
		}

		if len(c1.Operands) != len(c2.Operands) {
			return false
		}

		for i := range c1.Operands {
			if !c1.Operands[i].Equals(c2.Operands[i]) {
				return false
			}
		}
	}
	return true
}

func (s *Call) String() string {
	operands := []string{}
	for _, operand := range s.Operands {
		operands = append(operands, operand.String())
	}
	return fmt.Sprintf("(%s %s)", s.Operator, strings.Join(operands, " "))
}


type Fn struct {
	Params []*Symbol
	Body Expr
}

func (f1 *Fn) Equals(other Expr) bool {
	if other == nil {
		return false
	} else if f2, ok := other.(*Fn); ok {
		if len(f1.Params) != len(f2.Params) {
			return false
		}

		for i := range f1.Params {
			if !f1.Params[i].Equals(f2.Params[i]) {
				return false
			}
		}

		if !f1.Body.Equals(f2.Body) {
			return false
		}
	}
	return true
}

func (f *Fn) String() string {
	params := []string{}
	for _, param := range f.Params {
		params = append(params, param.String())
	}
	return fmt.Sprintf("fn %s . %s", strings.Join(params, " "), f.Body)
}