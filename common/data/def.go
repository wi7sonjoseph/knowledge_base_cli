package data

import (
	"sync"
	"context"
)

var (
	dbSession *DBSession
	once      sync.Once
)

//DBSession contains all database sessions
type DBSession struct {
}



// GetConn will return db session, which is created only once and then same connection pool is reused everytime.
func GetConn(ctx context.Context) (dbSession *DBSession, err error) {

	once.Do(func() {
		dbSession = new(DBSession)
		
	})
	return
}
