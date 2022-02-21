package parsqltongue_test

import (
	"reflect"
	"testing"

	parsql "github.com/a-poor/parsqltongue"
	"github.com/alecthomas/repr"
)

func TestNothing(t *testing.T) {
	t.Log("Nothing worked!")
}

func TestCreateParser(t *testing.T) {
	_ = parsql.NewParser()
}

func TestParseSelect(t *testing.T) {
	p := parsql.NewParser()

	// int
	s := "SELECT 1"
	stmt, err := p.Parse(s)
	if err != nil {
		t.Errorf("Error parsing SELECT statement %q: %v", s, err)
	}
	repr.Println(stmt)

	// col name
	s = "SELECT username"
	stmt, err = p.Parse(s)
	if err != nil {
		t.Errorf("Error parsing SELECT statement %q: %v", s, err)
	}
	repr.Println(stmt)

	// string
	s = "SELECT \"test\""
	stmt, err = p.Parse(s)
	if err != nil {
		t.Errorf("Error parsing SELECT statement %q: %v", s, err)
	}
	repr.Println(stmt)

	// float
	s = "SELECT 1."
	stmt, err = p.Parse(s)
	if err != nil {
		t.Errorf("Error parsing SELECT statement %q: %v", s, err)
	}
	repr.Println(stmt)

	// star
	s = "SELECT *"
	stmt, err = p.Parse(s)
	if err != nil {
		t.Errorf("Error parsing SELECT statement %q: %v", s, err)
	}
	star := "*"
	expect := &parsql.SelectStatement{
		Select: &parsql.SelectClause{
			Value: &parsql.Value{
				Star: &star,
			},
		},
	}
	if !reflect.DeepEqual(stmt, expect) {
		t.Logf("Error parsing %q", s)
		repr.Println(stmt)
		t.Fail()
	}
}

func TestSelectAlias(t *testing.T) {
	p := parsql.NewParser()

	// alias with "AS"
	s := "SELECT 1 AS a_number"
	stmt, err := p.Parse(s)
	if err != nil {
		t.Errorf("Error parsing SELECT statement %q: %v", s, err)
	}
	repr.Println(stmt)

	// alias no "AS"
	s = "SELECT username user"
	stmt, err = p.Parse(s)
	if err != nil {
		t.Errorf("Error parsing SELECT statement %q: %v", s, err)
	}
	repr.Println(stmt)
}

func TestSelectTableName(t *testing.T) {
	p := parsql.NewParser()

	// unquoted table name
	s := "SELECT users.username"
	stmt, err := p.Parse(s)
	if err != nil {
		t.Errorf("Error parsing SELECT statement %q: %v", s, err)
	}
	repr.Println(stmt)
}

func TestSelectExpr(t *testing.T) {
	p := parsql.NewParser()

	s := "SELECT 1 + 1"
	stmt, err := p.Parse(s)
	if err != nil {
		t.Errorf("Error parsing SELECT statement %q: %v", s, err)
	}
	repr.Println(stmt)
}
