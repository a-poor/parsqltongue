package parsqltongue

import (
	"github.com/alecthomas/participle/v2"
)

type Parser struct {
	p *participle.Parser
}

func NewParser() *Parser {
	return &Parser{
		p: participle.MustBuild(&SelectStatement{}),
	}
}

func (p *Parser) Parse(stmt string) (*SelectStatement, error) {
	s := &SelectStatement{}
	err := p.p.ParseString("", stmt, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

type SelectStatement struct {
	Select *SelectClause `"SELECT" @@` // Literal, column name, function
	From   *FromClause
	// Where   *WhereClause
	// GroupBy *GroupByClause
	// Having  *HavingClause
	// OrderBy *OrderByClause
	// Limit   *LimitClause
	// Offset  *OffsetClause
}

type SelectClause struct {
	// TableName *string     `@Ident "."`
	Value *LiteralVal `@@`
	// Alias *string     `["AS"] @Ident`
}

type LiteralVal struct {
	Name   *string  `@Ident`
	String *string  `| @String`
	Int    *int     `| @Int`
	Float  *float64 `| @Float`
}

type FromClause struct {
}

type WhereClause struct {
}

type GroupByClause struct {
}

type HavingClause struct {
}

type OrderByClause struct {
}

type LimitClause struct {
}

type OffsetValClause struct {
}
