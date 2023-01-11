package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить информацию об аккаунте.
func (c Client) GetAccountStatus(ctx context.Context) (*schema.Status, error) {
	// GET /account/status
	var endpoint = genApiPath([]string{"account", "status"})

	var data = &schema.TypicalResponse[*schema.Status]{}
	resp, err := c.self.R().SetResult(data).SetError(data).Get(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data.Result, err
}
