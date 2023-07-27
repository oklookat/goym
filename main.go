package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
	"github.com/oklookat/vantuz"
)

// Получить Client для запросов к API.
//
// accessToken - его можно получить выполнив авторизацию в Яндексе.
func New(accessToken string) (*Client, error) {
	httpCl := vantuz.C().SetAuthorization("OAuth " + accessToken)
	cl := &Client{
		Http: httpCl,
	}
	cl.SetUserAgent("goym")

	status, err := cl.AccountStatus(context.Background())
	if err != nil {
		return nil, err
	}
	cl.UserId = status.Result.Account.UID

	return cl, err
}

// Клиент для запросов к API.
type Client struct {
	// ID текущего пользователя.
	UserId schema.ID

	// Отправляет запросы.
	Http *vantuz.Client
}

// Установить user agent для запросов.
func (c *Client) SetUserAgent(val string) {
	c.Http.SetUserAgent(val)
}
