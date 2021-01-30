package model

import (
	"github.com/doug-martin/goqu/v9/exp"
)

type Visitor interface {
	VisitAdd(*CompositeAdd) int
	VisitNumber(*CompositeNumber) int
	VisitMultiply(*CompositeMultiply) int
	VisitAnd(*CompositeAnd) exp.ExpressionList
	VisitOr(*CompositeOr) exp.ExpressionList
	VisitEq(*CompositeEq) exp.ExpressionList
}

type Filter struct {
	Composite Composite
}

type Composite interface {
	Accept(Visitor) exp.ExpressionList
	ToInt(Visitor) int
}

type CompositeAdd struct {
	Composite
	Lhs Composite
	Rhs Composite
}

func (self *CompositeAdd) ToInt(v Visitor) int {
	return v.VisitAdd(self)
}

type CompositeMultiply struct {
	Composite
	Lhs Composite
	Rhs Composite
}

func (self *CompositeMultiply) ToInt(v Visitor) int { return v.VisitMultiply(self) }

type CompositeNumber struct {
	Composite
	Val int
}

func (self *CompositeNumber) ToInt(v Visitor) int { return v.VisitNumber(self) }

type CompositeAnd struct {
	Composite
	Lhs Composite
	Rhs Composite
}

func (self *CompositeAnd) Accept(v Visitor) exp.ExpressionList {
	return v.VisitAnd(self)
}

type CompositeOr struct {
	Composite
	Lhs Composite
	Rhs Composite
}

func (self *CompositeOr) Accept(v Visitor) exp.ExpressionList {
	return v.VisitOr(self)
}

type CompositeEq struct {
	Composite
	Lhs exp.IdentifierExpression
	Rhs Composite
}

func (self *CompositeEq) Accept(v Visitor) exp.ExpressionList {
	return v.VisitEq(self)
}
