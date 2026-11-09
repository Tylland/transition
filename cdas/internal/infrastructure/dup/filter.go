package dup

import (
	"fmt"
	"strings"
)

func EQ(name string, value any) Expression {
	return Comparison{
		Operator: "EQ",
		Name:     name,
		Value:    value,
	}
}

func NE(name string, value any) Expression {
	return Comparison{
		Operator: "NE",
		Name:     name,
		Value:    value,
	}
}

func GT(name string, value any) Expression {
	return Comparison{
		Operator: "GT",
		Name:     name,
		Value:    value,
	}
}

func GTE(name string, value any) Expression {
	return Comparison{
		Operator: "GTE",
		Name:     name,
		Value:    value,
	}
}

func LT(name string, value any) Expression {
	return Comparison{
		Operator: "LT",
		Name:     name,
		Value:    value,
	}
}

func LTE(name string, value any) Expression {
	return Comparison{
		Operator: "LTE",
		Name:     name,
		Value:    value,
	}
}

func EXISTS(name string) Expression {
	return Comparison{
		Operator: "EXISTS",
		Name:     name,
		Value:    true,
	}
}

func NOTEXISTS(name string) Expression {
	return Comparison{
		Operator: "EXISTS",
		Name:     name,
		Value:    false,
	}
}

func LIKE(name string, pattern string) Expression {
	return Comparison{
		Operator: "LIKE",
		Name:     name,
		Value:    pattern,
	}
}

func NOTLIKE(name string, pattern string) Expression {
	return Comparison{
		Operator: "NOTLIKE",
		Name:     name,
		Value:    pattern,
	}
}

func IN(name string, values ...any) Expression {
	parts := make([]string, len(values))

	for i, value := range values {
		parts[i] = fmt.Sprint(value)
	}

	return Comparison{
		Operator: "IN",
		Name:     name,
		Value:    strings.Join(parts, ","),
	}
}

func NOTIN(name string, values ...any) Expression {
	parts := make([]string, len(values))

	for i, value := range values {
		parts[i] = fmt.Sprint(value)
	}

	return Comparison{
		Operator: "NOTIN",
		Name:     name,
		Value:    strings.Join(parts, ","),
	}
}

func AND(expressions ...Expression) Expression {
	return Logical{
		Operator: "and",
		Children: expressions,
	}
}

func OR(expressions ...Expression) Expression {
	return Logical{
		Operator: "or",
		Children: expressions,
	}
}

func NOT(expression Expression) Expression {
	return Logical{
		Operator: "not",
		Children: []Expression{expression},
	}
}
