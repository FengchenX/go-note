package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type PageFilter struct {
	Sort   string
	Order  string
	Limit  string
	Offset string
}

func ParsePageFilter(c *gin.Context) (*PageFilter, error) {
	var filter PageFilter
	e := c.ShouldBindQuery(&filter)
	return &filter, e
}

func PageFilterSql(sql, keyCol string, filter *PageFilter) string {
	if len(sql) == 0 || filter == nil {
		return sql
	}
	if len(filter.Sort) > 0 && len(filter.Order) == 0 {
		filter.Order = "DESC"
	}
	if len(filter.Sort) > 0 && filter.Sort != keyCol {
		sql = fmt.Sprintf("%s ORDER BY %s %s", sql, filter.Sort, filter.Order)
		if len(keyCol) > 0 {
			sql = fmt.Sprintf("%s, %s %s", sql, keyCol, filter.Order)
		}
	} else if len(keyCol) > 0 {
		sql = fmt.Sprintf("%s ORDER BY %s %s", sql, keyCol, filter.Order)
	}
	if len(filter.Offset) > 0 && len(filter.Limit) > 0 {
		sql = fmt.Sprintf("%s LIMIT %s OFFSET %s", sql, filter.Limit, filter.Offset)
	}
	return sql
}
