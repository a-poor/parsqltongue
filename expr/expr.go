package expr

import (
	"github.com/alecthomas/participle/v2"
)

type Parser struct {
	p *participle.Parser
}

func NewParser() *Parser {
	return &Parser{
		participle.MustBuild(&AST{}),
	}
}

func (p *Parser) Parse(s string) (*AST, error) {
	var exp AST
	err := p.p.ParseString("", s, &exp)
	if err != nil {
		return nil, err
	}
	return &exp, nil
}

type AST struct {
	Root *Value `@@`
}

// type Value struct {
// 	Ident   *Ident        `  @@`
// 	Literal *Literal      `| @@`
// 	Func    *FunctionCall `| @@`
// }

type Value struct {
	Ident   *Ident        `  @@`
	Literal *Literal      `| @@`
	Func    *FunctionCall `| @@`
	Expr    *Expression   `| ( "(" @@ ")" )`
}

type Ident struct {
	Name string `@Ident`
}

type Literal struct {
	String *string  `  @String`
	Int    *int     `| @Int`
	Float  *float64 `| @Float`
	Bool   *Boolean `| @("true" | "false")`
}

type Expression struct {
	Left  *Value     `@@`
	Op    *Operation `@( "+" | "-" | "*" | "/" )`
	Right *Value     `@@`
}

type FunctionCall struct {
	FunctionName string   `@Ident`
	Args         []*Value `"(" ( @@ ( "," @@ )* )? ")"`
}
