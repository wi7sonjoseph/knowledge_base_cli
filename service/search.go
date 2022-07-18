package service

import (
	"context"
	"fmt"
)
func (*Service) Search(ctx context.Context, keyword string) interface{} {
	fmt.Println("Slack search integration code can be further segregated here...: keyword:", keyword)
	return nil
}