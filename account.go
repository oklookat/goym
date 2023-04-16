package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить информацию об аккаунте.
func (c Client) AccountStatus(ctx context.Context) (*schema.Status, error) {
	// GET /account/status
	endpoint := genApiPath("account", "status")

	data := &schema.Response[*schema.Status]{}
	resp, err := c.Http.R().SetResult(data).SetError(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return data.Result, err
}

// Активировать промокод.
//
// Метод не тестировался.
func (c Client) AccountConsumePromocode(ctx context.Context, code string, language string) (*schema.PromocodeStatus, error) {
	// POST /account/consume-promo-code
	endpoint := genApiPath("account", "consume-promo-code")

	body := schema.AccountConsumePromocodeRequestBody{
		Code:     code,
		Language: language,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	data := &schema.Response[*schema.PromocodeStatus]{}
	resp, err := c.Http.R().SetResult(data).SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return data.Result, err
}

// Получить настройки аккаунта.
func (c Client) AccountSettings(ctx context.Context) (*schema.AccountSettings, error) {
	// GET /account/settings
	endpoint := genApiPath("account", "settings")

	data := &schema.Response[*schema.AccountSettings]{}
	resp, err := c.Http.R().SetResult(data).SetError(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return data.Result, err
}

// Изменить настройки аккаунта.
//
// Настройку нельзя изменить, если в поле AccountSettings есть url:"-".
//
// Может вернуть как AccountSettings, так и ничего.
func (c Client) ChangeAccountSettings(ctx context.Context, set schema.AccountSettings) (any, error) {
	// POST /account/settings
	// todo(?) иногда выдает ошибку json unmarshal.

	endpoint := genApiPath("account", "settings")

	body := schema.ChangeAccountSettingsRequestBody{}
	body.Change(set)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	data := &schema.Response[any]{}
	resp, err := c.Http.R().SetResult(data).SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return data.Result, err
}
