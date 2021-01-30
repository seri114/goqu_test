package main

import (
	"fmt"
	"goqu_test/visitor"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

// composi
type Composite struct {
}

func main() {

	ons := make([]exp.Expression, 0)
	ons = append(ons, goqu.T("A").Col("a").Eq(goqu.T("B").Col("b")))
	ons = append(ons, goqu.T("A").Col("a").Eq(goqu.T("B").Col("b")))

	var aa exp.ExpressionList
	aa = goqu.And(goqu.Ex{
		"a": "b",
	})

	sql, _, _ := goqu.From("test").Where(aa).
		GroupBy(goqu.T("A").Col("a")).
		LeftJoin(goqu.T("A"), goqu.On(ons...)).
		ToSQL()

	// sql, _, _ = goqu.Expression(goqu.T("a").Eq(3))

	fmt.Println(sql)

	sql, _, _ = goqu.From("AA").Where(visitor.TestVisitor()).ToSQL()
	// sql, _, _ = goqu.From("AA").Where(
	// 	goqu.And(goqu.And(goqu.T("a").Eq(3), goqu.T("a").Eq(3)),
	// 		goqu.And(goqu.T("a").Eq(3), goqu.T("a").Eq(3)))).
	// 	ToSQL()
	fmt.Println(sql)

}
