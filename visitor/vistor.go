package visitor

import (
	"goqu_test/model"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type VisitorExpression struct {
	model.Visitor
	list exp.ExpressionList
}

func (self *VisitorExpression) VisitAdd(f *model.CompositeAdd) int {
	return f.Lhs.ToInt(self) + f.Rhs.ToInt(self)
}
func (self *VisitorExpression) VisitNumber(f *model.CompositeNumber) int {
	return f.Val
}
func (self *VisitorExpression) VisitMultiply(f *model.CompositeMultiply) int {
	return f.Lhs.ToInt(self) * f.Rhs.ToInt(self)
}
func (self *VisitorExpression) VisitAnd(f *model.CompositeAnd) exp.ExpressionList {
	return goqu.And(f.Lhs.Accept(self), f.Rhs.Accept(self))
}
func (self *VisitorExpression) VisitOr(f *model.CompositeOr) exp.ExpressionList {
	return goqu.Or(f.Lhs.Accept(self), f.Rhs.Accept(self))
}
func (self *VisitorExpression) VisitEq(f *model.CompositeEq) exp.ExpressionList {
	return goqu.And(f.Lhs.Eq(f.Rhs.ToInt(self)))
}

func TestVisitor() exp.ExpressionList {
	// aa := &CompositeEq{
	// 	Lhs: goqu.T("TestTable"),
	// 	Rhs: &CompositeNumber{val: 1},
	// }
	var aa model.Composite
	aa = &model.CompositeEq{
		Lhs: goqu.T("TestTable"),
		Rhs: &model.CompositeAdd{
			Lhs: &model.CompositeMultiply{
				Lhs: &model.CompositeNumber{Val: 1},
				Rhs: &model.CompositeNumber{Val: 2},
			},
			Rhs: &model.CompositeNumber{Val: 2},
		},
	}

	aa = &model.CompositeAnd{
		Lhs: aa,
		Rhs: aa,
	}
	aa = &model.CompositeOr{
		Lhs: aa,
		Rhs: aa,
	}

	f := model.Filter{
		Composite: aa,
	}
	ret := f.Composite.Accept(&VisitorExpression{})
	return ret
}
