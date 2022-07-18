package service

import (
	"context"
	"fmt"
)
func (*Service) Web(ctx context.Context, keyword string) interface{} {
	fmt.Println("Slack web integration code can be further segregated here..., keyword:", keyword)
	return nil
}