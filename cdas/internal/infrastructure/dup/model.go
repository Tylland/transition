package dup

type Expression interface {
	isExpression()
}

type Comparison struct {
	Operator string
	Name     string
	Value    any
}

type Logical struct {
	Operator string
	Children []Expression
}

func (Comparison) isExpression() {}
func (Logical) isExpression()    {}
