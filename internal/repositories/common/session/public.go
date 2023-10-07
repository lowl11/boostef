package session

import "github.com/lowl11/boostef/data/interfaces/irepository"

func (session *Session) SetPageSize(pageSize int) irepository.Session {
	session.pageSize = pageSize
	return session
}
