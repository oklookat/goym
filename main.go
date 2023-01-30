package goym

import (
	"context"
	"errors"
	"time"

	"github.com/oklookat/goym/auth"
	"github.com/oklookat/goym/schema"
	"github.com/oklookat/goym/vantuz"
)

const (
	errPrefix  = "goym: "
	_userAgent = "goym/v0.2.7 (github.com/oklookat/goym)"
)

var (
	ErrNilResponse = errors.New(errPrefix + "nil http or schema response (dev error?)")
	//
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

	var vCl = vantuz.C().
		SetGlobalHeaders(map[string]string{
			"User-Agent":    _userAgent,
			"Authorization": "OAuth " + tokens.AccessToken,
		})
	var cl = &Client{
		self: vCl,
	}

	// get uid
	status, err := cl.GetAccountStatus(context.Background())
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
	UserId schema.UniqueID

	// Для создания эндпоинтов.
	// Чтоб не конвертировать по 100 раз UserId.
	userId string

	// Отправляет запросы.
	self *vantuz.Client
}

// Включить вывод HTTP запросов в консоль.
func (c Client) EnableDevMode() {
	c.self.EnableDevMode()
}

// Отключить вывод HTTP запросов в консоль.
func (c Client) DisableDevMode() {
	c.self.DisableDevMode()
}

// Установить максимальное количество запросов.
//
// Например: 1 запрос в секунду. Если будет два запроса в секунду, придется ждать 2 секунды.
//
// Стандартное значение: requests = 0 (отключает ограничение).
func (c Client) SetRateLimit(requests int, per time.Duration) {
	c.self.SetRateLimit(requests, per)
}
