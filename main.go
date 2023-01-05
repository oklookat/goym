package goym

import (
	"errors"

	"github.com/oklookat/goym/goymauth"
	"github.com/oklookat/goym/holly"
)

// Получить Client для запросов к API.
//
// Получить tokens можно войдя в аккаунт, используя пакет goymauth.
func New(tokens *goymauth.Tokens) (*Client, error) {
	if tokens == nil {
		return nil, errors.New("nil tokens")
	}

	var cl = &Client{
		self: *holly.AC(tokens.AccessToken),
	}

	// get uid
	resp, err := cl.GetAccountStatus()
	if err != nil {
		return nil, err
	}

	cl.UserId = i2s(resp.Result.Account.UID)
	return cl, err
}

// Клиент для запросов к API.
type Client struct {
	UserId string
	self   holly.Client
}

// Включить вывод HTTP запросов в консоль.
func (c *Client) EnableDevMode() {
	c.self.EnableDevMode()
}

// Отключить вывод HTTP запросов в консоль.
func (c *Client) DisableDevMode() {
	c.self.DisableDevMode()
}
