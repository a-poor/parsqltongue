package calc

import "github.com/alecthomas/participle/v2"

type Parser struct {
	p *participle.Parser
}

func NewParser() *Parser {
	return &Parser{
		participle.MustBuild(&Expr{}),
	}
}

func (p *Parser) Parse(s string) (*Expr, error) {
	return nil, nil
}

type Expr struct {
	Left       Value `@@`
	OpAndRight []*struct {
		Op    Operation `@( "+" | "-" | "*" | "/" )`
		Right Value     `@@`
	} `@@*`
}

type Value struct {
	Int *int `@Int`
}
