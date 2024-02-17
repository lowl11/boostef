package join

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"strings"
)

func (join *Join) Get(_ ...string) string {
	query := strings.Builder{}

	query.WriteString(join.joinType)
	query.WriteString(join.table)
	query.WriteString(" AS ")
	query.WriteString(join.alias)
	query.WriteString(" ON ")
	query.WriteString("(")
	query.WriteString(join.joinColumn)
	query.WriteString(" = ")
	query.WriteString(join.mainColumn)
	query.WriteString(")")

	return query.String()
}

func (join *Join) Table(tableName string) iquery.Join {
	join.table = tableName
	return join
}

func (join *Join) Alias(aliasName string) iquery.Join {
	join.alias = aliasName
	return join
}

func (join *Join) JoinColumn(column string) iquery.Join {
	before, after, found := strings.Cut(column, ".")
	if found {
		join.joinColumn = "\"" + before + "\"." + after
	} else {
		join.joinColumn = column
	}
	return join
}

func (join *Join) MainColumn(column string) iquery.Join {
	before, after, found := strings.Cut(column, ".")
	if found {
		join.mainColumn = "\"" + before + "\"." + after
	} else {
		join.mainColumn = column
	}
	return join
}
