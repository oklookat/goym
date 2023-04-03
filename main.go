package goym

import (
	"context"
	"errors"

	"github.com/oklookat/goym/auth"
	"github.com/oklookat/goym/schema"
	"github.com/oklookat/goym/vantuz"
)

const (
	errPrefix = "goym: "
)

var (
	// Токены авторизации истекли. Нужно вызывать Tokens.Refresh.
	ErrTokensExpired = errors.New(errPrefix + "tokens expired. You need to refresh your current tokens")

	//
	ErrNilResponse      = errors.New(errPrefix + "nil http or schema response (dev error?)")
	ErrNilResponseError = errors.New(errPrefix + "nil Response.Error (API changed?)")
	ErrNilStatus        = errors.New(errPrefix + "nil Status (bad auth or API changed?)")
	ErrNilAccount       = errors.New(errPrefix + "nil Status.Account (API changed?)")
)

// Получить Client для запросов к API.
//
// Получить tokens можно войдя в аккаунт, используя пакет auth.
func New(tokens *auth.Tokens) (*Client, error) {
	if tokens == nil {
		return nil, nil
	}

	httpCl := vantuz.C().
		SetGlobalHeaders(map[string]string{
			"Authorization": "OAuth " + tokens.AccessToken,
		})
	cl := &Client{
		Http: httpCl,
	}

	// get uid
	status, err := cl.AccountStatus(context.Background())
	if err != nil {
		return nil, err
	}
	if status == nil {
		return nil, ErrNilStatus
	}

	if status.Account == nil {
		return nil, ErrNilAccount
	}

	cl.UserId = status.Account.UID
	cl.userId = cl.UserId.String()
	return cl, err
}

// Клиент для запросов к API.
type Client struct {
	// ID текущего пользователя.
	UserId schema.ID

	// Для создания эндпоинтов.
	// Чтоб не конвертировать по 100 раз UserId.
	userId string

	// Отправляет запросы.
	Http *vantuz.Client
}
