package session

import (
	"context"
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/data/interfaces/irepo"
	"github.com/lowl11/boostef/ef"
)

func (session *Session[T]) Get(ctx context.Context) ([]T, error) {
	q := session.q.Get()
	ef.DebugPrint(q)

	statement, err := session.connection.PreparexContext(ctx, q)
	if err != nil {
		return nil, err
	}

	rows, err := statement.QueryxContext(ctx, session.args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			// do something...
		}
	}()

	list := make([]T, 0)
	for rows.Next() {
		var item T
		if err = rows.StructScan(&item); err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

func (session *Session[T]) Where(condition func(iquery.Where)) irepo.Session[T] {
	session.q.Where(condition)
	return session
}

func (session *Session[T]) SetPage(page int) irepo.Session[T] {
	session.q.Page(session.pageSize, page)
	return session
}
