package goym

import (
	"errors"

	"github.com/oklookat/goym/goymauth"
	"github.com/oklookat/goym/vantuz"
)

// Получить Client для запросов к API.
//
// Получить tokens можно войдя в аккаунт, используя пакет goymauth.
func New(tokens *goymauth.Tokens) (*Client, error) {
	if tokens == nil {
		return nil, errors.New("nil tokens")
	}

	var cl = &Client{
		self: *vantuz.AC(tokens.AccessToken),
	}

	// get uid
	resp, err := cl.GetAccountStatus()
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("nil response")
	}

	var result = resp.Result
	if result.Account == nil {
		return nil, errors.New("nil account")
	}

	cl.UserId = resp.Result.Account.UID
	cl.userId = i2s(cl.UserId)
	return cl, err
}

// Клиент для запросов к API.
type Client struct {
	UserId int64

	// Для создания эндпоинтов.
	userId string
	self   vantuz.Client
}

// Включить вывод HTTP запросов в консоль.
func (c *Client) EnableDevMode() {
	c.self.EnableDevMode()
}

// Отключить вывод HTTP запросов в консоль.
func (c *Client) DisableDevMode() {
	c.self.DisableDevMode()
}
