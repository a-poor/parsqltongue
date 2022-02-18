package parsqltongue_test

import (
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
}
