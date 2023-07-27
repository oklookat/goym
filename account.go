package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить информацию об аккаунте.
func (c Client) AccountStatus(ctx context.Context) (schema.Response[schema.Status], error) {
	// GET /account/status
	endpoint := genApiPath("account", "status")

	data := &schema.Response[schema.Status]{}
	resp, err := c.Http.R().SetResult(data).SetError(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return *data, err
}

// Получить настройки аккаунта.
func (c Client) AccountSettings(ctx context.Context) (schema.Response[schema.AccountSettings], error) {
	// GET /account/settings
	endpoint := genApiPath("account", "settings")
	data := &schema.Response[schema.AccountSettings]{}

	resp, err := c.Http.R().SetResult(data).SetError(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return *data, err
}
