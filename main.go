package goym

import (
	"context"
	"errors"

	"github.com/oklookat/goym/schema"
	"github.com/oklookat/vantuz"
)

const (
	errPrefix = "goym: "
)

var (
	// Нужно обновить токен.
	ErrTokensExpired = errors.New(errPrefix + "tokens expired. You need to refresh tokens")

	//
	ErrNilResponse      = errors.New(errPrefix + "nil http or schema response (???)")
	ErrNilResponseError = errors.New(errPrefix + "nil Response.Error (API changed?)")
	ErrNilStatus        = errors.New(errPrefix + "nil Status (bad auth / API changed?)")
)

// Получить Client для запросов к API.
//
// accessToken - его можно получить выполнив авторизацию в Яндексе.
func New(accessToken string) (*Client, error) {
	httpCl := vantuz.C().SetAuthorization("OAuth " + accessToken)
	cl := &Client{
		Http: httpCl,
	}

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
