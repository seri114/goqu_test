package filter

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type Filter struct {
	Composite Composite
}

type Composite interface {
	ToInt() int
	ToExpression() exp.ExpressionList
}

type CompositeAdd struct {
	Composite
	Lhs Composite
	Rhs Composite
}

func (self *CompositeAdd) ToInt() int { return self.Lhs.ToInt() + self.Rhs.ToInt() }

type CompositeMultiply struct {
	Composite
	Lhs Composite
	Rhs Composite
}

func (self *CompositeMultiply) ToInt() int { return self.Lhs.ToInt() * self.Rhs.ToInt() }

type CompositeNumber struct {
	Composite
	val int
}

func (self *CompositeNumber) ToInt() int { return self.val }

type CompositeAnd struct {
	Composite
	Lhs Composite
	Rhs Composite
}

func (self *CompositeAnd) ToExpression() exp.ExpressionList {
	return goqu.And(self.Lhs.ToExpression(), self.Rhs.ToExpression())
}

type CompositeOr struct {
	Composite
	Lhs Composite
	Rhs Composite
}

func (self *CompositeOr) ToExpression() exp.ExpressionList {
	return goqu.Or(self.Lhs.ToExpression(), self.Rhs.ToExpression())
}

type CompositeEq struct {
	Composite
	Lhs exp.IdentifierExpression
	Rhs Composite
}

func (self *CompositeEq) ToExpression() exp.ExpressionList {
	return goqu.And(self.Lhs.Eq(self.Rhs.ToInt()))
}

func TestFilter() exp.ExpressionList {
	// aa := &CompositeEq{
	// 	Lhs: goqu.T("TestTable"),
	// 	Rhs: &CompositeNumber{val: 1},
	// }
	var aa Composite
	aa = &CompositeEq{
		Lhs: goqu.T("TestTable"),
		Rhs: &CompositeMultiply{
			Lhs: &CompositeAdd{
				Lhs: &CompositeNumber{val: 1},
				Rhs: &CompositeNumber{val: 2},
			},
			Rhs: &CompositeNumber{
				val: 2,
			},
		},
	}
	aa = &CompositeAnd{
		Lhs: aa,
		Rhs: aa,
	}
	aa = &CompositeOr{
		Lhs: aa,
		Rhs: aa,
	}

	f := Filter{
		Composite: aa,
	}
	ret := f.Composite.ToExpression()
	return ret
}
