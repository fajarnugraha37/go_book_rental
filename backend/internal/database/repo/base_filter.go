package repo

import (
	"backend/pkg/helper"
	"reflect"

	"github.com/uptrace/bun"
)

const (
	tagKeyComparator = "comparator"
	tagKeyField      = "field"
)

type AuditFilter struct {
	ID  *string   `comparator:"eq" field:"id"`
	IDs *[]string `comparator:"in" field:"id"`

	CreatedBy       *string `comparator:"like" field:"created_by"`
	CreatedAt       *string `comparator:"eq" field:"created_at"`
	CreatedAtBefore *string `comparator:"le" field:"created_at"`
	CreatedAtAfter  *string `comparator:"ge" field:"created_at"`

	UpdatedBy       *string `comparator:"like" field:"updated_by"`
	UpdatedAt       *string `comparator:"eq" field:"updated_at"`
	UpdatedAtBefore *string `comparator:"le" field:"updated_at"`
	UpdatedAtAfter  *string `comparator:"ge" field:"updated_at"`

	DeletedFlag     *bool   `comparator:"eq" field:"deleted_flag"`
	DeletedBy       *string `comparator:"like" field:"deleted_by"`
	DeletedAt       *string `comparator:"eq" field:"deleted_at"`
	DeletedAtBefore *string `comparator:"le" field:"deleted_at"`
	DeletedAtAfter  *string `comparator:"ge" field:"deleted_at"`
}

func filterToQueryBuilder[TFilter any](filter TFilter) func(q bun.QueryBuilder) bun.QueryBuilder {
	return func(q bun.QueryBuilder) bun.QueryBuilder {
		fields := helper.GetAllFields(
			reflect.TypeOf(filter),
			reflect.ValueOf(filter),
		)
		for _, field := range fields {
			if field.Value == nil {
				continue
			}

			tagValue, ok := field.Tag.Lookup(tagKeyComparator)
			if !ok {
				continue
			}
			tagField, ok := field.Tag.Lookup(tagKeyField)
			if !ok {
				tagField = helper.ToSnakeCase(field.Name)
			}

			switch tagValue {
			case "eq":
				{
					q.Where(tagField+" = ? ", field.Value)
				}
			case "neq":
				{
					q.Where(tagField+" != ? ", field.Value)
				}
			case "lt":
				{
					q.Where(tagField+" < ? ", field.Value)
				}
			case "lte":
				{
					q.Where(tagField+" <= ? ", field.Value)
				}
			case "gt":
				{
					q.Where(tagField+" > ? ", field.Value)
				}
			case "gte":
				{
					q.Where(tagField+" >= ? ", field.Value)
				}
			case "like":
				{
					q.Where(tagField+" LIKE ? ", field.Value)
				}
			case "nlike":
				{
					q.Where(tagField+" NOT LIKE ? ", field.Value)
				}
			case "in":
				{
					q.Where(tagField+" IN (?) ", bun.In(field.Value))
				}
			case "nin":
				{
					q.Where(tagField+" NOT IN (?) ", bun.In(field.Value))
				}
			}
		}

		return q
	}
}
