package service

import (
	"context"
	"knowledge_base_cli/common/service"
)

// Service model
type Service struct {
	*service.BaseService
}

// New will return new service instance
func New(ctx context.Context) (*Service, error) {
	baseService, err := service.New(ctx)
	return &Service{BaseService: baseService}, err
}
