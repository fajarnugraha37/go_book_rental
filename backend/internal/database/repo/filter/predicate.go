package filter

import (
	"github.com/uptrace/bun"
)

type Predicate struct {
	Operation        OperationType
	Filters          []QueryFilter
	Pageable         *Pageable
	Sortable         *[]Sortable
	Relations        []string
	RelationsWithOpt []struct {
		Name string
		Opt  bun.RelationOpts
	}
	Joins []Join
}

func (predicate Predicate) ToQuery() [](func(*bun.SelectQuery) *bun.SelectQuery) {
	return [](func(*bun.SelectQuery) *bun.SelectQuery){
		predicate.toRelations(),
		predicate.toJoins(),
		predicate.toFilters(),
	}
}

func (predicate Predicate) toJoins() func(*bun.SelectQuery) *bun.SelectQuery {
	return func(mainQuery *bun.SelectQuery) *bun.SelectQuery {
		for _, join := range predicate.Joins {
			mainQuery.Join(join.Clause)
			for _, on := range join.Ons {
				mainQuery.JoinOn(on)
			}
		}

		return mainQuery
	}
}

func (predicate Predicate) toRelations() func(*bun.SelectQuery) *bun.SelectQuery {
	return func(mainQuery *bun.SelectQuery) *bun.SelectQuery {
		for _, relation := range predicate.Relations {
			mainQuery.Relation(relation)
		}
		for _, relation := range predicate.RelationsWithOpt {
			mainQuery.RelationWithOpts(relation.Name, relation.Opt)
		}
		return mainQuery
	}
}

func (predicate Predicate) toFilters() func(*bun.SelectQuery) *bun.SelectQuery {
	return func(mainQuery *bun.SelectQuery) *bun.SelectQuery {
		for _, filter := range predicate.Filters {
			expression, arguments := filter.Comparator.ToExpression(filter.Field, filter.Value)
			if predicate.Operation.IsOr() {
				mainQuery.WhereOr(expression, arguments...)
			} else {
				mainQuery.Where(expression, arguments...)
			}

			if filter.Predicate != nil {
				mainQuery.WhereGroup(
					predicate.Operation.ToOperation(),
					func(predicateQuery *bun.SelectQuery) *bun.SelectQuery {

						return filter.Predicate.toFilters()(predicateQuery)
					},
				)
			}
		}

		return mainQuery
	}
}

// v2
// return func(mainQuery *bun.SelectQuery) *bun.SelectQuery {
// 	mainQuery.WhereGroup(
// 		predicate.Operation.ToOperation(),
// 		func(filterQuery *bun.SelectQuery) *bun.SelectQuery {
// 			for _, filter := range predicate.Filters {
// 				expression, arguments := filter.Comparator.ToExpression(filter.Field, filter.Value)
// 				if predicate.Operation.IsOr() {
// 					filterQuery.WhereOr(expression, arguments...)
// 				} else {
// 					filterQuery.Where(expression, arguments...)
// 				}

// 				if filter.Predicate != nil {
// 					filterQuery.WhereGroup(
// 						predicate.Operation.ToOperation(),
// 						func(predicateQuery *bun.SelectQuery) *bun.SelectQuery {

// 							return filter.Predicate.ToQuery()(predicateQuery)
// 						},
// 					)
// 				}
// 			}

// 			return filterQuery
// 		},
// 	)

// 	return mainQuery
// }

// v1
// return func(q *bun.SelectQuery) *bun.SelectQuery {
// 	for _, filter := range predicate.Filters {
// 		expression, arguments := filter.Comparator.ToExpression(filter.Field, filter.Value)
// 		if predicate.Operation.IsOr() {
// 			q.WhereOr(expression, arguments...)
// 		} else {
// 			q.Where(expression, arguments...)
// 		}

// 		if filter.Predicate != nil {
// 			q.WhereGroup(
// 				filter.Predicate.Operation.ToOperation(),
// 				func(queryPerdicte *bun.SelectQuery) *bun.SelectQuery {
// 					return filter.Predicate.ToQuery()(queryPerdicte)
// 				},
// 			)
// 		}
// 	}

// 	return q
// }
