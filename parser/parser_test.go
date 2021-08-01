package parser

import (
	"testing"

	"github.com/xstrengthofonex/chingu/ast"
)

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
