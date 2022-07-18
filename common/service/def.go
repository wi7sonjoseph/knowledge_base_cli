package service

import (
	"context"
	"knowledge_base_cli/common/data"
)
// BaseService model
type BaseService struct {
	DBSession *data.DBSession
}

// New will return new service instance
func New(ctx context.Context) (*BaseService, error) {

	ctx = context.WithValue(ctx, "AUDIT_DB", "audit-db")
	dbSession, err := data.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	return &BaseService{DBSession: dbSession}, nil
}
