package parsqltongue

import (
	"errors"

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
	// From   *FromClause   `"FROM" @@`
	// Where   *WhereClause
	// GroupBy *GroupByClause
	// Having  *HavingClause
	// OrderBy *OrderByClause
	// Limit   *LimitClause
	// Offset  *OffsetClause
}

type SelectClause struct {
	TableName *string `[@Ident "."]`
	Value     *Value  `@@`
	Expr      *Expr   `| @@`
	Alias     *string `[["AS"] @Ident]`
}

type Operation int

const (
	OpAdd Operation = iota
	OpSub
	OpMul
	OpDiv
)

var opMap = map[string]Operation{
	"+": OpAdd,
	"-": OpSub,
	"*": OpMul,
	"/": OpDiv,
}

func (o *Operation) Capture(s []string) error {
	if len(s) != 1 {
		return errors.New("too many or too few operations supplied")
	}

	op, ok := opMap[s[0]]
	if !ok {
		return errors.New("unknown operation")
	}
	o = &op
	return nil
}

type Expr struct {
	Left  *Value     `@@`
	Op    *Operation `@("+" | "-" | "*" | "/")`
	Right *Value     `@@`
}

type Func struct {
}

type FuncArg struct {
	Val *Value `@@`
	// Func *Func  `@@`
}

type Value struct {
	Star   *string  `@"*"`
	Name   *string  `| @Ident`
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
