package basic

import (
	"context"
	"defaas/core/helper"
)

func (client *BasicClient) FaaSTokenEventSub(ctx context.Context) error {

	return helper.FaaSTokenEventSub(ctx, client.FaaSToken.Contract)
}
