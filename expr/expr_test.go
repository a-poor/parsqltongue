package expr_test

import (
	"testing"

	"github.com/a-poor/parsqltongue/expr"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/repr"
)

///////////////////////
// Testing parser... //
///////////////////////

func TestNewParser(t *testing.T) {
	_ = expr.NewParser()
}

////////////////////////////////////
// Testing literals and idents... //
////////////////////////////////////

func TestIntValue(t *testing.T) {
	s := "1"
	ast, err := expr.NewParser().Parse(s)
	if err != nil {
		t.Errorf("Error parsing %q: %v", s, err)
	}
	repr.Println(ast)
}

func TestStringValue(t *testing.T) {
	s := `"test"`
	ast, err := expr.NewParser().Parse(s)
	if err != nil {
		t.Errorf("Error parsing %q: %v", s, err)
	}
	repr.Println(ast)
}

func TestIdentValue(t *testing.T) {
	s := `test`
	ast, err := expr.NewParser().Parse(s)
	if err != nil {
		t.Errorf("Error parsing %q: %v", s, err)
	}
	repr.Println(ast)
}

//////////////////////////
// Testing operation... //
//////////////////////////

func TestCaptureOp(t *testing.T) {
	// Testing errors...
	t.Run("too-few", func(t *testing.T) {
		var o expr.Operation
		err := o.Capture([]string{})
		if err != expr.ErrTooFewValues {
			t.Errorf(
				"Expected error %q, got: %s",
				expr.ErrTooFewValues,
				err,
			)
		}
	})
	t.Run("too-many", func(t *testing.T) {
		var o expr.Operation
		err := o.Capture([]string{"+", "-"})
		if err != expr.ErrTooManyValues {
			t.Errorf(
				"Expected error %q, got: %s",
				expr.ErrTooManyValues,
				err,
			)
		}
	})
	t.Run("unknown-op", func(t *testing.T) {
		var o expr.Operation
		err := o.Capture([]string{"%"})
		if err != expr.ErrUnknownOperation {
			t.Errorf(
				"Expected error %q, got: %s",
				expr.ErrUnknownOperation,
				err,
			)
		}
	})

	// Testing non-errors...
	t.Run("add", func(t *testing.T) {
		var o expr.Operation
		err := o.Capture([]string{"+"})
		if err != nil {
			t.Errorf("Expected %q, got error: %s", "+", err)
		}
		if o != expr.OpAdd {
			t.Errorf(
				"Expected Operation %v, got %v",
				expr.OpAdd,
				o,
			)
		}
	})
	t.Run("subtract", func(t *testing.T) {
		var o expr.Operation
		err := o.Capture([]string{"-"})
		if err != nil {
			t.Errorf("Expected %q, got error: %s", "-", err)
		}
		if o != expr.OpSub {
			t.Errorf(
				"Expected Operation %v, got %v",
				expr.OpSub,
				o,
			)
		}
	})
	t.Run("multiply", func(t *testing.T) {
		var o expr.Operation
		err := o.Capture([]string{"*"})
		if err != nil {
			t.Errorf("Expected %q, got error: %s", "*", err)
		}
		if o != expr.OpMul {
			t.Errorf(
				"Expected Operation %v, got %v",
				expr.OpMul,
				o,
			)
		}
	})
	t.Run("divide", func(t *testing.T) {
		var o expr.Operation
		err := o.Capture([]string{"/"})
		if err != nil {
			t.Errorf("Expected %q, got error: %s", "/", err)
		}
		if o != expr.OpDiv {
			t.Errorf(
				"Expected Operation %v, got %v",
				expr.OpDiv,
				o,
			)
		}
	})

}

////////////////////////////
// Testing expressions... //
////////////////////////////

func TestExpressionAdd(t *testing.T) {
	// Input to test
	s := `1 + 2`

	// Build the parser...
	p := participle.MustBuild(&expr.Expression{})
	var exp expr.Expression

	// Parse the result
	err := p.ParseString("", s, &exp)
	if err != nil {
		t.Errorf("Error parsing %q: %v", s, err)
	}
	repr.Println(exp)
}

func TestExpressionSub(t *testing.T) {
	// Input to test
	s := `1 - 2`

	// Build the parser...
	p := participle.MustBuild(&expr.Expression{})
	var exp expr.Expression

	// Parse the result
	err := p.ParseString("", s, &exp)
	if err != nil {
		t.Errorf("Error parsing %q: %v", s, err)
	}
	repr.Println(exp)
}

//////////////////////////
// Testing functions... //
//////////////////////////

func TestEmptyFunctionValue(t *testing.T) {
	s := `now()`
	p := participle.MustBuild(&expr.FunctionCall{})
	var fn expr.FunctionCall
	err := p.ParseString("", s, &fn)
	if err != nil {
		t.Errorf("Error parsing %q: %v", s, err)
	}
	repr.Println(fn)
}

func TestOneArgFunctionValue(t *testing.T) {
	s := `abs(1)`
	p := participle.MustBuild(&expr.FunctionCall{})
	var fn expr.FunctionCall
	err := p.ParseString("", s, &fn)
	if err != nil {
		t.Errorf("Error parsing %q: %v", s, err)
	}
	repr.Println(fn)
}

func TestMultiArgFunctionValue(t *testing.T) {
	s := `add(1, 2)`
	p := participle.MustBuild(&expr.FunctionCall{})
	var fn expr.FunctionCall
	err := p.ParseString("", s, &fn)
	if err != nil {
		t.Errorf("Error parsing %q: %v", s, err)
	}
	repr.Println(fn)
}
