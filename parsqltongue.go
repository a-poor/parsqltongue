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

type SelectStatement struct {
	Select  []*SelectVal
	From    []*FromVal
	Where   []*WhereVal
	GroupBy []*GroupByVal
	Having  []*HavingVal
	OrderBy []*OrderByVal
	Limit   []*LimitVal
	Offset  []*OffsetVal
}

type SelectVal struct {
}

type FromVal struct {
}

type WhereVal struct {
}

type GroupByVal struct {
}

type HavingVal struct {
}

type OrderByVal struct {
}

type LimitVal struct {
}

type OffsetVal struct {
}
